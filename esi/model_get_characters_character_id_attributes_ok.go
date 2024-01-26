/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.21
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

/* A list of GetCharactersCharacterIdAttributesOk. */
//easyjson:json
type GetCharactersCharacterIdAttributesOkList []GetCharactersCharacterIdAttributesOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdAttributesOk struct {
	AccruedRemapCooldownDate time.Time `json:"accrued_remap_cooldown_date,omitempty"` /* Neural remapping cooldown after a character uses remap accrued over time */
	BonusRemaps              int32     `json:"bonus_remaps,omitempty"`                /* Number of available bonus character neural remaps */
	Charisma                 int32     `json:"charisma,omitempty"`                    /* charisma integer */
	Intelligence             int32     `json:"intelligence,omitempty"`                /* intelligence integer */
	LastRemapDate            time.Time `json:"last_remap_date,omitempty"`             /* Datetime of last neural remap, including usage of bonus remaps */
	Memory                   int32     `json:"memory,omitempty"`                      /* memory integer */
	Perception               int32     `json:"perception,omitempty"`                  /* perception integer */
	Willpower                int32     `json:"willpower,omitempty"`                   /* willpower integer */
}
