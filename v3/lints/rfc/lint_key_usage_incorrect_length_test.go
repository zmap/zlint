package rfc

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestIncorrectKeyUsageLength(t *testing.T) {
	data := []struct {
		file string
		want lint.LintStatus
	}{
		{
			"incorrect_ku_length.pem", lint.Error,
		},
		{
			"facebookOnionV3Address.pem", lint.Pass,
		},
	}
	for _, testData := range data {
		data := testData
		t.Run(data.file, func(t *testing.T) {
			out := test.TestLint("e_key_usage_incorrect_length", data.file)
			if out.Status != data.want {
				t.Errorf("%s: expected %s, got %s", data.file, data.want, out.Status)
			}
		})
	}
}

func TestIncorrectKeyUsageLengthDirectly(t *testing.T) {
	type input struct {
		input []byte
		want  lint.LintStatus
	}
	data := make([]input, 0)
	// We have to do zero by hand because big.Int
	// will represent 0 as an empty slice rather than 0.
	data = append(data, input{
		input: []byte{3, 2, 0, 0},
		want:  lint.Pass,
	})
	for i := 1; i < 512; i++ {
		b := big.NewInt(int64(i))
		bytes := b.Bytes()
		// Padding cannot exceed 7 bits, so even though there are
		// eight trailing zeroes, we need to declare them as being used.
		var unused byte
		if i == 256 {
			unused = 0
		} else {
			unused = byte(b.TrailingZeroBits())
		}
		data = append(data, input{
			input: append([]byte{3, byte(1 + len(bytes)), unused}, bytes...),
			want:  lint.Pass,
		})
	}
	data = append(data, []input{
		{
			input: []byte{},
			want:  lint.Error,
		},
		{
			input: []byte{3},
			want:  lint.Error,
		},
		{
			input: []byte{1, 2, 0, 0},
			want:  lint.Error,
		},
		{
			input: []byte{3, 3, 7, 0b10000000, 0b10000000},
			want:  lint.Pass,
		},
		{
			input: []byte{3, 3, 0, 0b00000011, 0b11111111},
			want:  lint.Error,
		},
		{
			input: []byte{3, 3, 0, 0b00000001, 0b11111111},
			want:  lint.Pass,
		},
		{
			input: []byte{3, 3, 1, 0b00000011, 0b11111110},
			want:  lint.Pass,
		},
		{
			input: []byte{3, 3, 8, 0b00000011, 0b00000000},
			want:  lint.Error,
		},
	}...)
	for _, d := range data {
		dd := d
		t.Run(fmt.Sprintf("%v", dd.input), func(t *testing.T) {
			got := keyUsageIncorrectLengthBytes(d.input)
			if got.Status != d.want {
				t.Errorf("expected %v, got %v (details:'%s')", dd.want, got.Status, got.Details)
				t.Error(got.Details)
			}
		})
	}
}
