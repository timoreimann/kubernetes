/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testsuites

import (
	"fmt"
	"testing"
)

func TestShortenCSIDriverName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "my.driver.acme.com",
			want: "m.d.a.c",
		},
		{
			in:   "not-a-dns-name",
			want: "not-a-dns-name",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=%s", tt.in, tt.want), func(t *testing.T) {
			got := shortenCSIDriverName(tt.in)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
