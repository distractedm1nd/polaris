#!/bin/bash

# Stage 1: Clone repos
cd ~/
mkdir op-stack-deployment
cd op-stack-deployment
git clone https://github.com/celestiaorg/optimism.git
git clone https://github.com/ethereum-optimism/op-geth.git

# Stage 2: Setup nvm
brew install nvm

# Check if ~/.nvm directory doesn't exist
if [ ! -d "$nvm_dir" ]; then
    mkdir "$nvm_dir"
    echo "Created ~/.nvm directory."
fi
# hardhat needs ^16.0.0
. ~/.nvm/nvm.sh
. ~/.zshrc
. $(brew --prefix nvm)/nvm.sh  # if installed via Brew
nvm use 16

# Stage 3: Install op-node op-batcher op-proposer
cd optimism
git checkout v0.1.1-OP_v1.0.6-CN_v0.11.0-rc1
yarn install
make op-node op-batcher op-proposer
yarn build
cd ..

# Stage 4: Install op-geth
cd op-geth
git checkout v1.101105.1
make geth
cd ..