/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.6
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

/* A list of GetUniverseBloodlines200Ok. */
//easyjson:json
type GetUniverseBloodlines200OkList []GetUniverseBloodlines200Ok

/* 200 ok object */
//easyjson:json
type GetUniverseBloodlines200Ok struct {
	BloodlineId   int32  `json:"bloodline_id,omitempty"`   /* bloodline_id integer */
	Name          string `json:"name,omitempty"`           /* name string */
	Description   string `json:"description,omitempty"`    /* description string */
	RaceId        int32  `json:"race_id,omitempty"`        /* race_id integer */
	ShipTypeId    int32  `json:"ship_type_id,omitempty"`   /* ship_type_id integer */
	CorporationId int32  `json:"corporation_id,omitempty"` /* corporation_id integer */
	Perception    int32  `json:"perception,omitempty"`     /* perception integer */
	Willpower     int32  `json:"willpower,omitempty"`      /* willpower integer */
	Charisma      int32  `json:"charisma,omitempty"`       /* charisma integer */
	Memory        int32  `json:"memory,omitempty"`         /* memory integer */
	Intelligence  int32  `json:"intelligence,omitempty"`   /* intelligence integer */
}
