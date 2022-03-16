// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// StreamingServiceAction An action that delivers to an Oracle Stream Service stream.
type StreamingServiceAction struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the action.
	Id *string `mandatory:"true" json:"id"`

	// A message generated by the Events service about the current state of this action.
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream to which messages are delivered.
	StreamId *string `mandatory:"false" json:"streamId"`

	// The current state of the rule.
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m StreamingServiceAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m StreamingServiceAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m StreamingServiceAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

//GetIsEnabled returns IsEnabled
func (m StreamingServiceAction) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m StreamingServiceAction) GetDescription() *string {
	return m.Description
}

func (m StreamingServiceAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamingServiceAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingActionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetActionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StreamingServiceAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamingServiceAction StreamingServiceAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeStreamingServiceAction
	}{
		"OSS",
		(MarshalTypeStreamingServiceAction)(m),
	}

	return json.Marshal(&s)
}