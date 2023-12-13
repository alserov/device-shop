// Code generated by MockGen. DO NOT EDIT.
// Source: internal/db/repo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	models "github.com/alserov/device-shop/device-service/internal/db/models"
	gomock "github.com/golang/mock/gomock"
)

// MockDeviceRepo is a mock of DeviceRepo interface.
type MockDeviceRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceRepoMockRecorder
}

// MockDeviceRepoMockRecorder is the mock recorder for MockDeviceRepo.
type MockDeviceRepoMockRecorder struct {
	mock *MockDeviceRepo
}

// NewMockDeviceRepo creates a new mock instance.
func NewMockDeviceRepo(ctrl *gomock.Controller) *MockDeviceRepo {
	mock := &MockDeviceRepo{ctrl: ctrl}
	mock.recorder = &MockDeviceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceRepo) EXPECT() *MockDeviceRepoMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockDeviceRepo) CreateDevice(arg0 context.Context, arg1 models.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockDeviceRepoMockRecorder) CreateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockDeviceRepo)(nil).CreateDevice), arg0, arg1)
}

// DecreaseDevicesAmountTx mocks base method.
func (m *MockDeviceRepo) DecreaseDevicesAmountTx(ctx context.Context, devices []*models.OrderDevice) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseDevicesAmountTx", ctx, devices)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseDevicesAmountTx indicates an expected call of DecreaseDevicesAmountTx.
func (mr *MockDeviceRepoMockRecorder) DecreaseDevicesAmountTx(ctx, devices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseDevicesAmountTx", reflect.TypeOf((*MockDeviceRepo)(nil).DecreaseDevicesAmountTx), ctx, devices)
}

// DeleteDevice mocks base method.
func (m *MockDeviceRepo) DeleteDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockDeviceRepoMockRecorder) DeleteDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockDeviceRepo)(nil).DeleteDevice), arg0, arg1)
}

// GetAllDevices mocks base method.
func (m *MockDeviceRepo) GetAllDevices(ctx context.Context, index, amount uint32) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevices", ctx, index, amount)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDevices indicates an expected call of GetAllDevices.
func (mr *MockDeviceRepoMockRecorder) GetAllDevices(ctx, index, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevices", reflect.TypeOf((*MockDeviceRepo)(nil).GetAllDevices), ctx, index, amount)
}

// GetDeviceByUUID mocks base method.
func (m *MockDeviceRepo) GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUID", ctx, uuid)
	ret0, _ := ret[0].(models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID.
func (mr *MockDeviceRepoMockRecorder) GetDeviceByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockDeviceRepo)(nil).GetDeviceByUUID), ctx, uuid)
}

// GetDeviceByUUIDWithAmount mocks base method.
func (m *MockDeviceRepo) GetDeviceByUUIDWithAmount(ctx context.Context, deviceUUID string, amount uint32) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUIDWithAmount", ctx, deviceUUID, amount)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUIDWithAmount indicates an expected call of GetDeviceByUUIDWithAmount.
func (mr *MockDeviceRepoMockRecorder) GetDeviceByUUIDWithAmount(ctx, deviceUUID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUIDWithAmount", reflect.TypeOf((*MockDeviceRepo)(nil).GetDeviceByUUIDWithAmount), ctx, deviceUUID, amount)
}

// GetDevicesByManufacturer mocks base method.
func (m *MockDeviceRepo) GetDevicesByManufacturer(ctx context.Context, manu string) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByManufacturer", ctx, manu)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByManufacturer indicates an expected call of GetDevicesByManufacturer.
func (mr *MockDeviceRepoMockRecorder) GetDevicesByManufacturer(ctx, manu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByManufacturer", reflect.TypeOf((*MockDeviceRepo)(nil).GetDevicesByManufacturer), ctx, manu)
}

// GetDevicesByPrice mocks base method.
func (m *MockDeviceRepo) GetDevicesByPrice(ctx context.Context, min, max uint) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByPrice", ctx, min, max)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByPrice indicates an expected call of GetDevicesByPrice.
func (mr *MockDeviceRepoMockRecorder) GetDevicesByPrice(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByPrice", reflect.TypeOf((*MockDeviceRepo)(nil).GetDevicesByPrice), ctx, min, max)
}

// GetDevicesByTitle mocks base method.
func (m *MockDeviceRepo) GetDevicesByTitle(ctx context.Context, title string) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByTitle", ctx, title)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByTitle indicates an expected call of GetDevicesByTitle.
func (mr *MockDeviceRepoMockRecorder) GetDevicesByTitle(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByTitle", reflect.TypeOf((*MockDeviceRepo)(nil).GetDevicesByTitle), ctx, title)
}

// IncreaseDeviceAmountByUUID mocks base method.
func (m *MockDeviceRepo) IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseDeviceAmountByUUID", ctx, deviceUUID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseDeviceAmountByUUID indicates an expected call of IncreaseDeviceAmountByUUID.
func (mr *MockDeviceRepoMockRecorder) IncreaseDeviceAmountByUUID(ctx, deviceUUID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseDeviceAmountByUUID", reflect.TypeOf((*MockDeviceRepo)(nil).IncreaseDeviceAmountByUUID), ctx, deviceUUID, amount)
}

// UpdateDevice mocks base method.
func (m *MockDeviceRepo) UpdateDevice(arg0 context.Context, arg1 models.UpdateDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockDeviceRepoMockRecorder) UpdateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockDeviceRepo)(nil).UpdateDevice), arg0, arg1)
}

// MockGetActions is a mock of GetActions interface.
type MockGetActions struct {
	ctrl     *gomock.Controller
	recorder *MockGetActionsMockRecorder
}

// MockGetActionsMockRecorder is the mock recorder for MockGetActions.
type MockGetActionsMockRecorder struct {
	mock *MockGetActions
}

// NewMockGetActions creates a new mock instance.
func NewMockGetActions(ctrl *gomock.Controller) *MockGetActions {
	mock := &MockGetActions{ctrl: ctrl}
	mock.recorder = &MockGetActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetActions) EXPECT() *MockGetActionsMockRecorder {
	return m.recorder
}

// GetAllDevices mocks base method.
func (m *MockGetActions) GetAllDevices(ctx context.Context, index, amount uint32) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevices", ctx, index, amount)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDevices indicates an expected call of GetAllDevices.
func (mr *MockGetActionsMockRecorder) GetAllDevices(ctx, index, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevices", reflect.TypeOf((*MockGetActions)(nil).GetAllDevices), ctx, index, amount)
}

// GetDeviceByUUID mocks base method.
func (m *MockGetActions) GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUID", ctx, uuid)
	ret0, _ := ret[0].(models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID.
func (mr *MockGetActionsMockRecorder) GetDeviceByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockGetActions)(nil).GetDeviceByUUID), ctx, uuid)
}

// GetDeviceByUUIDWithAmount mocks base method.
func (m *MockGetActions) GetDeviceByUUIDWithAmount(ctx context.Context, deviceUUID string, amount uint32) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUIDWithAmount", ctx, deviceUUID, amount)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUIDWithAmount indicates an expected call of GetDeviceByUUIDWithAmount.
func (mr *MockGetActionsMockRecorder) GetDeviceByUUIDWithAmount(ctx, deviceUUID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUIDWithAmount", reflect.TypeOf((*MockGetActions)(nil).GetDeviceByUUIDWithAmount), ctx, deviceUUID, amount)
}

// GetDevicesByManufacturer mocks base method.
func (m *MockGetActions) GetDevicesByManufacturer(ctx context.Context, manu string) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByManufacturer", ctx, manu)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByManufacturer indicates an expected call of GetDevicesByManufacturer.
func (mr *MockGetActionsMockRecorder) GetDevicesByManufacturer(ctx, manu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByManufacturer", reflect.TypeOf((*MockGetActions)(nil).GetDevicesByManufacturer), ctx, manu)
}

// GetDevicesByPrice mocks base method.
func (m *MockGetActions) GetDevicesByPrice(ctx context.Context, min, max uint) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByPrice", ctx, min, max)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByPrice indicates an expected call of GetDevicesByPrice.
func (mr *MockGetActionsMockRecorder) GetDevicesByPrice(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByPrice", reflect.TypeOf((*MockGetActions)(nil).GetDevicesByPrice), ctx, min, max)
}

// GetDevicesByTitle mocks base method.
func (m *MockGetActions) GetDevicesByTitle(ctx context.Context, title string) ([]*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByTitle", ctx, title)
	ret0, _ := ret[0].([]*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByTitle indicates an expected call of GetDevicesByTitle.
func (mr *MockGetActionsMockRecorder) GetDevicesByTitle(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByTitle", reflect.TypeOf((*MockGetActions)(nil).GetDevicesByTitle), ctx, title)
}

// MockChangeAmountActions is a mock of ChangeAmountActions interface.
type MockChangeAmountActions struct {
	ctrl     *gomock.Controller
	recorder *MockChangeAmountActionsMockRecorder
}

// MockChangeAmountActionsMockRecorder is the mock recorder for MockChangeAmountActions.
type MockChangeAmountActionsMockRecorder struct {
	mock *MockChangeAmountActions
}

// NewMockChangeAmountActions creates a new mock instance.
func NewMockChangeAmountActions(ctrl *gomock.Controller) *MockChangeAmountActions {
	mock := &MockChangeAmountActions{ctrl: ctrl}
	mock.recorder = &MockChangeAmountActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChangeAmountActions) EXPECT() *MockChangeAmountActionsMockRecorder {
	return m.recorder
}

// DecreaseDevicesAmountTx mocks base method.
func (m *MockChangeAmountActions) DecreaseDevicesAmountTx(ctx context.Context, devices []*models.OrderDevice) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseDevicesAmountTx", ctx, devices)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseDevicesAmountTx indicates an expected call of DecreaseDevicesAmountTx.
func (mr *MockChangeAmountActionsMockRecorder) DecreaseDevicesAmountTx(ctx, devices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseDevicesAmountTx", reflect.TypeOf((*MockChangeAmountActions)(nil).DecreaseDevicesAmountTx), ctx, devices)
}

// IncreaseDeviceAmountByUUID mocks base method.
func (m *MockChangeAmountActions) IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseDeviceAmountByUUID", ctx, deviceUUID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseDeviceAmountByUUID indicates an expected call of IncreaseDeviceAmountByUUID.
func (mr *MockChangeAmountActionsMockRecorder) IncreaseDeviceAmountByUUID(ctx, deviceUUID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseDeviceAmountByUUID", reflect.TypeOf((*MockChangeAmountActions)(nil).IncreaseDeviceAmountByUUID), ctx, deviceUUID, amount)
}

// MockAdminActions is a mock of AdminActions interface.
type MockAdminActions struct {
	ctrl     *gomock.Controller
	recorder *MockAdminActionsMockRecorder
}

// MockAdminActionsMockRecorder is the mock recorder for MockAdminActions.
type MockAdminActionsMockRecorder struct {
	mock *MockAdminActions
}

// NewMockAdminActions creates a new mock instance.
func NewMockAdminActions(ctrl *gomock.Controller) *MockAdminActions {
	mock := &MockAdminActions{ctrl: ctrl}
	mock.recorder = &MockAdminActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminActions) EXPECT() *MockAdminActionsMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockAdminActions) CreateDevice(arg0 context.Context, arg1 models.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockAdminActionsMockRecorder) CreateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockAdminActions)(nil).CreateDevice), arg0, arg1)
}

// DeleteDevice mocks base method.
func (m *MockAdminActions) DeleteDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockAdminActionsMockRecorder) DeleteDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockAdminActions)(nil).DeleteDevice), arg0, arg1)
}

// UpdateDevice mocks base method.
func (m *MockAdminActions) UpdateDevice(arg0 context.Context, arg1 models.UpdateDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockAdminActionsMockRecorder) UpdateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockAdminActions)(nil).UpdateDevice), arg0, arg1)
}
