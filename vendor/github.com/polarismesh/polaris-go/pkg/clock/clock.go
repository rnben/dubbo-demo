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

// Package clock provides a global clock.
package clock

import (
	"sync/atomic"
	"time"
)

// 全局时钟
var globalClock *clockImpl

// Clock 时钟接口
type Clock interface {
	// Now 当前集群
	Now() time.Time
}

// clockImpl 时钟的实现
type clockImpl struct {
	currentTime atomic.Value
}

// Now 获取当前时间
func (c *clockImpl) Now() time.Time {
	nowPtr := c.currentTime.Load().(*time.Time)
	return *nowPtr
}

// TimeStep 时间轮的步长
func TimeStep() time.Duration {
	return 10 * time.Millisecond
}

// updateTime 定期更新时间
func (c *clockImpl) updateTime() {
	ticker := time.NewTicker(TimeStep())
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			c.currentTime.Store(&now)
		}
	}
}

// GetClock 获取全局时钟
func GetClock() Clock {
	return globalClock
}

// init 初始化全局时钟
func init() {
	globalClock = &clockImpl{}
	now := time.Now()
	globalClock.currentTime.Store(&now)
	go globalClock.updateTime()
}
