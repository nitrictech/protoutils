// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protoutils

import (
	"encoding/json"
	"unicode/utf8"

	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/structpb"
)

// NewStruct This is a copy of structpb.NewStruct() except we use FromValue()
// instead of structpb.NewValue()
func NewStruct(v map[string]interface{}) (*structpb.Struct, error) {
	x := &structpb.Struct{Fields: make(map[string]*structpb.Value, len(v))}
	for k, v := range v {
		if !utf8.ValidString(k) {
			return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", k)
		}
		var err error
		x.Fields[k], err = ToValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

// ToValue converts custom types to a structpb.Value. This helper was added
// since structpb.NewValue has a limited set of types that it supports.
// More details here: https://github.com/golang/protobuf/issues/1302#issuecomment-805453221
func ToValue(data interface{}) (*structpb.Value, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	v := &structpb.Value{}
	err = v.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}

	return v, nil
}
