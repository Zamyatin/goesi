/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.19
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetContractsPublicItemsContractId200Ok. */
//easyjson:json
type GetContractsPublicItemsContractId200OkList []GetContractsPublicItemsContractId200Ok

/* 200 ok object */
//easyjson:json
type GetContractsPublicItemsContractId200Ok struct {
	IsBlueprintCopy    bool  `json:"is_blueprint_copy,omitempty"`   /* is_blueprint_copy boolean */
	IsIncluded         bool  `json:"is_included,omitempty"`         /* true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract */
	ItemId             int64 `json:"item_id,omitempty"`             /* Unique ID for the item being sold. Not present if item is being requested by contract rather than sold with contract */
	MaterialEfficiency int32 `json:"material_efficiency,omitempty"` /* Material Efficiency Level of the blueprint */
	Quantity           int32 `json:"quantity,omitempty"`            /* Number of items in the stack */
	RecordId           int64 `json:"record_id,omitempty"`           /* Unique ID for the item, used by the contract system */
	Runs               int32 `json:"runs,omitempty"`                /* Number of runs remaining if the blueprint is a copy, -1 if it is an original */
	TimeEfficiency     int32 `json:"time_efficiency,omitempty"`     /* Time Efficiency Level of the blueprint */
	TypeId             int32 `json:"type_id,omitempty"`             /* Type ID for item */
}
