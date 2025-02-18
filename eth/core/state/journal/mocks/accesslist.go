// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	journal "pkg.berachain.dev/polaris/eth/core/state/journal"

	mock "github.com/stretchr/testify/mock"
)

// Accesslist is an autogenerated mock type for the Accesslist type
type Accesslist struct {
	mock.Mock
}

type Accesslist_Expecter struct {
	mock *mock.Mock
}

func (_m *Accesslist) EXPECT() *Accesslist_Expecter {
	return &Accesslist_Expecter{mock: &_m.Mock}
}

// AddAddressToAccessList provides a mock function with given fields: _a0
func (_m *Accesslist) AddAddressToAccessList(_a0 common.Address) {
	_m.Called(_a0)
}

// Accesslist_AddAddressToAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAddressToAccessList'
type Accesslist_AddAddressToAccessList_Call struct {
	*mock.Call
}

// AddAddressToAccessList is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Accesslist_Expecter) AddAddressToAccessList(_a0 interface{}) *Accesslist_AddAddressToAccessList_Call {
	return &Accesslist_AddAddressToAccessList_Call{Call: _e.mock.On("AddAddressToAccessList", _a0)}
}

func (_c *Accesslist_AddAddressToAccessList_Call) Run(run func(_a0 common.Address)) *Accesslist_AddAddressToAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Accesslist_AddAddressToAccessList_Call) Return() *Accesslist_AddAddressToAccessList_Call {
	_c.Call.Return()
	return _c
}

func (_c *Accesslist_AddAddressToAccessList_Call) RunAndReturn(run func(common.Address)) *Accesslist_AddAddressToAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// AddSlotToAccessList provides a mock function with given fields: _a0, _a1
func (_m *Accesslist) AddSlotToAccessList(_a0 common.Address, _a1 common.Hash) {
	_m.Called(_a0, _a1)
}

// Accesslist_AddSlotToAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSlotToAccessList'
type Accesslist_AddSlotToAccessList_Call struct {
	*mock.Call
}

// AddSlotToAccessList is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 common.Hash
func (_e *Accesslist_Expecter) AddSlotToAccessList(_a0 interface{}, _a1 interface{}) *Accesslist_AddSlotToAccessList_Call {
	return &Accesslist_AddSlotToAccessList_Call{Call: _e.mock.On("AddSlotToAccessList", _a0, _a1)}
}

func (_c *Accesslist_AddSlotToAccessList_Call) Run(run func(_a0 common.Address, _a1 common.Hash)) *Accesslist_AddSlotToAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *Accesslist_AddSlotToAccessList_Call) Return() *Accesslist_AddSlotToAccessList_Call {
	_c.Call.Return()
	return _c
}

func (_c *Accesslist_AddSlotToAccessList_Call) RunAndReturn(run func(common.Address, common.Hash)) *Accesslist_AddSlotToAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// AddressInAccessList provides a mock function with given fields: _a0
func (_m *Accesslist) AddressInAccessList(_a0 common.Address) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Accesslist_AddressInAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddressInAccessList'
type Accesslist_AddressInAccessList_Call struct {
	*mock.Call
}

// AddressInAccessList is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Accesslist_Expecter) AddressInAccessList(_a0 interface{}) *Accesslist_AddressInAccessList_Call {
	return &Accesslist_AddressInAccessList_Call{Call: _e.mock.On("AddressInAccessList", _a0)}
}

func (_c *Accesslist_AddressInAccessList_Call) Run(run func(_a0 common.Address)) *Accesslist_AddressInAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Accesslist_AddressInAccessList_Call) Return(_a0 bool) *Accesslist_AddressInAccessList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Accesslist_AddressInAccessList_Call) RunAndReturn(run func(common.Address) bool) *Accesslist_AddressInAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// Clone provides a mock function with given fields:
func (_m *Accesslist) Clone() journal.Accesslist {
	ret := _m.Called()

	var r0 journal.Accesslist
	if rf, ok := ret.Get(0).(func() journal.Accesslist); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(journal.Accesslist)
		}
	}

	return r0
}

// Accesslist_Clone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clone'
type Accesslist_Clone_Call struct {
	*mock.Call
}

// Clone is a helper method to define mock.On call
func (_e *Accesslist_Expecter) Clone() *Accesslist_Clone_Call {
	return &Accesslist_Clone_Call{Call: _e.mock.On("Clone")}
}

func (_c *Accesslist_Clone_Call) Run(run func()) *Accesslist_Clone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Accesslist_Clone_Call) Return(_a0 journal.Accesslist) *Accesslist_Clone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Accesslist_Clone_Call) RunAndReturn(run func() journal.Accesslist) *Accesslist_Clone_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields:
func (_m *Accesslist) Finalize() {
	_m.Called()
}

// Accesslist_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type Accesslist_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
func (_e *Accesslist_Expecter) Finalize() *Accesslist_Finalize_Call {
	return &Accesslist_Finalize_Call{Call: _e.mock.On("Finalize")}
}

func (_c *Accesslist_Finalize_Call) Run(run func()) *Accesslist_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Accesslist_Finalize_Call) Return() *Accesslist_Finalize_Call {
	_c.Call.Return()
	return _c
}

func (_c *Accesslist_Finalize_Call) RunAndReturn(run func()) *Accesslist_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// RegistryKey provides a mock function with given fields:
func (_m *Accesslist) RegistryKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Accesslist_RegistryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegistryKey'
type Accesslist_RegistryKey_Call struct {
	*mock.Call
}

// RegistryKey is a helper method to define mock.On call
func (_e *Accesslist_Expecter) RegistryKey() *Accesslist_RegistryKey_Call {
	return &Accesslist_RegistryKey_Call{Call: _e.mock.On("RegistryKey")}
}

func (_c *Accesslist_RegistryKey_Call) Run(run func()) *Accesslist_RegistryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Accesslist_RegistryKey_Call) Return(_a0 string) *Accesslist_RegistryKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Accesslist_RegistryKey_Call) RunAndReturn(run func() string) *Accesslist_RegistryKey_Call {
	_c.Call.Return(run)
	return _c
}

// RevertToSnapshot provides a mock function with given fields: _a0
func (_m *Accesslist) RevertToSnapshot(_a0 int) {
	_m.Called(_a0)
}

// Accesslist_RevertToSnapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertToSnapshot'
type Accesslist_RevertToSnapshot_Call struct {
	*mock.Call
}

// RevertToSnapshot is a helper method to define mock.On call
//   - _a0 int
func (_e *Accesslist_Expecter) RevertToSnapshot(_a0 interface{}) *Accesslist_RevertToSnapshot_Call {
	return &Accesslist_RevertToSnapshot_Call{Call: _e.mock.On("RevertToSnapshot", _a0)}
}

func (_c *Accesslist_RevertToSnapshot_Call) Run(run func(_a0 int)) *Accesslist_RevertToSnapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Accesslist_RevertToSnapshot_Call) Return() *Accesslist_RevertToSnapshot_Call {
	_c.Call.Return()
	return _c
}

func (_c *Accesslist_RevertToSnapshot_Call) RunAndReturn(run func(int)) *Accesslist_RevertToSnapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SlotInAccessList provides a mock function with given fields: _a0, _a1
func (_m *Accesslist) SlotInAccessList(_a0 common.Address, _a1 common.Hash) (bool, bool) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	var r1 bool
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) (bool, bool)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(common.Address, common.Hash) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Accesslist_SlotInAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SlotInAccessList'
type Accesslist_SlotInAccessList_Call struct {
	*mock.Call
}

// SlotInAccessList is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 common.Hash
func (_e *Accesslist_Expecter) SlotInAccessList(_a0 interface{}, _a1 interface{}) *Accesslist_SlotInAccessList_Call {
	return &Accesslist_SlotInAccessList_Call{Call: _e.mock.On("SlotInAccessList", _a0, _a1)}
}

func (_c *Accesslist_SlotInAccessList_Call) Run(run func(_a0 common.Address, _a1 common.Hash)) *Accesslist_SlotInAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *Accesslist_SlotInAccessList_Call) Return(addressPresent bool, slotPresent bool) *Accesslist_SlotInAccessList_Call {
	_c.Call.Return(addressPresent, slotPresent)
	return _c
}

func (_c *Accesslist_SlotInAccessList_Call) RunAndReturn(run func(common.Address, common.Hash) (bool, bool)) *Accesslist_SlotInAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// Snapshot provides a mock function with given fields:
func (_m *Accesslist) Snapshot() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Accesslist_Snapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Snapshot'
type Accesslist_Snapshot_Call struct {
	*mock.Call
}

// Snapshot is a helper method to define mock.On call
func (_e *Accesslist_Expecter) Snapshot() *Accesslist_Snapshot_Call {
	return &Accesslist_Snapshot_Call{Call: _e.mock.On("Snapshot")}
}

func (_c *Accesslist_Snapshot_Call) Run(run func()) *Accesslist_Snapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Accesslist_Snapshot_Call) Return(_a0 int) *Accesslist_Snapshot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Accesslist_Snapshot_Call) RunAndReturn(run func() int) *Accesslist_Snapshot_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccesslist creates a new instance of Accesslist. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccesslist(t interface {
	mock.TestingT
	Cleanup(func())
}) *Accesslist {
	mock := &Accesslist{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
