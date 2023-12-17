// Code generated by MockGen. DO NOT EDIT.
// Source: proto/gen/device/device_grpc.pb.go

// Package deviceservicemock is a generated GoMock package.
package deviceservicemock

import (
	context "context"
	reflect "reflect"

	device "github.com/alserov/device-shop/proto/gen/device"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockDevicesClient is a mock of DevicesClient interface.
type MockDevicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockDevicesClientMockRecorder
}

// MockDevicesClientMockRecorder is the mock recorder for MockDevicesClient.
type MockDevicesClientMockRecorder struct {
	mock *MockDevicesClient
}

// NewMockDevicesClient creates a new mock instance.
func NewMockDevicesClient(ctrl *gomock.Controller) *MockDevicesClient {
	mock := &MockDevicesClient{ctrl: ctrl}
	mock.recorder = &MockDevicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevicesClient) EXPECT() *MockDevicesClientMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockDevicesClient) CreateDevice(ctx context.Context, in *device.CreateDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDevice", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockDevicesClientMockRecorder) CreateDevice(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockDevicesClient)(nil).CreateDevice), varargs...)
}

// DeleteDevice mocks base method.
func (m *MockDevicesClient) DeleteDevice(ctx context.Context, in *device.DeleteDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDevice", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockDevicesClientMockRecorder) DeleteDevice(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockDevicesClient)(nil).DeleteDevice), varargs...)
}

// GetAllDevices mocks base method.
func (m *MockDevicesClient) GetAllDevices(ctx context.Context, in *device.GetAllDevicesReq, opts ...grpc.CallOption) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllDevices", varargs...)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDevices indicates an expected call of GetAllDevices.
func (mr *MockDevicesClientMockRecorder) GetAllDevices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevices", reflect.TypeOf((*MockDevicesClient)(nil).GetAllDevices), varargs...)
}

// GetDeviceByUUID mocks base method.
func (m *MockDevicesClient) GetDeviceByUUID(ctx context.Context, in *device.GetDeviceByUUIDReq, opts ...grpc.CallOption) (*device.Device, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeviceByUUID", varargs...)
	ret0, _ := ret[0].(*device.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID.
func (mr *MockDevicesClientMockRecorder) GetDeviceByUUID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockDevicesClient)(nil).GetDeviceByUUID), varargs...)
}

// GetDevicesByManufacturer mocks base method.
func (m *MockDevicesClient) GetDevicesByManufacturer(ctx context.Context, in *device.GetByManufacturer, opts ...grpc.CallOption) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDevicesByManufacturer", varargs...)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByManufacturer indicates an expected call of GetDevicesByManufacturer.
func (mr *MockDevicesClientMockRecorder) GetDevicesByManufacturer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByManufacturer", reflect.TypeOf((*MockDevicesClient)(nil).GetDevicesByManufacturer), varargs...)
}

// GetDevicesByPrice mocks base method.
func (m *MockDevicesClient) GetDevicesByPrice(ctx context.Context, in *device.GetByPrice, opts ...grpc.CallOption) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDevicesByPrice", varargs...)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByPrice indicates an expected call of GetDevicesByPrice.
func (mr *MockDevicesClientMockRecorder) GetDevicesByPrice(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByPrice", reflect.TypeOf((*MockDevicesClient)(nil).GetDevicesByPrice), varargs...)
}

// GetDevicesByTitle mocks base method.
func (m *MockDevicesClient) GetDevicesByTitle(ctx context.Context, in *device.GetDeviceByTitleReq, opts ...grpc.CallOption) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDevicesByTitle", varargs...)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByTitle indicates an expected call of GetDevicesByTitle.
func (mr *MockDevicesClientMockRecorder) GetDevicesByTitle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByTitle", reflect.TypeOf((*MockDevicesClient)(nil).GetDevicesByTitle), varargs...)
}

// IncreaseDeviceAmount mocks base method.
func (m *MockDevicesClient) IncreaseDeviceAmount(ctx context.Context, in *device.IncreaseDeviceAmountReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IncreaseDeviceAmount", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncreaseDeviceAmount indicates an expected call of IncreaseDeviceAmount.
func (mr *MockDevicesClientMockRecorder) IncreaseDeviceAmount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseDeviceAmount", reflect.TypeOf((*MockDevicesClient)(nil).IncreaseDeviceAmount), varargs...)
}

// UpdateDevice mocks base method.
func (m *MockDevicesClient) UpdateDevice(ctx context.Context, in *device.UpdateDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDevice", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockDevicesClientMockRecorder) UpdateDevice(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockDevicesClient)(nil).UpdateDevice), varargs...)
}

// MockDevicesServer is a mock of DevicesServer interface.
type MockDevicesServer struct {
	ctrl     *gomock.Controller
	recorder *MockDevicesServerMockRecorder
}

// MockDevicesServerMockRecorder is the mock recorder for MockDevicesServer.
type MockDevicesServerMockRecorder struct {
	mock *MockDevicesServer
}

// NewMockDevicesServer creates a new mock instance.
func NewMockDevicesServer(ctrl *gomock.Controller) *MockDevicesServer {
	mock := &MockDevicesServer{ctrl: ctrl}
	mock.recorder = &MockDevicesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevicesServer) EXPECT() *MockDevicesServerMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockDevicesServer) CreateDevice(arg0 context.Context, arg1 *device.CreateDeviceReq) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockDevicesServerMockRecorder) CreateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockDevicesServer)(nil).CreateDevice), arg0, arg1)
}

// DeleteDevice mocks base method.
func (m *MockDevicesServer) DeleteDevice(arg0 context.Context, arg1 *device.DeleteDeviceReq) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockDevicesServerMockRecorder) DeleteDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockDevicesServer)(nil).DeleteDevice), arg0, arg1)
}

// GetAllDevices mocks base method.
func (m *MockDevicesServer) GetAllDevices(arg0 context.Context, arg1 *device.GetAllDevicesReq) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevices", arg0, arg1)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDevices indicates an expected call of GetAllDevices.
func (mr *MockDevicesServerMockRecorder) GetAllDevices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevices", reflect.TypeOf((*MockDevicesServer)(nil).GetAllDevices), arg0, arg1)
}

// GetDeviceByUUID mocks base method.
func (m *MockDevicesServer) GetDeviceByUUID(arg0 context.Context, arg1 *device.GetDeviceByUUIDReq) (*device.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUID", arg0, arg1)
	ret0, _ := ret[0].(*device.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID.
func (mr *MockDevicesServerMockRecorder) GetDeviceByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockDevicesServer)(nil).GetDeviceByUUID), arg0, arg1)
}

// GetDevicesByManufacturer mocks base method.
func (m *MockDevicesServer) GetDevicesByManufacturer(arg0 context.Context, arg1 *device.GetByManufacturer) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByManufacturer", arg0, arg1)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByManufacturer indicates an expected call of GetDevicesByManufacturer.
func (mr *MockDevicesServerMockRecorder) GetDevicesByManufacturer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByManufacturer", reflect.TypeOf((*MockDevicesServer)(nil).GetDevicesByManufacturer), arg0, arg1)
}

// GetDevicesByPrice mocks base method.
func (m *MockDevicesServer) GetDevicesByPrice(arg0 context.Context, arg1 *device.GetByPrice) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByPrice", arg0, arg1)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByPrice indicates an expected call of GetDevicesByPrice.
func (mr *MockDevicesServerMockRecorder) GetDevicesByPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByPrice", reflect.TypeOf((*MockDevicesServer)(nil).GetDevicesByPrice), arg0, arg1)
}

// GetDevicesByTitle mocks base method.
func (m *MockDevicesServer) GetDevicesByTitle(arg0 context.Context, arg1 *device.GetDeviceByTitleReq) (*device.DevicesRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByTitle", arg0, arg1)
	ret0, _ := ret[0].(*device.DevicesRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesByTitle indicates an expected call of GetDevicesByTitle.
func (mr *MockDevicesServerMockRecorder) GetDevicesByTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByTitle", reflect.TypeOf((*MockDevicesServer)(nil).GetDevicesByTitle), arg0, arg1)
}

// IncreaseDeviceAmount mocks base method.
func (m *MockDevicesServer) IncreaseDeviceAmount(arg0 context.Context, arg1 *device.IncreaseDeviceAmountReq) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseDeviceAmount", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncreaseDeviceAmount indicates an expected call of IncreaseDeviceAmount.
func (mr *MockDevicesServerMockRecorder) IncreaseDeviceAmount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseDeviceAmount", reflect.TypeOf((*MockDevicesServer)(nil).IncreaseDeviceAmount), arg0, arg1)
}

// UpdateDevice mocks base method.
func (m *MockDevicesServer) UpdateDevice(arg0 context.Context, arg1 *device.UpdateDeviceReq) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockDevicesServerMockRecorder) UpdateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockDevicesServer)(nil).UpdateDevice), arg0, arg1)
}

// mustEmbedUnimplementedDevicesServer mocks base method.
func (m *MockDevicesServer) mustEmbedUnimplementedDevicesServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDevicesServer")
}

// mustEmbedUnimplementedDevicesServer indicates an expected call of mustEmbedUnimplementedDevicesServer.
func (mr *MockDevicesServerMockRecorder) mustEmbedUnimplementedDevicesServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDevicesServer", reflect.TypeOf((*MockDevicesServer)(nil).mustEmbedUnimplementedDevicesServer))
}

// MockUnsafeDevicesServer is a mock of UnsafeDevicesServer interface.
type MockUnsafeDevicesServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDevicesServerMockRecorder
}

// MockUnsafeDevicesServerMockRecorder is the mock recorder for MockUnsafeDevicesServer.
type MockUnsafeDevicesServerMockRecorder struct {
	mock *MockUnsafeDevicesServer
}

// NewMockUnsafeDevicesServer creates a new mock instance.
func NewMockUnsafeDevicesServer(ctrl *gomock.Controller) *MockUnsafeDevicesServer {
	mock := &MockUnsafeDevicesServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDevicesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDevicesServer) EXPECT() *MockUnsafeDevicesServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDevicesServer mocks base method.
func (m *MockUnsafeDevicesServer) mustEmbedUnimplementedDevicesServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDevicesServer")
}

// mustEmbedUnimplementedDevicesServer indicates an expected call of mustEmbedUnimplementedDevicesServer.
func (mr *MockUnsafeDevicesServerMockRecorder) mustEmbedUnimplementedDevicesServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDevicesServer", reflect.TypeOf((*MockUnsafeDevicesServer)(nil).mustEmbedUnimplementedDevicesServer))
}
