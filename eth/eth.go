// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package eth

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/consensus"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/node"
	"pkg.berachain.dev/polaris/eth/polar"
	"pkg.berachain.dev/polaris/eth/rpc"
)

type (

	// Miner represents the `Miner` that exists on the backend of the execution layer.
	Miner interface {
		BuildPayload(*miner.BuildPayloadArgs) (*miner.Payload, error)
		Etherbase() common.Address
	}

	// TxPool represents the `TxPool` that exists on the backend of the execution layer.
	TxPool interface {
		Add([]*coretypes.Transaction, bool, bool) []error
		Stats() (int, int)
		SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription
	}

	// ExecutionLayerNode is the entrypoint for the evm execution environment.
	NetworkingStack interface {
		// IsExtRPCEnabled returns true if the networking stack is configured to expose JSON-RPC API.
		ExtRPCEnabled() bool

		// RegisterHandler manually registers a new handler into the networking stack.
		RegisterHandler(string, string, http.Handler)

		// RegisterAPIs registers JSON-RPC handlers for the networking stack.
		RegisterAPIs([]rpc.API)

		// RegisterLifecycles registers objects to have their lifecycle manged by the stack.
		RegisterLifecycle(node.Lifecycle)

		// Start starts the networking stack.
		Start() error

		// Close stops the networking stack
		Close() error
	}

	// ExecutionLayer represents the execution layer for a polaris EVM chain.
	ExecutionLayer struct {
		// stack handles all networking aspects of the execution layer. mainly JSON-RPC.
		stack NetworkingStack
		// backend is the entry point to the core logic of the execution layer.
		backend *polar.Polaris
	}

	// Config struct holds the configuration for Polaris and Node.
	Config struct {
		Polar polar.Config
		Node  node.Config
	}
)

// New creates a new execution layer with the provided host chain.
// It takes a client type, configuration, host chain, consensus engine, and log handler
// as parameters. It returns a pointer to the ExecutionLayer and an error if any.
func New(
	client string, cfg any, host core.PolarisHostChain,
	engine consensus.Engine, logHandler log.Handler,
) (*ExecutionLayer, error) {
	clientFactories := map[string]func(
		any, core.PolarisHostChain, consensus.Engine, log.Handler,
	) (*ExecutionLayer, error){
		"geth": newGethExecutionLayer,
	}

	factory, ok := clientFactories[client]
	if !ok {
		return nil, fmt.Errorf("unknown execution layer: %s", client)
	}

	return factory(cfg, host, engine, logHandler)
}

// newGethExecutionLayer creates a new geth execution layer.
// It returns a pointer to the ExecutionLayer and an error if any.
func newGethExecutionLayer(
	anyCfg any, host core.PolarisHostChain,
	engine consensus.Engine, logHandler log.Handler,
) (*ExecutionLayer, error) {
	cfg, ok := anyCfg.(*Config)
	if !ok {
		// If the configuration type is invalid, return an error
		return nil, fmt.Errorf("invalid config type")
	}

	gethNode, err := node.New(&cfg.Node)
	if err != nil {
		return nil, err
	}

	// In Polaris we don't use P2P at the geth level.
	gethNode.SetP2PDisabled(true)

	// Create a new Polaris backend
	backend := polar.New(&cfg.Polar, host, engine, gethNode, logHandler)

	// Return a new ExecutionLayer with the created gethNode and backend
	return &ExecutionLayer{
		stack:   gethNode,
		backend: backend,
	}, nil
}

// RegisterSyncStatusProvider registers a sync status provider to the backend of the
// execution layer.
func (el *ExecutionLayer) RegisterSyncStatusProvider(provider polar.SyncStatusProvider) {
	el.backend.RegisterSyncStatusProvider(provider)
}

// RegisterLifecycle registers a lifecycle to the networking stack of the execution layer.
func (el *ExecutionLayer) RegisterLifecycle(lifecycle node.Lifecycle) {
	el.stack.RegisterLifecycle(lifecycle)
}

// Start starts the networking stack of the execution layer.
// It returns an error if the start operation fails.
func (el *ExecutionLayer) Start() error {
	return el.stack.Start()
}

// Close stops the networking stack of the execution layer.
// It returns an error if the close operation fails.
func (el *ExecutionLayer) Close() error {
	return el.stack.Close()
}

// Miner returns the miner interface of the backend of the execution layer.
func (el *ExecutionLayer) Miner() Miner {
	return el.backend.Miner()
}

// TxPool returns the transaction pool interface of the backend of the execution layer.
func (el *ExecutionLayer) TxPool() TxPool {
	return el.backend.TxPool()
}

// Blockchain returns the blockchain interface of the backend of the execution layer.
func (el *ExecutionLayer) Blockchain() core.Blockchain {
	return el.backend.Blockchain()
}
