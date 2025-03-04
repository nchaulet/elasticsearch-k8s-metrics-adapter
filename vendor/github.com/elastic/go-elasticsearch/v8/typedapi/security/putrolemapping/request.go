// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package putrolemapping

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putrolemapping
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/security/put_role_mapping/SecurityPutRoleMappingRequest.ts#L24-L43
type Request struct {
	Enabled  *bool                  `json:"enabled,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Roles    []string               `json:"roles,omitempty"`
	Rules    *types.RoleMappingRule `json:"rules,omitempty"`
	RunAs    []string               `json:"run_as,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putrolemapping request: %w", err)
	}

	return &req, nil
}
