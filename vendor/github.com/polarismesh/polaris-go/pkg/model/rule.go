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

package model

import (
	"fmt"
	"sync"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/golang/protobuf/proto"
)

// 网格规则类型.
const (
	// MeshServiceEntry ServiceEntry.
	MeshServiceEntry string = "networking.istio.io/v1alpha3/ServiceEntry"
	// MeshVirtualService VirtualService.
	MeshVirtualService string = "networking.istio.io/v1alpha3/VirtualService"
	// MeshDestinationRule DestinationRule.
	MeshDestinationRule string = "networking.istio.io/v1alpha3/DestinationRule"
	// MeshEnvoyFilter EnvoyFilter.
	MeshEnvoyFilter string = "networking.istio.io/v1alpha3/EnvoyFilter"
	// MeshGateway Gateway.
	MeshGateway string = "networking.istio.io/v1alpha3/Gateway"
)

var (
	// MeshPrefix 网格请求key前缀.
	MeshPrefix = "mesh_resource"
	// MeshKeyLen 网格请求key分割长度.
	MeshKeyLen = 3
	// MeshKeySpliter 网格请求key分隔符.
	MeshKeySpliter = "MESHa071a34fecSPLITER"
)

// ServiceRule 服务配置通用接口.
type ServiceRule interface {
	RegistryValue
	// 获取路由所属服务的命名空间
	GetNamespace() string
	// 获取路由所属服务的服务名
	GetService() string
	// 获取值
	// PB场景下，路由规则类型为*Routing
	GetValue() interface{}
	// 获取规则缓存信息
	GetRuleCache() RuleCache
	// 获取规则校验失败异常
	GetValidateError() error

	IsCacheLoaded() bool
}

// RuleCache 服务规则缓存.
type RuleCache interface {
	// 通过字面值获取表达式对象
	GetRegexMatcher(message string) (*regexp.Regexp, error)
	// 获取消息缓存
	GetMessageCache(message proto.Message) interface{}
	// 设置消息缓存
	SetMessageCache(message proto.Message, cacheValue interface{})
}

// NewRuleCache 创建规则缓存对象.
func NewRuleCache() RuleCache {
	return &ruleCache{
		regexMatchers: make(map[string]*regexp.Regexp),
		messageCaches: make(map[proto.Message]interface{}),
	}
}

// ruleCache 路由规则缓存实现.
type ruleCache struct {
	mutex         sync.RWMutex
	regexMatchers map[string]*regexp.Regexp
	messageCaches map[proto.Message]interface{}
}

// GetRegexMatcher 通过字面值获取表达式对象.
func (r *ruleCache) GetRegexMatcher(message string) (*regexp.Regexp, error) {
	r.mutex.RLock()
	regexObj, ok := r.regexMatchers[message]
	r.mutex.RUnlock()
	if ok {
		return regexObj, nil
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	regexObj, ok = r.regexMatchers[message]
	if ok {
		return regexObj, nil
	}
	var err error
	regexObj, err = regexp.Compile(message, regexp.RE2)
	if err != nil {
		return nil, fmt.Errorf("invalid regex expression %s, error is %v", message, err)
	}
	r.regexMatchers[message] = regexObj
	return regexObj, nil
}

// GetMessageCache 获取hash值.
func (r *ruleCache) GetMessageCache(message proto.Message) interface{} {
	return r.messageCaches[message]
}

// SetMessageCache 设置hash值.
func (r *ruleCache) SetMessageCache(message proto.Message, cacheValue interface{}) {
	r.messageCaches[message] = cacheValue
}

// GetServiceRuleRequest 获取服务规则请求.
type GetServiceRuleRequest struct {
	// 可选，流水号，用于跟踪用户的请求，默认0
	FlowID uint64
	// 命名空间
	Namespace string
	// 服务名
	Service string
	// 可选，仅对同步请求有效，本次查询最大超时信息，可选，默认直接获取全局的超时配置
	Timeout *time.Duration
	// 可选，重试次数，默认直接获取全局的超时配置
	RetryCount *int
	// 应答对象，由主流程填充并返回
	response ServiceRuleResponse
}

// GetService 获取服务名.
func (g *GetServiceRuleRequest) GetService() string {
	return g.Service
}

// GetNamespace 获取命名空间.
func (g *GetServiceRuleRequest) GetNamespace() string {
	return g.Namespace
}

// GetMetadata 获取元数据信息.
func (g *GetServiceRuleRequest) GetMetadata() map[string]string {
	return nil
}

// SetTimeout 设置超时时间.
func (g *GetServiceRuleRequest) SetTimeout(duration time.Duration) {
	g.Timeout = ToDurationPtr(duration)
}

// SetRetryCount 设置重试次数.
func (g *GetServiceRuleRequest) SetRetryCount(retryCount int) {
	g.RetryCount = &retryCount
}

// GetResponse 获取应答.
func (g *GetServiceRuleRequest) GetResponse() *ServiceRuleResponse {
	return &g.response
}

// GetTimeoutPtr 获取超时值指针.
func (g *GetServiceRuleRequest) GetTimeoutPtr() *time.Duration {
	return g.Timeout
}

// GetRetryCountPtr 获取重试次数指针.
func (g *GetServiceRuleRequest) GetRetryCountPtr() *int {
	return g.RetryCount
}

// Validate 校验获取服务规则请求对象.
func (g *GetServiceRuleRequest) Validate() error {
	if nil == g {
		return NewSDKError(ErrCodeAPIInvalidArgument, nil, "GetServiceRuleRequest can not be nil")
	}
	if err := validateServiceMetadata("GetServiceRuleRequest", g); err != nil {
		return NewSDKError(ErrCodeAPIInvalidArgument, err, "fail to validate GetServiceRuleRequest")
	}
	return nil
}

// ServiceRuleResponse 服务规则应答.
type ServiceRuleResponse struct {
	// 规则类型
	Type EventType
	// 所属服务
	Service ServiceKey
	// 规则对象，不同EventType对应不同类型实例
	Value interface{}
	// 规则版本信息
	Revision string
	// 规则缓存
	RuleCache RuleCache
	// 规则校验异常
	ValidateError error
}

// GetType 获取配置类型.
func (s *ServiceRuleResponse) GetType() EventType {
	return s.Type
}

// GetValue 获取值
// PB场景下，路由规则类型为*Routing.
func (s *ServiceRuleResponse) GetValue() interface{} {
	return s.Value
}

// IsInitialized 配置规则是否已经加载.
func (s *ServiceRuleResponse) IsInitialized() bool {
	return true
}

func (s *ServiceRuleResponse) IsCacheLoaded() bool {
	return false
}

// GetRevision 获取配置规则的修订版本信息.
func (s *ServiceRuleResponse) GetRevision() string {
	return s.Revision
}

// GetRuleCache 获取规则缓存信息.
func (s *ServiceRuleResponse) GetRuleCache() RuleCache {
	return s.RuleCache
}

// GetNamespace 获取命名空间.
func (s *ServiceRuleResponse) GetNamespace() string {
	return s.Service.Namespace
}

// GetService 获取服务名.
func (s *ServiceRuleResponse) GetService() string {
	return s.Service.Service
}

// GetValidateError 获取规则校验异常.
func (s *ServiceRuleResponse) GetValidateError() error {
	return s.ValidateError
}
