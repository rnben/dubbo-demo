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

package config

import (
	"github.com/polarismesh/polaris-go/pkg/plugin/common"
)

// StatReporterConfigImpl global.statReporter.
type StatReporterConfigImpl struct {
	// 是否启动上报
	Enable *bool `yaml:"enable" json:"enable"`
	// 上报插件链
	Chain []string `yaml:"chain" json:"chain"`
	// 插件相关配置
	Plugin PluginConfigs `yaml:"plugin" json:"plugin"`
}

// IsEnable 是否启用上报.
func (s *StatReporterConfigImpl) IsEnable() bool {
	return *s.Enable
}

// SetEnable 设置是否启用上报.
func (s *StatReporterConfigImpl) SetEnable(enable bool) {
	s.Enable = &enable
}

// GetChain 插件链条.
func (s *StatReporterConfigImpl) GetChain() []string {
	return s.Chain
}

// SetChain 设置插件链条.
func (s *StatReporterConfigImpl) SetChain(chain []string) {
	s.Chain = chain
}

// GetPluginConfig 获取一个插件的配置.
func (s *StatReporterConfigImpl) GetPluginConfig(name string) BaseConfig {
	value, ok := s.Plugin[name]
	if !ok {
		return nil
	}
	return value.(BaseConfig)
}

// Verify 检测statReporter配置.
func (s *StatReporterConfigImpl) Verify() error {
	return s.Plugin.Verify()
}

// SetDefault 设置statReporter默认值.
func (s *StatReporterConfigImpl) SetDefault() {
	if nil == s.Enable {
		enable := DefaultStatReportEnabled
		s.Enable = &enable
	}
	if len(s.Chain) == 0 {
		s.Chain = []string{}
	}
	s.Plugin.SetDefault(common.TypeStatReporter)
}

// Init 配置初始化.
func (s *StatReporterConfigImpl) Init() {
	s.Plugin = PluginConfigs{}
	s.Plugin.Init(common.TypeStatReporter)
}

// SetPluginConfig 输出插件具体配置.
func (s *StatReporterConfigImpl) SetPluginConfig(plugName string, value BaseConfig) error {
	return s.Plugin.SetPluginConfig(common.TypeStatReporter, plugName, value)
}
