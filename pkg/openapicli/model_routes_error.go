/*
Swagger baaa API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapicli

import (
	"encoding/json"
)

// RoutesError struct for RoutesError
type RoutesError struct {
	Name string `json:"name"`
}

// NewRoutesError instantiates a new RoutesError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutesError(name string) *RoutesError {
	this := RoutesError{}
	this.Name = name
	return &this
}

// NewRoutesErrorWithDefaults instantiates a new RoutesError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutesErrorWithDefaults() *RoutesError {
	this := RoutesError{}
	return &this
}

// GetName returns the Name field value
func (o *RoutesError) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoutesError) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoutesError) SetName(v string) {
	o.Name = v
}

func (o RoutesError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRoutesError struct {
	value *RoutesError
	isSet bool
}

func (v NullableRoutesError) Get() *RoutesError {
	return v.value
}

func (v *NullableRoutesError) Set(val *RoutesError) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutesError) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutesError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutesError(val *RoutesError) *NullableRoutesError {
	return &NullableRoutesError{value: val, isSet: true}
}

func (v NullableRoutesError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutesError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

