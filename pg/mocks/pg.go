// Automatically generated by MockGen. DO NOT EDIT!
// Source: pg/interfaces/pg.go

package mocks

import (
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	gomock "github.com/golang/mock/gomock"
	"github.com/topfreegames/extensions/pg/interfaces"
)

// Mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *_MockDBRecorder
}

// Recorder for MockDB (not exported)
type _MockDBRecorder struct {
	mock *MockDB
}

func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &_MockDBRecorder{mock}
	return mock
}

func (_m *MockDB) EXPECT() *_MockDBRecorder {
	return _m.recorder
}

func (_m *MockDB) Exec(_param0 interface{}, _param1 ...interface{}) (orm.Result, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Exec", _s...)
	ret0, _ := ret[0].(orm.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Exec(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exec", _s...)
}

func (_m *MockDB) ExecOne(_param0 interface{}, _param1 ...interface{}) (orm.Result, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ExecOne", _s...)
	ret0, _ := ret[0].(orm.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) ExecOne(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecOne", _s...)
}

func (_m *MockDB) Query(_param0 interface{}, _param1 interface{}, _param2 ...interface{}) (orm.Result, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Query", _s...)
	ret0, _ := ret[0].(orm.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", _s...)
}

func (_m *MockDB) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDBRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockDB) Begin() (*pg.Tx, error) {
	ret := _m.ctrl.Call(_m, "Begin")
	ret0, _ := ret[0].(*pg.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Begin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Begin")
}

// Mock of Tx interface
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *_MockTxRecorder
}

// Recorder for MockTx (not exported)
type _MockTxRecorder struct {
	mock *MockTx
}

func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &_MockTxRecorder{mock}
	return mock
}

func (_m *MockTx) EXPECT() *_MockTxRecorder {
	return _m.recorder
}

func (_m *MockTx) Exec(_param0 interface{}, _param1 ...interface{}) (orm.Result, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Exec", _s...)
	ret0, _ := ret[0].(orm.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTxRecorder) Exec(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exec", _s...)
}

func (_m *MockTx) ExecOne(_param0 interface{}, _param1 ...interface{}) (orm.Result, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ExecOne", _s...)
	ret0, _ := ret[0].(orm.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTxRecorder) ExecOne(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecOne", _s...)
}

func (_m *MockTx) Query(_param0 interface{}, _param1 interface{}, _param2 ...interface{}) (orm.Result, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Query", _s...)
	ret0, _ := ret[0].(orm.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTxRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", _s...)
}

func (_m *MockTx) Rollback() error {
	ret := _m.ctrl.Call(_m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTxRecorder) Rollback() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rollback")
}

func (_m *MockTx) Commit() error {
	ret := _m.ctrl.Call(_m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTxRecorder) Commit() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Commit")
}

// Mock of TxWrapper interface
type MockTxWrapper struct {
	ctrl     *gomock.Controller
	recorder *_MockTxWrapperRecorder
}

// Recorder for MockTxWrapper (not exported)
type _MockTxWrapperRecorder struct {
	mock *MockTxWrapper
}

func NewMockTxWrapper(ctrl *gomock.Controller) *MockTxWrapper {
	mock := &MockTxWrapper{ctrl: ctrl}
	mock.recorder = &_MockTxWrapperRecorder{mock}
	return mock
}

func (_m *MockTxWrapper) EXPECT() *_MockTxWrapperRecorder {
	return _m.recorder
}

func (_m *MockTxWrapper) DbBegin(db interfaces.DB) (interfaces.Tx, error) {
	ret := _m.ctrl.Call(_m, "DbBegin", db)
	ret0, _ := ret[0].(interfaces.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTxWrapperRecorder) DbBegin(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DbBegin", arg0)
}
