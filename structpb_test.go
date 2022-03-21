// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorList_Error(t *testing.T) {
	test := map[string]interface{}{
		"SS": []string{"one"},
		"SM": map[string]string{"yellow": "good"},
		"IS": []int{3},
		"IM": map[string]int{"red": 5},
	}
	expect := map[string]interface{}{
		"SS": []interface{}{"one"},
		"SM": map[string]interface{}{"yellow": "good"},
		"IS": []interface{}{float64(3)},
		"IM": map[string]interface{}{"red": float64(5)},
	}
	got, err := NewStruct(test)
	assert.NoError(t, err)
	assert.Equal(t, expect, got.AsMap())
}
