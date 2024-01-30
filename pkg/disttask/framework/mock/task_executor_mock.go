// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/pkg/disttask/framework/taskexecutor (interfaces: TaskTable,Pool,TaskExecutor,Extension)
//
// Generated by this command:
//
//	mockgen -package mock github.com/pingcap/tidb/pkg/disttask/framework/taskexecutor TaskTable,Pool,TaskExecutor,Extension
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	proto "github.com/pingcap/tidb/pkg/disttask/framework/proto"
	storage "github.com/pingcap/tidb/pkg/disttask/framework/storage"
	execute "github.com/pingcap/tidb/pkg/disttask/framework/taskexecutor/execute"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskTable is a mock of TaskTable interface.
type MockTaskTable struct {
	ctrl     *gomock.Controller
	recorder *MockTaskTableMockRecorder
}

// MockTaskTableMockRecorder is the mock recorder for MockTaskTable.
type MockTaskTableMockRecorder struct {
	mock *MockTaskTable
}

// NewMockTaskTable creates a new mock instance.
func NewMockTaskTable(ctrl *gomock.Controller) *MockTaskTable {
	mock := &MockTaskTable{ctrl: ctrl}
	mock.recorder = &MockTaskTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskTable) EXPECT() *MockTaskTableMockRecorder {
	return m.recorder
}

// CancelSubtask mocks base method.
func (m *MockTaskTable) CancelSubtask(arg0 context.Context, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelSubtask", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelSubtask indicates an expected call of CancelSubtask.
func (mr *MockTaskTableMockRecorder) CancelSubtask(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSubtask", reflect.TypeOf((*MockTaskTable)(nil).CancelSubtask), arg0, arg1, arg2)
}

// FailSubtask mocks base method.
func (m *MockTaskTable) FailSubtask(arg0 context.Context, arg1 string, arg2 int64, arg3 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailSubtask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FailSubtask indicates an expected call of FailSubtask.
func (mr *MockTaskTableMockRecorder) FailSubtask(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailSubtask", reflect.TypeOf((*MockTaskTable)(nil).FailSubtask), arg0, arg1, arg2, arg3)
}

// FinishSubtask mocks base method.
func (m *MockTaskTable) FinishSubtask(arg0 context.Context, arg1 string, arg2 int64, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishSubtask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishSubtask indicates an expected call of FinishSubtask.
func (mr *MockTaskTableMockRecorder) FinishSubtask(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishSubtask", reflect.TypeOf((*MockTaskTable)(nil).FinishSubtask), arg0, arg1, arg2, arg3)
}

// GetFirstSubtaskInStates mocks base method.
func (m *MockTaskTable) GetFirstSubtaskInStates(arg0 context.Context, arg1 string, arg2 int64, arg3 proto.Step, arg4 ...proto.SubtaskState) (*proto.Subtask, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFirstSubtaskInStates", varargs...)
	ret0, _ := ret[0].(*proto.Subtask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstSubtaskInStates indicates an expected call of GetFirstSubtaskInStates.
func (mr *MockTaskTableMockRecorder) GetFirstSubtaskInStates(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstSubtaskInStates", reflect.TypeOf((*MockTaskTable)(nil).GetFirstSubtaskInStates), varargs...)
}

// GetSubtasksByExecIDAndStepAndStates mocks base method.
func (m *MockTaskTable) GetSubtasksByExecIDAndStepAndStates(arg0 context.Context, arg1 string, arg2 int64, arg3 proto.Step, arg4 ...proto.SubtaskState) ([]*proto.Subtask, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubtasksByExecIDAndStepAndStates", varargs...)
	ret0, _ := ret[0].([]*proto.Subtask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubtasksByExecIDAndStepAndStates indicates an expected call of GetSubtasksByExecIDAndStepAndStates.
func (mr *MockTaskTableMockRecorder) GetSubtasksByExecIDAndStepAndStates(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubtasksByExecIDAndStepAndStates", reflect.TypeOf((*MockTaskTable)(nil).GetSubtasksByExecIDAndStepAndStates), varargs...)
}

// GetTaskByID mocks base method.
func (m *MockTaskTable) GetTaskByID(arg0 context.Context, arg1 int64) (*proto.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", arg0, arg1)
	ret0, _ := ret[0].(*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockTaskTableMockRecorder) GetTaskByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTaskTable)(nil).GetTaskByID), arg0, arg1)
}

// GetTaskExecInfoByExecID mocks base method.
func (m *MockTaskTable) GetTaskExecInfoByExecID(arg0 context.Context, arg1 string) ([]*storage.TaskExecInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskExecInfoByExecID", arg0, arg1)
	ret0, _ := ret[0].([]*storage.TaskExecInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskExecInfoByExecID indicates an expected call of GetTaskExecInfoByExecID.
func (mr *MockTaskTableMockRecorder) GetTaskExecInfoByExecID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskExecInfoByExecID", reflect.TypeOf((*MockTaskTable)(nil).GetTaskExecInfoByExecID), arg0, arg1)
}

// GetTasksInStates mocks base method.
func (m *MockTaskTable) GetTasksInStates(arg0 context.Context, arg1 ...any) ([]*proto.Task, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTasksInStates", varargs...)
	ret0, _ := ret[0].([]*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksInStates indicates an expected call of GetTasksInStates.
func (mr *MockTaskTableMockRecorder) GetTasksInStates(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksInStates", reflect.TypeOf((*MockTaskTable)(nil).GetTasksInStates), varargs...)
}

// HasSubtasksInStates mocks base method.
func (m *MockTaskTable) HasSubtasksInStates(arg0 context.Context, arg1 string, arg2 int64, arg3 proto.Step, arg4 ...proto.SubtaskState) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasSubtasksInStates", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasSubtasksInStates indicates an expected call of HasSubtasksInStates.
func (mr *MockTaskTableMockRecorder) HasSubtasksInStates(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSubtasksInStates", reflect.TypeOf((*MockTaskTable)(nil).HasSubtasksInStates), varargs...)
}

// InitMeta mocks base method.
func (m *MockTaskTable) InitMeta(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitMeta", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitMeta indicates an expected call of InitMeta.
func (mr *MockTaskTableMockRecorder) InitMeta(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitMeta", reflect.TypeOf((*MockTaskTable)(nil).InitMeta), arg0, arg1, arg2)
}

// PauseSubtasks mocks base method.
func (m *MockTaskTable) PauseSubtasks(arg0 context.Context, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseSubtasks", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PauseSubtasks indicates an expected call of PauseSubtasks.
func (mr *MockTaskTableMockRecorder) PauseSubtasks(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseSubtasks", reflect.TypeOf((*MockTaskTable)(nil).PauseSubtasks), arg0, arg1, arg2)
}

// RecoverMeta mocks base method.
func (m *MockTaskTable) RecoverMeta(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecoverMeta", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecoverMeta indicates an expected call of RecoverMeta.
func (mr *MockTaskTableMockRecorder) RecoverMeta(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverMeta", reflect.TypeOf((*MockTaskTable)(nil).RecoverMeta), arg0, arg1, arg2)
}

// RunningSubtasksBack2Pending mocks base method.
func (m *MockTaskTable) RunningSubtasksBack2Pending(arg0 context.Context, arg1 []*proto.Subtask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunningSubtasksBack2Pending", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunningSubtasksBack2Pending indicates an expected call of RunningSubtasksBack2Pending.
func (mr *MockTaskTableMockRecorder) RunningSubtasksBack2Pending(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunningSubtasksBack2Pending", reflect.TypeOf((*MockTaskTable)(nil).RunningSubtasksBack2Pending), arg0, arg1)
}

// StartSubtask mocks base method.
func (m *MockTaskTable) StartSubtask(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSubtask", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartSubtask indicates an expected call of StartSubtask.
func (mr *MockTaskTableMockRecorder) StartSubtask(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSubtask", reflect.TypeOf((*MockTaskTable)(nil).StartSubtask), arg0, arg1, arg2)
}

// UpdateSubtaskStateAndError mocks base method.
func (m *MockTaskTable) UpdateSubtaskStateAndError(arg0 context.Context, arg1 string, arg2 int64, arg3 proto.SubtaskState, arg4 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubtaskStateAndError", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubtaskStateAndError indicates an expected call of UpdateSubtaskStateAndError.
func (mr *MockTaskTableMockRecorder) UpdateSubtaskStateAndError(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubtaskStateAndError", reflect.TypeOf((*MockTaskTable)(nil).UpdateSubtaskStateAndError), arg0, arg1, arg2, arg3, arg4)
}

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// ReleaseAndWait mocks base method.
func (m *MockPool) ReleaseAndWait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseAndWait")
}

// ReleaseAndWait indicates an expected call of ReleaseAndWait.
func (mr *MockPoolMockRecorder) ReleaseAndWait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseAndWait", reflect.TypeOf((*MockPool)(nil).ReleaseAndWait))
}

// Run mocks base method.
func (m *MockPool) Run(arg0 func()) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockPoolMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPool)(nil).Run), arg0)
}

// RunWithConcurrency mocks base method.
func (m *MockPool) RunWithConcurrency(arg0 chan func(), arg1 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunWithConcurrency", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWithConcurrency indicates an expected call of RunWithConcurrency.
func (mr *MockPoolMockRecorder) RunWithConcurrency(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWithConcurrency", reflect.TypeOf((*MockPool)(nil).RunWithConcurrency), arg0, arg1)
}

// MockTaskExecutor is a mock of TaskExecutor interface.
type MockTaskExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockTaskExecutorMockRecorder
}

// MockTaskExecutorMockRecorder is the mock recorder for MockTaskExecutor.
type MockTaskExecutorMockRecorder struct {
	mock *MockTaskExecutor
}

// NewMockTaskExecutor creates a new mock instance.
func NewMockTaskExecutor(ctrl *gomock.Controller) *MockTaskExecutor {
	mock := &MockTaskExecutor{ctrl: ctrl}
	mock.recorder = &MockTaskExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskExecutor) EXPECT() *MockTaskExecutorMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockTaskExecutor) Cancel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel")
}

// Cancel indicates an expected call of Cancel.
func (mr *MockTaskExecutorMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockTaskExecutor)(nil).Cancel))
}

// CancelRunningSubtask mocks base method.
func (m *MockTaskExecutor) CancelRunningSubtask() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRunningSubtask")
}

// CancelRunningSubtask indicates an expected call of CancelRunningSubtask.
func (mr *MockTaskExecutorMockRecorder) CancelRunningSubtask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRunningSubtask", reflect.TypeOf((*MockTaskExecutor)(nil).CancelRunningSubtask))
}

// Close mocks base method.
func (m *MockTaskExecutor) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockTaskExecutorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTaskExecutor)(nil).Close))
}

// GetTask mocks base method.
func (m *MockTaskExecutor) GetTask() *proto.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask")
	ret0, _ := ret[0].(*proto.Task)
	return ret0
}

// GetTask indicates an expected call of GetTask.
func (mr *MockTaskExecutorMockRecorder) GetTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTaskExecutor)(nil).GetTask))
}

// Init mocks base method.
func (m *MockTaskExecutor) Init(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockTaskExecutorMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockTaskExecutor)(nil).Init), arg0)
}

// IsRetryableError mocks base method.
func (m *MockTaskExecutor) IsRetryableError(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRetryableError", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRetryableError indicates an expected call of IsRetryableError.
func (mr *MockTaskExecutorMockRecorder) IsRetryableError(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRetryableError", reflect.TypeOf((*MockTaskExecutor)(nil).IsRetryableError), arg0)
}

// Run mocks base method.
func (m *MockTaskExecutor) Run(arg0 *proto.StepResource) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", arg0)
}

// Run indicates an expected call of Run.
func (mr *MockTaskExecutorMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTaskExecutor)(nil).Run), arg0)
}

// MockExtension is a mock of Extension interface.
type MockExtension struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionMockRecorder
}

// MockExtensionMockRecorder is the mock recorder for MockExtension.
type MockExtensionMockRecorder struct {
	mock *MockExtension
}

// NewMockExtension creates a new mock instance.
func NewMockExtension(ctrl *gomock.Controller) *MockExtension {
	mock := &MockExtension{ctrl: ctrl}
	mock.recorder = &MockExtensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtension) EXPECT() *MockExtensionMockRecorder {
	return m.recorder
}

// GetStepExecutor mocks base method.
func (m *MockExtension) GetStepExecutor(arg0 *proto.Task, arg1 *execute.Summary, arg2 *proto.StepResource) (execute.StepExecutor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStepExecutor", arg0, arg1, arg2)
	ret0, _ := ret[0].(execute.StepExecutor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStepExecutor indicates an expected call of GetStepExecutor.
func (mr *MockExtensionMockRecorder) GetStepExecutor(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStepExecutor", reflect.TypeOf((*MockExtension)(nil).GetStepExecutor), arg0, arg1, arg2)
}

// IsIdempotent mocks base method.
func (m *MockExtension) IsIdempotent(arg0 *proto.Subtask) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIdempotent", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIdempotent indicates an expected call of IsIdempotent.
func (mr *MockExtensionMockRecorder) IsIdempotent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIdempotent", reflect.TypeOf((*MockExtension)(nil).IsIdempotent), arg0)
}

// IsRetryableError mocks base method.
func (m *MockExtension) IsRetryableError(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRetryableError", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRetryableError indicates an expected call of IsRetryableError.
func (mr *MockExtensionMockRecorder) IsRetryableError(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRetryableError", reflect.TypeOf((*MockExtension)(nil).IsRetryableError), arg0)
}
