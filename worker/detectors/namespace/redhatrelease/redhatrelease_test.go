// Copyright 2016 clair authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redhatrelease

import (
	"testing"

	"github.com/coreos/clair/database"
	"github.com/coreos/clair/worker/detectors/namespace"
)

func TestRedhatReleaseNamespaceDetector(t *testing.T) {
	testData := []namespace.TestData{
		{
			ExpectedNamespace: &database.Namespace{Name: "oracle:6"},
			Data: map[string][]byte{
				"etc/oracle-release": []byte(`Oracle Linux Server release 6.8`),
			},
		},
		{
			ExpectedNamespace: &database.Namespace{Name: "oracle:7"},
			Data: map[string][]byte{
				"etc/oracle-release": []byte(`Oracle Linux Server release 7.2`),
			},
		},
		{
			ExpectedNamespace: &database.Namespace{Name: "centos:6"},
			Data: map[string][]byte{
				"etc/centos-release": []byte(`CentOS release 6.6 (Final)`),
			},
		},
		{
			ExpectedNamespace: &database.Namespace{Name: "centos:7"},
			Data: map[string][]byte{
				"etc/system-release": []byte(`CentOS Linux release 7.1.1503 (Core)`),
			},
		},
	}

	namespace.TestDetector(t, &RedhatReleaseNamespaceDetector{}, testData)
}
