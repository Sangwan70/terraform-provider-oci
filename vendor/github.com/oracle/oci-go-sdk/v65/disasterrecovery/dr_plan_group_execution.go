// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanGroupExecution Summary information about a group execution in a DR Plan Execution.
type DrPlanGroupExecution struct {

	// The unique id of the group. Must not be modified by user.
	// Example: `sgid1.group..examplegroupsgid`
	GroupId *string `mandatory:"true" json:"groupId"`

	// The plan group type.
	Type DrPlanGroupTypeEnum `mandatory:"true" json:"type"`

	// The display name of group that was executed.
	// Example: `DATABASE_SWITCHOVER`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The status of the group execution.
	Status DrPlanGroupExecutionStatusEnum `mandatory:"true" json:"status"`

	// A list of details of each step executed in this group.
	StepExecutions []DrPlanStepExecution `mandatory:"true" json:"stepExecutions"`

	// Additional details about the group execution status.
	// Example: `A total of three steps failed in the group`
	StatusDetails *string `mandatory:"false" json:"statusDetails"`

	// The time at which group execution began. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time at which group execution ended. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The total duration in seconds taken to complete group execution.
	// Example: `120`
	ExecutionDurationInSec *int `mandatory:"false" json:"executionDurationInSec"`
}

func (m DrPlanGroupExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanGroupExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrPlanGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrPlanGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrPlanGroupExecutionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDrPlanGroupExecutionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
