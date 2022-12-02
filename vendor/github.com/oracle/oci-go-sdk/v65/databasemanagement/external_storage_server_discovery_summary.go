// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalStorageServerDiscoverySummary The Exadata storage server.
type ExternalStorageServerDiscoverySummary struct {

	// The name of the entity.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Null for new discover case.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the agent could be used for monitoring.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated connector.
	ConnectorId *string `mandatory:"false" json:"connectorId"`

	// The version of the entity.
	Version *string `mandatory:"false" json:"version"`

	// The internal identifier.
	InternalId *string `mandatory:"false" json:"internalId"`

	// The status of the entity.
	Status *string `mandatory:"false" json:"status"`

	// The error code of the discovery on the resource
	DiscoverErrorCode *string `mandatory:"false" json:"discoverErrorCode"`

	// The error message of the discovery on the resource
	DiscoverErrorMsg *string `mandatory:"false" json:"discoverErrorMsg"`

	// The IP address of the storage server.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The make model of the storage server.
	MakeModel *string `mandatory:"false" json:"makeModel"`

	// The cpu count of the storage server.
	CpuCount *int `mandatory:"false" json:"cpuCount"`

	// The memory size in GB of the storage server.
	MemoryGB *float64 `mandatory:"false" json:"memoryGB"`

	// The connector name of the storage server in rediscovery case.
	ConnectorName *string `mandatory:"false" json:"connectorName"`

	// The status of the entity discover.
	DiscoverStatus EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
}

//GetId returns Id
func (m ExternalStorageServerDiscoverySummary) GetId() *string {
	return m.Id
}

//GetAgentId returns AgentId
func (m ExternalStorageServerDiscoverySummary) GetAgentId() *string {
	return m.AgentId
}

//GetConnectorId returns ConnectorId
func (m ExternalStorageServerDiscoverySummary) GetConnectorId() *string {
	return m.ConnectorId
}

//GetDisplayName returns DisplayName
func (m ExternalStorageServerDiscoverySummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetVersion returns Version
func (m ExternalStorageServerDiscoverySummary) GetVersion() *string {
	return m.Version
}

//GetInternalId returns InternalId
func (m ExternalStorageServerDiscoverySummary) GetInternalId() *string {
	return m.InternalId
}

//GetStatus returns Status
func (m ExternalStorageServerDiscoverySummary) GetStatus() *string {
	return m.Status
}

//GetDiscoverStatus returns DiscoverStatus
func (m ExternalStorageServerDiscoverySummary) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

//GetDiscoverErrorCode returns DiscoverErrorCode
func (m ExternalStorageServerDiscoverySummary) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

//GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m ExternalStorageServerDiscoverySummary) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

func (m ExternalStorageServerDiscoverySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalStorageServerDiscoverySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEntityDiscoveredDiscoverStatusEnum(string(m.DiscoverStatus)); !ok && m.DiscoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoverStatus: %s. Supported values are: %s.", m.DiscoverStatus, strings.Join(GetEntityDiscoveredDiscoverStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalStorageServerDiscoverySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalStorageServerDiscoverySummary ExternalStorageServerDiscoverySummary
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeExternalStorageServerDiscoverySummary
	}{
		"STORAGE_SERVER_DISCOVER_SUMMARY",
		(MarshalTypeExternalStorageServerDiscoverySummary)(m),
	}

	return json.Marshal(&s)
}
