/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package report

import (
	gxset "github.com/dubbogo/gost/container/set"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/metadata/identifier"
	"dubbo.apache.org/dubbo-go/v3/registry"
)

// MetadataReport is an interface of remote metadata report.
type MetadataReport interface {
	// StoreProviderMetadata stores the metadata.
	// Metadata includes the basic info of the server,
	// provider info, and other user custom info.
	StoreProviderMetadata(*identifier.MetadataIdentifier, string) error

	// StoreConsumerMetadata stores the metadata.
	// Metadata includes the basic info of the server,
	// consumer info, and other user custom info.
	StoreConsumerMetadata(*identifier.MetadataIdentifier, string) error

	// SaveServiceMetadata saves the metadata.
	// Metadata includes the basic info of the server,
	// service info, and other user custom info.
	SaveServiceMetadata(*identifier.ServiceMetadataIdentifier, *common.URL) error

	// RemoveServiceMetadata removes the metadata.
	RemoveServiceMetadata(*identifier.ServiceMetadataIdentifier) error

	// GetExportedURLs gets the urls.
	// If not found, an empty list will be returned.
	GetExportedURLs(*identifier.ServiceMetadataIdentifier) ([]string, error)

	// SaveSubscribedData saves the urls.
	// If not found, an empty str will be returned.
	SaveSubscribedData(*identifier.SubscriberMetadataIdentifier, string) error

	// GetSubscribedURLs gets the urls.
	// If not found, an empty list will be returned.
	GetSubscribedURLs(*identifier.SubscriberMetadataIdentifier) ([]string, error)

	// GetServiceDefinition gets the service definition.
	GetServiceDefinition(*identifier.MetadataIdentifier) (string, error)

	// GetAppMetadata get metadata info from report
	GetAppMetadata(*identifier.SubscriberMetadataIdentifier) (*common.MetadataInfo, error)

	// PublishAppMetadata publish metadata info to reportss
	PublishAppMetadata(*identifier.SubscriberMetadataIdentifier, *common.MetadataInfo) error

	// RegisterServiceAppMapping map the specified Dubbo service interface to current Dubbo app name
	RegisterServiceAppMapping(string, string, string) error

	// GetServiceAppMapping get the app names from the specified Dubbo service interface
	GetServiceAppMapping(string, string, registry.MappingListener) (*gxset.HashSet, error)

	// RemoveServiceAppMappingListener remove the serviceMapping listener by key and group
	RemoveServiceAppMappingListener(string, string) error
}
