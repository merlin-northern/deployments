// Copyright 2019 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/deployments/model"
import time "time"

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// AbortDeployment provides a mock function with given fields: ctx, deploymentID
func (_m *App) AbortDeployment(ctx context.Context, deploymentID string) error {
	ret := _m.Called(ctx, deploymentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDeployment provides a mock function with given fields: ctx, constructor
func (_m *App) CreateDeployment(ctx context.Context, constructor *model.DeploymentConstructor) (string, error) {
	ret := _m.Called(ctx, constructor)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.DeploymentConstructor) string); ok {
		r0 = rf(ctx, constructor)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.DeploymentConstructor) error); ok {
		r1 = rf(ctx, constructor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImage provides a mock function with given fields: ctx, multipartUploadMsg
func (_m *App) CreateImage(ctx context.Context, multipartUploadMsg *model.MultipartUploadMsg) (string, error) {
	ret := _m.Called(ctx, multipartUploadMsg)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.MultipartUploadMsg) string); ok {
		r0 = rf(ctx, multipartUploadMsg)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.MultipartUploadMsg) error); ok {
		r1 = rf(ctx, multipartUploadMsg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateImage provides a mock function with given fields, ctx, multipartGenerateImageMsg
func (_m *App) GenerateImage(ctx context.Context, multipartGenerateImageMsg *model.MultipartGenerateImageMsg) (string, error) {
	ret := _m.Called(ctx, multipartGenerateImageMsg)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.MultipartGenerateImageMsg) string); ok {
		r0 = rf(ctx, multipartGenerateImageMsg)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.MultipartGenerateImageMsg) error); ok {
		r1 = rf(ctx, multipartGenerateImageMsg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecommissionDevice provides a mock function with given fields: ctx, deviceID
func (_m *App) DecommissionDevice(ctx context.Context, deviceID string) error {
	ret := _m.Called(ctx, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteImage provides a mock function with given fields: ctx, imageID
func (_m *App) DeleteImage(ctx context.Context, imageID string) error {
	ret := _m.Called(ctx, imageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, imageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DownloadLink provides a mock function with given fields: ctx, imageID, expire
func (_m *App) DownloadLink(ctx context.Context, imageID string, expire time.Duration) (*model.Link, error) {
	ret := _m.Called(ctx, imageID, expire)

	var r0 *model.Link
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) *model.Link); ok {
		r0 = rf(ctx, imageID, expire)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) error); ok {
		r1 = rf(ctx, imageID, expire)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditImage provides a mock function with given fields: ctx, id, constructorData
func (_m *App) EditImage(ctx context.Context, id string, constructorData *model.ImageMeta) (bool, error) {
	ret := _m.Called(ctx, id, constructorData)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.ImageMeta) bool); ok {
		r0 = rf(ctx, id, constructorData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *model.ImageMeta) error); ok {
		r1 = rf(ctx, id, constructorData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployment provides a mock function with given fields: ctx, deploymentID
func (_m *App) GetDeployment(ctx context.Context, deploymentID string) (*model.Deployment, error) {
	ret := _m.Called(ctx, deploymentID)

	var r0 *model.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Deployment); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentForDeviceWithCurrent provides a mock function with given fields: ctx, deviceID, current
func (_m *App) GetDeploymentForDeviceWithCurrent(ctx context.Context, deviceID string, current *model.InstalledDeviceDeployment) (*model.DeploymentInstructions, error) {
	ret := _m.Called(ctx, deviceID, current)

	var r0 *model.DeploymentInstructions
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.InstalledDeviceDeployment) *model.DeploymentInstructions); ok {
		r0 = rf(ctx, deviceID, current)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeploymentInstructions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *model.InstalledDeviceDeployment) error); ok {
		r1 = rf(ctx, deviceID, current)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentStats provides a mock function with given fields: ctx, deploymentID
func (_m *App) GetDeploymentStats(ctx context.Context, deploymentID string) (model.Stats, error) {
	ret := _m.Called(ctx, deploymentID)

	var r0 model.Stats
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Stats); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceDeploymentLog provides a mock function with given fields: ctx, deviceID, deploymentID
func (_m *App) GetDeviceDeploymentLog(ctx context.Context, deviceID string, deploymentID string) (*model.DeploymentLog, error) {
	ret := _m.Called(ctx, deviceID, deploymentID)

	var r0 *model.DeploymentLog
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.DeploymentLog); ok {
		r0 = rf(ctx, deviceID, deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeploymentLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, deviceID, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceStatusesForDeployment provides a mock function with given fields: ctx, deploymentID
func (_m *App) GetDeviceStatusesForDeployment(ctx context.Context, deploymentID string) ([]model.DeviceDeployment, error) {
	ret := _m.Called(ctx, deploymentID)

	var r0 []model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.DeviceDeployment); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImage provides a mock function with given fields: ctx, id
func (_m *App) GetImage(ctx context.Context, id string) (*model.Image, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Image
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Image); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLimit provides a mock function with given fields: ctx, name
func (_m *App) GetLimit(ctx context.Context, name string) (*model.Limit, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.Limit
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Limit); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Limit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasDeploymentForDevice provides a mock function with given fields: ctx, deploymentID, deviceID
func (_m *App) HasDeploymentForDevice(ctx context.Context, deploymentID string, deviceID string) (bool, error) {
	ret := _m.Called(ctx, deploymentID, deviceID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, deploymentID, deviceID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, deploymentID, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDeploymentFinished provides a mock function with given fields: ctx, deploymentID
func (_m *App) IsDeploymentFinished(ctx context.Context, deploymentID string) (bool, error) {
	ret := _m.Called(ctx, deploymentID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImages provides a mock function with given fields: ctx, filters
func (_m *App) ListImages(ctx context.Context, filters map[string]string) ([]*model.Image, error) {
	ret := _m.Called(ctx, filters)

	var r0 []*model.Image
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) []*model.Image); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupDeployment provides a mock function with given fields: ctx, query
func (_m *App) LookupDeployment(ctx context.Context, query model.Query) ([]*model.Deployment, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, model.Query) []*model.Deployment); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProvisionTenant provides a mock function with given fields: ctx, tenant_id
func (_m *App) ProvisionTenant(ctx context.Context, tenant_id string) error {
	ret := _m.Called(ctx, tenant_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenant_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDeviceDeploymentLog provides a mock function with given fields: ctx, deviceID, deploymentID, logs
func (_m *App) SaveDeviceDeploymentLog(ctx context.Context, deviceID string, deploymentID string, logs []model.LogMessage) error {
	ret := _m.Called(ctx, deviceID, deploymentID, logs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []model.LogMessage) error); ok {
		r0 = rf(ctx, deviceID, deploymentID, logs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceDeploymentStatus provides a mock function with given fields: ctx, deploymentID, deviceID, status
func (_m *App) UpdateDeviceDeploymentStatus(ctx context.Context, deploymentID string, deviceID string, status model.DeviceDeploymentStatus) error {
	ret := _m.Called(ctx, deploymentID, deviceID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.DeviceDeploymentStatus) error); ok {
		r0 = rf(ctx, deploymentID, deviceID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
