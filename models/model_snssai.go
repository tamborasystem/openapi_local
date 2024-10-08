// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * 3gpp-traffic-influence
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Snssai struct {
	Sst int32  `json:"sst" yaml:"sst" bson:"sst" mapstructure:"Sst"`
	Sd  string `json:"sd,omitempty" yaml:"sd" bson:"sd" mapstructure:"Sd"`
	Nid string `json:"nid,omitempty" yaml:"nid" bson:"nid" mapstructure:"nid"`
}
