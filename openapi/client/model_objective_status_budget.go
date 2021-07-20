/*
 * Pyrra
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ObjectiveStatusBudget struct for ObjectiveStatusBudget
type ObjectiveStatusBudget struct {
	Total     float64 `json:"total"`
	Remaining float64 `json:"remaining"`
	Max       float64 `json:"max"`
}

// NewObjectiveStatusBudget instantiates a new ObjectiveStatusBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectiveStatusBudget(total float64, remaining float64, max float64) *ObjectiveStatusBudget {
	this := ObjectiveStatusBudget{}
	this.Total = total
	this.Remaining = remaining
	this.Max = max
	return &this
}

// NewObjectiveStatusBudgetWithDefaults instantiates a new ObjectiveStatusBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectiveStatusBudgetWithDefaults() *ObjectiveStatusBudget {
	this := ObjectiveStatusBudget{}
	return &this
}

// GetTotal returns the Total field value
func (o *ObjectiveStatusBudget) GetTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ObjectiveStatusBudget) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ObjectiveStatusBudget) SetTotal(v float64) {
	o.Total = v
}

// GetRemaining returns the Remaining field value
func (o *ObjectiveStatusBudget) GetRemaining() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *ObjectiveStatusBudget) GetRemainingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *ObjectiveStatusBudget) SetRemaining(v float64) {
	o.Remaining = v
}

// GetMax returns the Max field value
func (o *ObjectiveStatusBudget) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *ObjectiveStatusBudget) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *ObjectiveStatusBudget) SetMax(v float64) {
	o.Max = v
}

func (o ObjectiveStatusBudget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["remaining"] = o.Remaining
	}
	if true {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableObjectiveStatusBudget struct {
	value *ObjectiveStatusBudget
	isSet bool
}

func (v NullableObjectiveStatusBudget) Get() *ObjectiveStatusBudget {
	return v.value
}

func (v *NullableObjectiveStatusBudget) Set(val *ObjectiveStatusBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectiveStatusBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectiveStatusBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectiveStatusBudget(val *ObjectiveStatusBudget) *NullableObjectiveStatusBudget {
	return &NullableObjectiveStatusBudget{value: val, isSet: true}
}

func (v NullableObjectiveStatusBudget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectiveStatusBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}