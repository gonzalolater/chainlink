// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/smartcontractkit/chainlink/core/utils"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// Chain provides a mock function with given fields: id
func (_m *ORM) Chain(id utils.Big) (types.Chain, error) {
	ret := _m.Called(id)

	var r0 types.Chain
	if rf, ok := ret.Get(0).(func(utils.Big) types.Chain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(types.Chain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(utils.Big) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Chains provides a mock function with given fields: offset, limit
func (_m *ORM) Chains(offset int, limit int) ([]types.Chain, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []types.Chain
	if rf, ok := ret.Get(0).(func(int, int) []types.Chain); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Chain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Clear provides a mock function with given fields: chainID, key
func (_m *ORM) Clear(chainID *big.Int, key string) error {
	ret := _m.Called(chainID, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, string) error); ok {
		r0 = rf(chainID, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateChain provides a mock function with given fields: id, config
func (_m *ORM) CreateChain(id utils.Big, config types.ChainCfg) (types.Chain, error) {
	ret := _m.Called(id, config)

	var r0 types.Chain
	if rf, ok := ret.Get(0).(func(utils.Big, types.ChainCfg) types.Chain); ok {
		r0 = rf(id, config)
	} else {
		r0 = ret.Get(0).(types.Chain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(utils.Big, types.ChainCfg) error); ok {
		r1 = rf(id, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNode provides a mock function with given fields: data
func (_m *ORM) CreateNode(data types.NewNode) (types.Node, error) {
	ret := _m.Called(data)

	var r0 types.Node
	if rf, ok := ret.Get(0).(func(types.NewNode) types.Node); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(types.Node)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.NewNode) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChain provides a mock function with given fields: id
func (_m *ORM) DeleteChain(id utils.Big) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(utils.Big) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNode provides a mock function with given fields: id
func (_m *ORM) DeleteNode(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnabledChainsWithNodes provides a mock function with given fields:
func (_m *ORM) EnabledChainsWithNodes() ([]types.Chain, error) {
	ret := _m.Called()

	var r0 []types.Chain
	if rf, ok := ret.Get(0).(func() []types.Chain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Chain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChainsByIDs provides a mock function with given fields: ids
func (_m *ORM) GetChainsByIDs(ids []utils.Big) ([]types.Chain, error) {
	ret := _m.Called(ids)

	var r0 []types.Chain
	if rf, ok := ret.Get(0).(func([]utils.Big) []types.Chain); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Chain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]utils.Big) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodesByChainIDs provides a mock function with given fields: chainIDs
func (_m *ORM) GetNodesByChainIDs(chainIDs []utils.Big) ([]types.Node, error) {
	ret := _m.Called(chainIDs)

	var r0 []types.Node
	if rf, ok := ret.Get(0).(func([]utils.Big) []types.Node); ok {
		r0 = rf(chainIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]utils.Big) error); ok {
		r1 = rf(chainIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Nodes provides a mock function with given fields: offset, limit
func (_m *ORM) Nodes(offset int, limit int) ([]types.Node, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []types.Node
	if rf, ok := ret.Get(0).(func(int, int) []types.Node); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Node)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NodesForChain provides a mock function with given fields: chainID, offset, limit
func (_m *ORM) NodesForChain(chainID utils.Big, offset int, limit int) ([]types.Node, int, error) {
	ret := _m.Called(chainID, offset, limit)

	var r0 []types.Node
	if rf, ok := ret.Get(0).(func(utils.Big, int, int) []types.Node); ok {
		r0 = rf(chainID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Node)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(utils.Big, int, int) int); ok {
		r1 = rf(chainID, offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(utils.Big, int, int) error); ok {
		r2 = rf(chainID, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StoreString provides a mock function with given fields: chainID, key, val
func (_m *ORM) StoreString(chainID *big.Int, key string, val string) error {
	ret := _m.Called(chainID, key, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, string, string) error); ok {
		r0 = rf(chainID, key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateChain provides a mock function with given fields: id, enabled, config
func (_m *ORM) UpdateChain(id utils.Big, enabled bool, config types.ChainCfg) (types.Chain, error) {
	ret := _m.Called(id, enabled, config)

	var r0 types.Chain
	if rf, ok := ret.Get(0).(func(utils.Big, bool, types.ChainCfg) types.Chain); ok {
		r0 = rf(id, enabled, config)
	} else {
		r0 = ret.Get(0).(types.Chain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(utils.Big, bool, types.ChainCfg) error); ok {
		r1 = rf(id, enabled, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
