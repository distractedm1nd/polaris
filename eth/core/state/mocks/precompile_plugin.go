// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	params "github.com/ethereum/go-ethereum/params"

	vm "github.com/ethereum/go-ethereum/core/vm"
)

// PrecompilePlugin is an autogenerated mock type for the PrecompilePlugin type
type PrecompilePlugin struct {
	mock.Mock
}

type PrecompilePlugin_Expecter struct {
	mock *mock.Mock
}

func (_m *PrecompilePlugin) EXPECT() *PrecompilePlugin_Expecter {
	return &PrecompilePlugin_Expecter{mock: &_m.Mock}
}

// DisableReentrancy provides a mock function with given fields: _a0
func (_m *PrecompilePlugin) DisableReentrancy(_a0 vm.PrecompileEVM) {
	_m.Called(_a0)
}

// PrecompilePlugin_DisableReentrancy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableReentrancy'
type PrecompilePlugin_DisableReentrancy_Call struct {
	*mock.Call
}

// DisableReentrancy is a helper method to define mock.On call
//   - _a0 vm.PrecompileEVM
func (_e *PrecompilePlugin_Expecter) DisableReentrancy(_a0 interface{}) *PrecompilePlugin_DisableReentrancy_Call {
	return &PrecompilePlugin_DisableReentrancy_Call{Call: _e.mock.On("DisableReentrancy", _a0)}
}

func (_c *PrecompilePlugin_DisableReentrancy_Call) Run(run func(_a0 vm.PrecompileEVM)) *PrecompilePlugin_DisableReentrancy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(vm.PrecompileEVM))
	})
	return _c
}

func (_c *PrecompilePlugin_DisableReentrancy_Call) Return() *PrecompilePlugin_DisableReentrancy_Call {
	_c.Call.Return()
	return _c
}

func (_c *PrecompilePlugin_DisableReentrancy_Call) RunAndReturn(run func(vm.PrecompileEVM)) *PrecompilePlugin_DisableReentrancy_Call {
	_c.Call.Return(run)
	return _c
}

// EnableReentrancy provides a mock function with given fields: _a0
func (_m *PrecompilePlugin) EnableReentrancy(_a0 vm.PrecompileEVM) {
	_m.Called(_a0)
}

// PrecompilePlugin_EnableReentrancy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnableReentrancy'
type PrecompilePlugin_EnableReentrancy_Call struct {
	*mock.Call
}

// EnableReentrancy is a helper method to define mock.On call
//   - _a0 vm.PrecompileEVM
func (_e *PrecompilePlugin_Expecter) EnableReentrancy(_a0 interface{}) *PrecompilePlugin_EnableReentrancy_Call {
	return &PrecompilePlugin_EnableReentrancy_Call{Call: _e.mock.On("EnableReentrancy", _a0)}
}

func (_c *PrecompilePlugin_EnableReentrancy_Call) Run(run func(_a0 vm.PrecompileEVM)) *PrecompilePlugin_EnableReentrancy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(vm.PrecompileEVM))
	})
	return _c
}

func (_c *PrecompilePlugin_EnableReentrancy_Call) Return() *PrecompilePlugin_EnableReentrancy_Call {
	_c.Call.Return()
	return _c
}

func (_c *PrecompilePlugin_EnableReentrancy_Call) RunAndReturn(run func(vm.PrecompileEVM)) *PrecompilePlugin_EnableReentrancy_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: addr, rules
func (_m *PrecompilePlugin) Get(addr common.Address, rules *params.Rules) (vm.PrecompiledContract, bool) {
	ret := _m.Called(addr, rules)

	var r0 vm.PrecompiledContract
	var r1 bool
	if rf, ok := ret.Get(0).(func(common.Address, *params.Rules) (vm.PrecompiledContract, bool)); ok {
		return rf(addr, rules)
	}
	if rf, ok := ret.Get(0).(func(common.Address, *params.Rules) vm.PrecompiledContract); ok {
		r0 = rf(addr, rules)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vm.PrecompiledContract)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, *params.Rules) bool); ok {
		r1 = rf(addr, rules)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// PrecompilePlugin_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PrecompilePlugin_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - addr common.Address
//   - rules *params.Rules
func (_e *PrecompilePlugin_Expecter) Get(addr interface{}, rules interface{}) *PrecompilePlugin_Get_Call {
	return &PrecompilePlugin_Get_Call{Call: _e.mock.On("Get", addr, rules)}
}

func (_c *PrecompilePlugin_Get_Call) Run(run func(addr common.Address, rules *params.Rules)) *PrecompilePlugin_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(*params.Rules))
	})
	return _c
}

func (_c *PrecompilePlugin_Get_Call) Return(_a0 vm.PrecompiledContract, _a1 bool) *PrecompilePlugin_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrecompilePlugin_Get_Call) RunAndReturn(run func(common.Address, *params.Rules) (vm.PrecompiledContract, bool)) *PrecompilePlugin_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetActive provides a mock function with given fields: _a0
func (_m *PrecompilePlugin) GetActive(_a0 params.Rules) []common.Address {
	ret := _m.Called(_a0)

	var r0 []common.Address
	if rf, ok := ret.Get(0).(func(params.Rules) []common.Address); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	return r0
}

// PrecompilePlugin_GetActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActive'
type PrecompilePlugin_GetActive_Call struct {
	*mock.Call
}

// GetActive is a helper method to define mock.On call
//   - _a0 params.Rules
func (_e *PrecompilePlugin_Expecter) GetActive(_a0 interface{}) *PrecompilePlugin_GetActive_Call {
	return &PrecompilePlugin_GetActive_Call{Call: _e.mock.On("GetActive", _a0)}
}

func (_c *PrecompilePlugin_GetActive_Call) Run(run func(_a0 params.Rules)) *PrecompilePlugin_GetActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(params.Rules))
	})
	return _c
}

func (_c *PrecompilePlugin_GetActive_Call) Return(_a0 []common.Address) *PrecompilePlugin_GetActive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PrecompilePlugin_GetActive_Call) RunAndReturn(run func(params.Rules) []common.Address) *PrecompilePlugin_GetActive_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *PrecompilePlugin) Register(_a0 vm.PrecompiledContract) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(vm.PrecompiledContract) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrecompilePlugin_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type PrecompilePlugin_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 vm.PrecompiledContract
func (_e *PrecompilePlugin_Expecter) Register(_a0 interface{}) *PrecompilePlugin_Register_Call {
	return &PrecompilePlugin_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *PrecompilePlugin_Register_Call) Run(run func(_a0 vm.PrecompiledContract)) *PrecompilePlugin_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(vm.PrecompiledContract))
	})
	return _c
}

func (_c *PrecompilePlugin_Register_Call) Return(_a0 error) *PrecompilePlugin_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PrecompilePlugin_Register_Call) RunAndReturn(run func(vm.PrecompiledContract) error) *PrecompilePlugin_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: evm, p, input, caller, value, suppliedGas, readonly
func (_m *PrecompilePlugin) Run(evm vm.PrecompileEVM, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error) {
	ret := _m.Called(evm, p, input, caller, value, suppliedGas, readonly)

	var r0 []byte
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(vm.PrecompileEVM, vm.PrecompiledContract, []byte, common.Address, *big.Int, uint64, bool) ([]byte, uint64, error)); ok {
		return rf(evm, p, input, caller, value, suppliedGas, readonly)
	}
	if rf, ok := ret.Get(0).(func(vm.PrecompileEVM, vm.PrecompiledContract, []byte, common.Address, *big.Int, uint64, bool) []byte); ok {
		r0 = rf(evm, p, input, caller, value, suppliedGas, readonly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(vm.PrecompileEVM, vm.PrecompiledContract, []byte, common.Address, *big.Int, uint64, bool) uint64); ok {
		r1 = rf(evm, p, input, caller, value, suppliedGas, readonly)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(vm.PrecompileEVM, vm.PrecompiledContract, []byte, common.Address, *big.Int, uint64, bool) error); ok {
		r2 = rf(evm, p, input, caller, value, suppliedGas, readonly)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PrecompilePlugin_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type PrecompilePlugin_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - evm vm.PrecompileEVM
//   - p vm.PrecompiledContract
//   - input []byte
//   - caller common.Address
//   - value *big.Int
//   - suppliedGas uint64
//   - readonly bool
func (_e *PrecompilePlugin_Expecter) Run(evm interface{}, p interface{}, input interface{}, caller interface{}, value interface{}, suppliedGas interface{}, readonly interface{}) *PrecompilePlugin_Run_Call {
	return &PrecompilePlugin_Run_Call{Call: _e.mock.On("Run", evm, p, input, caller, value, suppliedGas, readonly)}
}

func (_c *PrecompilePlugin_Run_Call) Run(run func(evm vm.PrecompileEVM, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool)) *PrecompilePlugin_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(vm.PrecompileEVM), args[1].(vm.PrecompiledContract), args[2].([]byte), args[3].(common.Address), args[4].(*big.Int), args[5].(uint64), args[6].(bool))
	})
	return _c
}

func (_c *PrecompilePlugin_Run_Call) Return(ret []byte, remainingGas uint64, err error) *PrecompilePlugin_Run_Call {
	_c.Call.Return(ret, remainingGas, err)
	return _c
}

func (_c *PrecompilePlugin_Run_Call) RunAndReturn(run func(vm.PrecompileEVM, vm.PrecompiledContract, []byte, common.Address, *big.Int, uint64, bool) ([]byte, uint64, error)) *PrecompilePlugin_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewPrecompilePlugin creates a new instance of PrecompilePlugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPrecompilePlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *PrecompilePlugin {
	mock := &PrecompilePlugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
