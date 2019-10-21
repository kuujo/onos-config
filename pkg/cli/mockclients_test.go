// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Client Mocks
package cli

import (
	"context"
	"github.com/onosproject/onos-config/pkg/northbound/admin"
	"github.com/onosproject/onos-config/pkg/northbound/diags"
	"github.com/onosproject/onos-config/pkg/types/snapshot/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// MockClientConfig is used by tests to set up which mock clients they want to use
type MockClientsConfig struct {
	registeredModelsClient          *MockConfigAdminServiceListRegisteredModelsClient
	opstateClient                   *MockOpStateDiagsGetOpStateClient
	netChangesClient                *MockConfigAdminServiceGetNetworkChangesClient
	configDiagsClientConfigurations *MockConfigDiagsGetConfigurationsClient
	configDiagsClientChanges        *MockConfigDiagsGetChangesClient
	listDeviceChangesClient         *MockChangeServiceListDeviceChangesClient
	listNetworkChangesClient        *MockChangeServiceListNetworkChangesClient
}

// mockConfigAdminServiceClient is the mock for the ConfigAdminServiceClient
type mockConfigAdminServiceClient struct {
	rollBackID             string
	registeredModelsClient *MockConfigAdminServiceListRegisteredModelsClient
	netChangesClient       *MockConfigAdminServiceGetNetworkChangesClient
}

var LastCreatedClient *mockConfigAdminServiceClient

func (c mockConfigAdminServiceClient) RegisterModel(ctx context.Context, in *admin.RegisterRequest, opts ...grpc.CallOption) (*admin.RegisterResponse, error) {
	response := &admin.RegisterResponse{
		Name:    in.GetSoFile(),
		Version: "1.0",
	}
	return response, nil
}

func (c mockConfigAdminServiceClient) UploadRegisterModel(ctx context.Context, opts ...grpc.CallOption) (admin.ConfigAdminService_UploadRegisterModelClient, error) {
	return nil, nil
}

func (c mockConfigAdminServiceClient) ListRegisteredModels(ctx context.Context, in *admin.ListModelsRequest, opts ...grpc.CallOption) (admin.ConfigAdminService_ListRegisteredModelsClient, error) {
	return c.registeredModelsClient, nil
}

func (c mockConfigAdminServiceClient) GetNetworkChanges(ctx context.Context, in *admin.NetworkChangesRequest, opts ...grpc.CallOption) (admin.ConfigAdminService_GetNetworkChangesClient, error) {
	return c.netChangesClient, nil
}

func (c mockConfigAdminServiceClient) RollbackNetworkChange(ctx context.Context, in *admin.RollbackRequest, opts ...grpc.CallOption) (*admin.RollbackResponse, error) {
	response := &admin.RollbackResponse{
		Message: "Rollback was successful",
	}
	LastCreatedClient.rollBackID = in.Name
	return response, nil
}

func (c mockConfigAdminServiceClient) GetSnapshot(ctx context.Context, in *admin.GetSnapshotRequest, opts ...grpc.CallOption) (*device.Snapshot, error) {
	return nil, nil
}

func (c mockConfigAdminServiceClient) ListSnapshots(ctx context.Context, in *admin.ListSnapshotsRequest, opts ...grpc.CallOption) (admin.ConfigAdminService_ListSnapshotsClient, error) {
	return nil, nil
}

func (c mockConfigAdminServiceClient) CompactChanges(ctx context.Context, in *admin.CompactChangesRequest, opts ...grpc.CallOption) (*admin.CompactChangesResponse, error) {
	return nil, nil
}

// MockConfigAdminServiceListRegisteredModelsClient is a mock of the ConfigAdminServiceListRegisteredModelsClient
// Function pointers are used to allow mocking specific APIs
type MockConfigAdminServiceListRegisteredModelsClient struct {
	recvFn      func() (*admin.ModelInfo, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(interface{}) error
	recvMsgFn   func(interface{}) error
}

func (c MockConfigAdminServiceListRegisteredModelsClient) Recv() (*admin.ModelInfo, error) {
	return c.recvFn()
}

func (c MockConfigAdminServiceListRegisteredModelsClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockConfigAdminServiceListRegisteredModelsClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockConfigAdminServiceListRegisteredModelsClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockConfigAdminServiceListRegisteredModelsClient) Context() context.Context {
	return c.contextFn()
}

func (c MockConfigAdminServiceListRegisteredModelsClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockConfigAdminServiceListRegisteredModelsClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// MockOpStateDiagsGetOpStateClient is a mock of the OpStateDiagsGetOpStateClient
// Function pointers are used to allow mocking specific APIs
type MockOpStateDiagsGetOpStateClient struct {
	recvFn      func() (*diags.OpStateResponse, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(interface{}) error
	recvMsgFn   func(interface{}) error
}

func (c MockOpStateDiagsGetOpStateClient) Recv() (*diags.OpStateResponse, error) {
	return c.recvFn()
}

func (c MockOpStateDiagsGetOpStateClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockOpStateDiagsGetOpStateClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockOpStateDiagsGetOpStateClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockOpStateDiagsGetOpStateClient) Context() context.Context {
	return c.contextFn()
}

func (c MockOpStateDiagsGetOpStateClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockOpStateDiagsGetOpStateClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// mockOpStateDiagsClient is the mock for the OpStateDiagsClient
type mockOpStateDiagsClient struct {
	getOpStateClient diags.OpStateDiags_GetOpStateClient
}

func (m mockOpStateDiagsClient) GetOpState(ctx context.Context, in *diags.OpStateRequest, opts ...grpc.CallOption) (diags.OpStateDiags_GetOpStateClient, error) {
	return m.getOpStateClient, nil
}

// MockConfigDiagsGetConfigurationsClient is a mock of the ConfigDiagsGetOpStateClient
// Function pointers are used to allow mocking specific APIs
type MockConfigDiagsGetConfigurationsClient struct {
	recvFn      func() (*diags.Configuration, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(interface{}) error
	recvMsgFn   func(interface{}) error
}

func (c MockConfigDiagsGetConfigurationsClient) Recv() (*diags.Configuration, error) {
	return c.recvFn()
}

func (c MockConfigDiagsGetConfigurationsClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockConfigDiagsGetConfigurationsClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockConfigDiagsGetConfigurationsClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockConfigDiagsGetConfigurationsClient) Context() context.Context {
	return c.contextFn()
}

func (c MockConfigDiagsGetConfigurationsClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockConfigDiagsGetConfigurationsClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// MockConfigDiagsGetChangesClient is a mock of the ConfigDiagsGetOpStateClient
// Function pointers are used to allow mocking specific APIs
type MockConfigDiagsGetChangesClient struct {
	recvFn      func() (*admin.Change, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(interface{}) error
	recvMsgFn   func(interface{}) error
}

func (c MockConfigDiagsGetChangesClient) Recv() (*admin.Change, error) {
	return c.recvFn()
}

func (c MockConfigDiagsGetChangesClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockConfigDiagsGetChangesClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockConfigDiagsGetChangesClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockConfigDiagsGetChangesClient) Context() context.Context {
	return c.contextFn()
}

func (c MockConfigDiagsGetChangesClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockConfigDiagsGetChangesClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// mockConfigDiagsClient is the mock for the ConfigDiagsClient
type mockConfigDiagsClient struct {
	getConfigDiagsClientConfigurations diags.ConfigDiags_GetConfigurationsClient
	getConfigDiagsClientChanges        diags.ConfigDiags_GetChangesClient
}

func (m mockConfigDiagsClient) GetConfigurations(ctx context.Context, in *diags.ConfigRequest, opts ...grpc.CallOption) (diags.ConfigDiags_GetConfigurationsClient, error) {
	return m.getConfigDiagsClientConfigurations, nil
}

func (m mockConfigDiagsClient) GetChanges(ctx context.Context, in *diags.ChangesRequest, opts ...grpc.CallOption) (diags.ConfigDiags_GetChangesClient, error) {
	return m.getConfigDiagsClientChanges, nil
}

// MockConfigAdminServiceGetNetworkChangesClient is a mock of the ConfigAdminService_GetNetworkChangesClient
// Function pointers are used to allow mocking specific APIs
type MockConfigAdminServiceGetNetworkChangesClient struct {
	recvFn      func() (*admin.NetChange, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(m interface{}) error
	recvMsgFn   func(m interface{}) error
}

func (c MockConfigAdminServiceGetNetworkChangesClient) Recv() (*admin.NetChange, error) {
	return c.recvFn()
}

func (c MockConfigAdminServiceGetNetworkChangesClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockConfigAdminServiceGetNetworkChangesClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockConfigAdminServiceGetNetworkChangesClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockConfigAdminServiceGetNetworkChangesClient) Context() context.Context {
	return c.contextFn()
}

func (c MockConfigAdminServiceGetNetworkChangesClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockConfigAdminServiceGetNetworkChangesClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// MockChangeServiceListDeviceChangesClient is a mock of the ChangeService_ListDeviceChangesClient
// Function pointers are used to allow mocking specific APIs
type MockChangeServiceListDeviceChangesClient struct {
	recvFn      func() (*diags.ListDeviceChangeResponse, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(interface{}) error
	recvMsgFn   func(interface{}) error
}

func (c MockChangeServiceListDeviceChangesClient) Recv() (*diags.ListDeviceChangeResponse, error) {
	return c.recvFn()
}

func (c MockChangeServiceListDeviceChangesClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockChangeServiceListDeviceChangesClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockChangeServiceListDeviceChangesClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockChangeServiceListDeviceChangesClient) Context() context.Context {
	return c.contextFn()
}

func (c MockChangeServiceListDeviceChangesClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockChangeServiceListDeviceChangesClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// MockChangeServiceListNetworkChangesClient is a mock of the ChangeService_ListNetworkChangesClient
// Function pointers are used to allow mocking specific APIs
type MockChangeServiceListNetworkChangesClient struct {
	recvFn      func() (*diags.ListNetworkChangeResponse, error)
	headerFn    func() (metadata.MD, error)
	trailerFn   func() metadata.MD
	closeSendFn func() error
	contextFn   func() context.Context
	sendMsgFn   func(interface{}) error
	recvMsgFn   func(interface{}) error
}

func (c MockChangeServiceListNetworkChangesClient) Recv() (*diags.ListNetworkChangeResponse, error) {
	return c.recvFn()
}

func (c MockChangeServiceListNetworkChangesClient) Header() (metadata.MD, error) {
	return c.headerFn()
}

func (c MockChangeServiceListNetworkChangesClient) Trailer() metadata.MD {
	return c.trailerFn()
}

func (c MockChangeServiceListNetworkChangesClient) CloseSend() error {
	return c.closeSendFn()
}

func (c MockChangeServiceListNetworkChangesClient) Context() context.Context {
	return c.contextFn()
}

func (c MockChangeServiceListNetworkChangesClient) SendMsg(m interface{}) error {
	return c.sendMsgFn(m)
}

func (c MockChangeServiceListNetworkChangesClient) RecvMsg(m interface{}) error {
	return c.recvMsgFn(m)
}

// mockChangeServiceClient is the mock for the ChangeServiceClient
type mockChangeServiceClient struct {
	getChangeServiceClientDeviceChanges  diags.ChangeService_ListDeviceChangesClient
	getChangeServiceClientNetworkChanges diags.ChangeService_ListNetworkChangesClient
}

func (m mockChangeServiceClient) ListNetworkChanges(ctx context.Context, in *diags.ListNetworkChangeRequest, opts ...grpc.CallOption) (diags.ChangeService_ListNetworkChangesClient, error) {
	return m.getChangeServiceClientNetworkChanges, nil
}

func (m mockChangeServiceClient) ListDeviceChanges(ctx context.Context, in *diags.ListDeviceChangeRequest, opts ...grpc.CallOption) (diags.ChangeService_ListDeviceChangesClient, error) {
	return m.getChangeServiceClientDeviceChanges, nil
}

// setUpMockClients sets up factories to create mocks of top level clients used by the CLI
func setUpMockClients(config MockClientsConfig) {
	admin.ConfigAdminClientFactory = func(cc *grpc.ClientConn) admin.ConfigAdminServiceClient {
		LastCreatedClient = &mockConfigAdminServiceClient{
			rollBackID:             "",
			registeredModelsClient: config.registeredModelsClient,
			netChangesClient:       config.netChangesClient,
		}
		return LastCreatedClient
	}
	diags.OpStateDiagsClientFactory = func(cc *grpc.ClientConn) diags.OpStateDiagsClient {
		return mockOpStateDiagsClient{
			getOpStateClient: config.opstateClient,
		}
	}
	diags.ConfigDiagsClientFactory = func(cc *grpc.ClientConn) diags.ConfigDiagsClient {
		return mockConfigDiagsClient{
			getConfigDiagsClientConfigurations: config.configDiagsClientConfigurations,
			getConfigDiagsClientChanges:        config.configDiagsClientChanges,
		}
	}
	diags.ChangeServiceClientFactory = func(cc *grpc.ClientConn) diags.ChangeServiceClient {
		return mockChangeServiceClient{
			getChangeServiceClientDeviceChanges:  config.listDeviceChangesClient,
			getChangeServiceClientNetworkChanges: config.listNetworkChangesClient,
		}
	}
}
