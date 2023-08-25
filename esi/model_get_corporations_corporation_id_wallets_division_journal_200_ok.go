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

import (
	"time"
)

/* A list of GetCorporationsCorporationIdWalletsDivisionJournal200Ok. */
//easyjson:json
type GetCorporationsCorporationIdWalletsDivisionJournal200OkList []GetCorporationsCorporationIdWalletsDivisionJournal200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdWalletsDivisionJournal200Ok struct {
	Amount        float64   `json:"amount,omitempty"`          /* The amount of ISK given or taken from the wallet as a result of the given transaction. Positive when ISK is deposited into the wallet and negative when ISK is withdrawn */
	Balance       float64   `json:"balance,omitempty"`         /* Wallet balance after transaction occurred */
	ContextId     int64     `json:"context_id,omitempty"`      /* An ID that gives extra context to the particular transaction. Because of legacy reasons the context is completely different per ref_type and means different things. It is also possible to not have a context_id */
	ContextIdType string    `json:"context_id_type,omitempty"` /* The type of the given context_id if present */
	Date          time.Time `json:"date,omitempty"`            /* Date and time of transaction */
	Description   string    `json:"description,omitempty"`     /* The reason for the transaction, mirrors what is seen in the client */
	FirstPartyId  int32     `json:"first_party_id,omitempty"`  /* The id of the first party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name */
	Id            int64     `json:"id,omitempty"`              /* Unique journal reference ID */
	Reason        string    `json:"reason,omitempty"`          /* The user stated reason for the transaction. Only applies to some ref_types */
	RefType       string    `json:"ref_type,omitempty"`        /* \"The transaction type for the given. transaction. Different transaction types will populate different attributes. Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can look at the following file to see string->int mappings: https://github.com/ccpgames/eve-glue/blob/master/eve_glue/wallet_journal_ref.py\" */
	SecondPartyId int32     `json:"second_party_id,omitempty"` /* The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name */
	Tax           float64   `json:"tax,omitempty"`             /* Tax amount received. Only applies to tax related transactions */
	TaxReceiverId int32     `json:"tax_receiver_id,omitempty"` /* The corporation ID receiving any tax paid. Only applies to tax related transactions */
}
