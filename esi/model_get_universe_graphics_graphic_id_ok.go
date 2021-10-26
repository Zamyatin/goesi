/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.8.2
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

/* A list of GetUniverseGraphicsGraphicIdOk. */
//easyjson:json
type GetUniverseGraphicsGraphicIdOkList []GetUniverseGraphicsGraphicIdOk

/* 200 ok object */
//easyjson:json
type GetUniverseGraphicsGraphicIdOk struct {
	CollisionFile string `json:"collision_file,omitempty"`  /* collision_file string */
	GraphicFile   string `json:"graphic_file,omitempty"`    /* graphic_file string */
	GraphicId     int32  `json:"graphic_id,omitempty"`      /* graphic_id integer */
	IconFolder    string `json:"icon_folder,omitempty"`     /* icon_folder string */
	SofDna        string `json:"sof_dna,omitempty"`         /* sof_dna string */
	SofFationName string `json:"sof_fation_name,omitempty"` /* sof_fation_name string */
	SofHullName   string `json:"sof_hull_name,omitempty"`   /* sof_hull_name string */
	SofRaceName   string `json:"sof_race_name,omitempty"`   /* sof_race_name string */
}
