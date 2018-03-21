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

import (
	"time"
)

/* A list of GetCharactersCharacterIdOk. */
//easyjson:json
type GetCharactersCharacterIdOkList []GetCharactersCharacterIdOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdOk struct {
	Name           string    `json:"name,omitempty"`            /* name string */
	Description    string    `json:"description,omitempty"`     /* description string */
	CorporationId  int32     `json:"corporation_id,omitempty"`  /* The character's corporation ID */
	AllianceId     int32     `json:"alliance_id,omitempty"`     /* The character's alliance ID */
	Birthday       time.Time `json:"birthday,omitempty"`        /* Creation date of the character */
	Gender         string    `json:"gender,omitempty"`          /* gender string */
	RaceId         int32     `json:"race_id,omitempty"`         /* race_id integer */
	BloodlineId    int32     `json:"bloodline_id,omitempty"`    /* bloodline_id integer */
	AncestryId     int32     `json:"ancestry_id,omitempty"`     /* ancestry_id integer */
	SecurityStatus float32   `json:"security_status,omitempty"` /* security_status number */
	FactionId      int32     `json:"faction_id,omitempty"`      /* ID of the faction the character is fighting for, if the character is enlisted in Factional Warfare */
}
