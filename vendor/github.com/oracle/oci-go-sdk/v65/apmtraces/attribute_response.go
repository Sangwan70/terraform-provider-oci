// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttributeResponse Response of an individual attribute item in the bulk activate or de-activate operation.
type AttributeResponse struct {

	// Attribute that was activated or de-activated by this bulk operation.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Type of the attribute.
	AttributeType AttributeResponseAttributeTypeEnum `mandatory:"true" json:"attributeType"`

	// Type of operation - activate or de-activate.
	OperationType AttributeResponseOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the attribute after this operation.  The attribute can have one of the following statuses after the activate or de-activate operation.  The attribute
	// can have either a success status or an error status.  The status of the attribute must be correlated with the operation status property in the bulk operation metadata
	// object.  The bulk operation will be successful only when all attributes in the bulk request are processed successful and they get a success status back.
	// The following are successful status values of individual attribute items in a bulk attribute activation operation.
	// ATTRIBUTE_ACTIVATED - The attribute is activated and is available to be queried.  Note that the span processor might still have not picked up the changes, and the
	// associated caches would not have refreshed yet to pick up the changes.
	// ATTRIBUTE_ALREADY_ACTIVE - The caller is trying to activate an attribute that is already active or in the process of getting activated.
	// ATTRIBUTE_DEACTIVATED - The attribute is de-activated and will not appear in searches.  The span processor might not have picked up the new changes and the associated caches
	// might not have refreshed yet.
	// ATTRIBUTE_ALREADY_DEACTIVATED - The caller is trying to de-activate an attribute that has already been de-activated or in the process of de-activation.
	// DUPLICATE_ATTRIBUTE - The attriubute is a duplicate of an attribute that was present in this bulk request.  Note that we de-duplicate the attribute collection, process only unique attributes,
	// and call out duplicates.  A duplicate attribute in a bulk request will not prevent the processing of further attributes in the bulk operation.
	// The following values are error statuses and the bulk processing is stopped when the first error is encountered.  None of the attributes in the bulk request would have been activated or
	// de-activated by this bulk operation.
	// DEACTIVATION_NOT_ALLOWED - The caller has asked for the de-activation of an out of box tag which is not permitted.
	// ATTRIBUTE_DOES_NOT_EXIST - The caller tries to de-activate an attribute that doesn't exist in the APM Domain.
	// INVALID_ATTRIBUTE - The attribute is invalid.
	// INVALID_ATTRIBUTE_TYPE_CONFLICT - The attribute is invalid.  There were two attributes with same name but different type in the bulk request.
	// ATTRIBUTE_NOT_PROCESSED - The attribute was not processed, as there was another attribute in this bulk request collection that resulted in a processing error.
	AttributeStatus AttributeResponseAttributeStatusEnum `mandatory:"true" json:"attributeStatus"`

	// Time when the attribute was activated or de-activated.  Note that the span processor might not have picked up the changes even if this time has elapsed.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Notes for the pinned attribute.
	Notes *string `mandatory:"false" json:"notes"`
}

func (m AttributeResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeResponseAttributeTypeEnum(string(m.AttributeType)); !ok && m.AttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeType: %s. Supported values are: %s.", m.AttributeType, strings.Join(GetAttributeResponseAttributeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeResponseOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetAttributeResponseOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeResponseAttributeStatusEnum(string(m.AttributeStatus)); !ok && m.AttributeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeStatus: %s. Supported values are: %s.", m.AttributeStatus, strings.Join(GetAttributeResponseAttributeStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeResponseAttributeTypeEnum Enum with underlying type: string
type AttributeResponseAttributeTypeEnum string

// Set of constants representing the allowable values for AttributeResponseAttributeTypeEnum
const (
	AttributeResponseAttributeTypeNumeric AttributeResponseAttributeTypeEnum = "NUMERIC"
	AttributeResponseAttributeTypeString  AttributeResponseAttributeTypeEnum = "STRING"
)

var mappingAttributeResponseAttributeTypeEnum = map[string]AttributeResponseAttributeTypeEnum{
	"NUMERIC": AttributeResponseAttributeTypeNumeric,
	"STRING":  AttributeResponseAttributeTypeString,
}

var mappingAttributeResponseAttributeTypeEnumLowerCase = map[string]AttributeResponseAttributeTypeEnum{
	"numeric": AttributeResponseAttributeTypeNumeric,
	"string":  AttributeResponseAttributeTypeString,
}

// GetAttributeResponseAttributeTypeEnumValues Enumerates the set of values for AttributeResponseAttributeTypeEnum
func GetAttributeResponseAttributeTypeEnumValues() []AttributeResponseAttributeTypeEnum {
	values := make([]AttributeResponseAttributeTypeEnum, 0)
	for _, v := range mappingAttributeResponseAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeResponseAttributeTypeEnumStringValues Enumerates the set of values in String for AttributeResponseAttributeTypeEnum
func GetAttributeResponseAttributeTypeEnumStringValues() []string {
	return []string{
		"NUMERIC",
		"STRING",
	}
}

// GetMappingAttributeResponseAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeResponseAttributeTypeEnum(val string) (AttributeResponseAttributeTypeEnum, bool) {
	enum, ok := mappingAttributeResponseAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeResponseOperationTypeEnum Enum with underlying type: string
type AttributeResponseOperationTypeEnum string

// Set of constants representing the allowable values for AttributeResponseOperationTypeEnum
const (
	AttributeResponseOperationTypeActivate   AttributeResponseOperationTypeEnum = "ACTIVATE"
	AttributeResponseOperationTypeDeactivate AttributeResponseOperationTypeEnum = "DEACTIVATE"
)

var mappingAttributeResponseOperationTypeEnum = map[string]AttributeResponseOperationTypeEnum{
	"ACTIVATE":   AttributeResponseOperationTypeActivate,
	"DEACTIVATE": AttributeResponseOperationTypeDeactivate,
}

var mappingAttributeResponseOperationTypeEnumLowerCase = map[string]AttributeResponseOperationTypeEnum{
	"activate":   AttributeResponseOperationTypeActivate,
	"deactivate": AttributeResponseOperationTypeDeactivate,
}

// GetAttributeResponseOperationTypeEnumValues Enumerates the set of values for AttributeResponseOperationTypeEnum
func GetAttributeResponseOperationTypeEnumValues() []AttributeResponseOperationTypeEnum {
	values := make([]AttributeResponseOperationTypeEnum, 0)
	for _, v := range mappingAttributeResponseOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeResponseOperationTypeEnumStringValues Enumerates the set of values in String for AttributeResponseOperationTypeEnum
func GetAttributeResponseOperationTypeEnumStringValues() []string {
	return []string{
		"ACTIVATE",
		"DEACTIVATE",
	}
}

// GetMappingAttributeResponseOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeResponseOperationTypeEnum(val string) (AttributeResponseOperationTypeEnum, bool) {
	enum, ok := mappingAttributeResponseOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeResponseAttributeStatusEnum Enum with underlying type: string
type AttributeResponseAttributeStatusEnum string

// Set of constants representing the allowable values for AttributeResponseAttributeStatusEnum
const (
	AttributeResponseAttributeStatusAttributeAlreadyActive       AttributeResponseAttributeStatusEnum = "ATTRIBUTE_ALREADY_ACTIVE"
	AttributeResponseAttributeStatusAttributeActivated           AttributeResponseAttributeStatusEnum = "ATTRIBUTE_ACTIVATED"
	AttributeResponseAttributeStatusAttributeDeactivated         AttributeResponseAttributeStatusEnum = "ATTRIBUTE_DEACTIVATED"
	AttributeResponseAttributeStatusDeactivationNotAllowed       AttributeResponseAttributeStatusEnum = "DEACTIVATION_NOT_ALLOWED"
	AttributeResponseAttributeStatusAttributeDoesNotExist        AttributeResponseAttributeStatusEnum = "ATTRIBUTE_DOES_NOT_EXIST"
	AttributeResponseAttributeStatusAttributeAlreadyDeactivated  AttributeResponseAttributeStatusEnum = "ATTRIBUTE_ALREADY_DEACTIVATED"
	AttributeResponseAttributeStatusDuplicateAttribute           AttributeResponseAttributeStatusEnum = "DUPLICATE_ATTRIBUTE"
	AttributeResponseAttributeStatusInvalidAttribute             AttributeResponseAttributeStatusEnum = "INVALID_ATTRIBUTE"
	AttributeResponseAttributeStatusInvalidAttributeTypeConflict AttributeResponseAttributeStatusEnum = "INVALID_ATTRIBUTE_TYPE_CONFLICT"
	AttributeResponseAttributeStatusAttributeNotProcessed        AttributeResponseAttributeStatusEnum = "ATTRIBUTE_NOT_PROCESSED"
)

var mappingAttributeResponseAttributeStatusEnum = map[string]AttributeResponseAttributeStatusEnum{
	"ATTRIBUTE_ALREADY_ACTIVE":        AttributeResponseAttributeStatusAttributeAlreadyActive,
	"ATTRIBUTE_ACTIVATED":             AttributeResponseAttributeStatusAttributeActivated,
	"ATTRIBUTE_DEACTIVATED":           AttributeResponseAttributeStatusAttributeDeactivated,
	"DEACTIVATION_NOT_ALLOWED":        AttributeResponseAttributeStatusDeactivationNotAllowed,
	"ATTRIBUTE_DOES_NOT_EXIST":        AttributeResponseAttributeStatusAttributeDoesNotExist,
	"ATTRIBUTE_ALREADY_DEACTIVATED":   AttributeResponseAttributeStatusAttributeAlreadyDeactivated,
	"DUPLICATE_ATTRIBUTE":             AttributeResponseAttributeStatusDuplicateAttribute,
	"INVALID_ATTRIBUTE":               AttributeResponseAttributeStatusInvalidAttribute,
	"INVALID_ATTRIBUTE_TYPE_CONFLICT": AttributeResponseAttributeStatusInvalidAttributeTypeConflict,
	"ATTRIBUTE_NOT_PROCESSED":         AttributeResponseAttributeStatusAttributeNotProcessed,
}

var mappingAttributeResponseAttributeStatusEnumLowerCase = map[string]AttributeResponseAttributeStatusEnum{
	"attribute_already_active":        AttributeResponseAttributeStatusAttributeAlreadyActive,
	"attribute_activated":             AttributeResponseAttributeStatusAttributeActivated,
	"attribute_deactivated":           AttributeResponseAttributeStatusAttributeDeactivated,
	"deactivation_not_allowed":        AttributeResponseAttributeStatusDeactivationNotAllowed,
	"attribute_does_not_exist":        AttributeResponseAttributeStatusAttributeDoesNotExist,
	"attribute_already_deactivated":   AttributeResponseAttributeStatusAttributeAlreadyDeactivated,
	"duplicate_attribute":             AttributeResponseAttributeStatusDuplicateAttribute,
	"invalid_attribute":               AttributeResponseAttributeStatusInvalidAttribute,
	"invalid_attribute_type_conflict": AttributeResponseAttributeStatusInvalidAttributeTypeConflict,
	"attribute_not_processed":         AttributeResponseAttributeStatusAttributeNotProcessed,
}

// GetAttributeResponseAttributeStatusEnumValues Enumerates the set of values for AttributeResponseAttributeStatusEnum
func GetAttributeResponseAttributeStatusEnumValues() []AttributeResponseAttributeStatusEnum {
	values := make([]AttributeResponseAttributeStatusEnum, 0)
	for _, v := range mappingAttributeResponseAttributeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeResponseAttributeStatusEnumStringValues Enumerates the set of values in String for AttributeResponseAttributeStatusEnum
func GetAttributeResponseAttributeStatusEnumStringValues() []string {
	return []string{
		"ATTRIBUTE_ALREADY_ACTIVE",
		"ATTRIBUTE_ACTIVATED",
		"ATTRIBUTE_DEACTIVATED",
		"DEACTIVATION_NOT_ALLOWED",
		"ATTRIBUTE_DOES_NOT_EXIST",
		"ATTRIBUTE_ALREADY_DEACTIVATED",
		"DUPLICATE_ATTRIBUTE",
		"INVALID_ATTRIBUTE",
		"INVALID_ATTRIBUTE_TYPE_CONFLICT",
		"ATTRIBUTE_NOT_PROCESSED",
	}
}

// GetMappingAttributeResponseAttributeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeResponseAttributeStatusEnum(val string) (AttributeResponseAttributeStatusEnum, bool) {
	enum, ok := mappingAttributeResponseAttributeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
