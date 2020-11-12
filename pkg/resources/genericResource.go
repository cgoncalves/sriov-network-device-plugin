// Copyright 2018 Intel Corp. All Rights Reserved.
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

package resources

import (
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type genericResource struct {
}

// NewGenericResource instantiate a generic DeviceInfoProvider
func NewGenericResource() types.DeviceInfoProvider {
	return &genericResource{}
}

func (rp *genericResource) GetDeviceSpecs(pciAddr string) []*pluginapi.DeviceSpec {
	devSpecs := make([]*pluginapi.DeviceSpec, 0)
	// NO device file, send empty DeviceSpec map
	return devSpecs
}

func (rp *genericResource) GetEnvVal(pciAddr string) string {
	return pciAddr
}

func (rp *genericResource) GetMounts(pciAddr string) []*pluginapi.Mount {
	mounts := make([]*pluginapi.Mount, 0)
	return mounts
}
