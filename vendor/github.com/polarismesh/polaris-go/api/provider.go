/**
 * Tencent is pleased to support the open source community by making polaris-go available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package api

import (
	"github.com/polarismesh/polaris-go/pkg/model"
)

// InstanceHeartbeatRequest 心跳上报请求
type InstanceHeartbeatRequest struct {
	model.InstanceHeartbeatRequest
}

// InstanceDeRegisterRequest 反注册服务请求
type InstanceDeRegisterRequest struct {
	model.InstanceDeRegisterRequest
}

// InstanceRegisterRequest 注册服务请求
type InstanceRegisterRequest struct {
	model.InstanceRegisterRequest
}

// ProviderAPI CL5服务端API的主接口
type ProviderAPI interface {
	SDKOwner
	// RegisterInstance
	// minimum supported version of polaris-server is v1.10.0
	RegisterInstance(instance *InstanceRegisterRequest) (*model.InstanceRegisterResponse, error)
	// Register
	// 同步注册服务，服务注册成功后会填充instance中的InstanceID字段
	// 用户可保持该instance对象用于反注册和心跳上报
	// Deprecated: Use RegisterInstance instead.
	Register(instance *InstanceRegisterRequest) (*model.InstanceRegisterResponse, error)
	// Deregister synchronize the anti registration service
	Deregister(instance *InstanceDeRegisterRequest) error
	// Heartbeat the heartbeat report
	// Deprecated: Use RegisterInstance instead.
	Heartbeat(instance *InstanceHeartbeatRequest) error
	// Destroy the api is destroyed and cannot be called again
	Destroy()
}

var (
	// NewProviderAPI 通过以默认域名为埋点server的默认配置创建ProviderAPI
	NewProviderAPI = newProviderAPI
	// NewProviderAPIByFile 通过配置文件创建SDK ProviderAPI对象
	NewProviderAPIByFile = newProviderAPIByFile
	// NewProviderAPIByConfig 通过配置对象创建SDK ProviderAPI对象
	NewProviderAPIByConfig = newProviderAPIByConfig
	// NewProviderAPIByContext 通过上下文创建SDK ProviderAPI对象
	NewProviderAPIByContext = newProviderAPIByContext
	// NewProviderAPIByDefaultConfigFile 通过系统默认配置文件创建ProviderAPI
	NewProviderAPIByDefaultConfigFile = newProviderAPIByDefaultConfigFile
	// NewProviderAPIByAddress 通过address创建ProviderAPI对象
	NewProviderAPIByAddress = newProviderAPIByAddress
)
