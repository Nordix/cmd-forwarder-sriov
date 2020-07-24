// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chains

import (
	"context"

	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// EmptyDevicePluginServer is a default chain element implementation for the Server interface
type EmptyDevicePluginServer struct{}

// GetDevicePluginOptions stub
func (e EmptyDevicePluginServer) GetDevicePluginOptions(_ context.Context) (*pluginapi.DevicePluginOptions, error) {
	return &pluginapi.DevicePluginOptions{}, nil
}

// ListAndWatch stub
func (e EmptyDevicePluginServer) ListAndWatch(_ chan<- *pluginapi.ListAndWatchResponse) error {
	return nil
}

// Allocate stub
func (e EmptyDevicePluginServer) Allocate(_ context.Context, request *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	return &pluginapi.AllocateResponse{
		ContainerResponses: make([]*pluginapi.ContainerAllocateResponse, len(request.ContainerRequests)),
	}, nil
}

// PreStartContainer stub
func (e EmptyDevicePluginServer) PreStartContainer(_ context.Context, _ *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	return &pluginapi.PreStartContainerResponse{}, nil
}
