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

/* A list of GetFwLeaderboardsActiveTotalActiveTotal1. */
//easyjson:json
type GetFwLeaderboardsActiveTotalActiveTotal1List []GetFwLeaderboardsActiveTotalActiveTotal1

/* active_total object */
//easyjson:json
type GetFwLeaderboardsActiveTotalActiveTotal1 struct {
	Amount    int32 `json:"amount,omitempty"`     /* Amount of victory points */
	FactionId int32 `json:"faction_id,omitempty"` /* faction_id integer */
}
