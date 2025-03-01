// Copyright 2021 ORAS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

// Manifest describes an ORAS artifact.
// This structure provides `application/vnd.oras.artifact.manifest.v1+json` mediatype when marshalled to JSON.
type Manifest struct {
	// ArtifactType is the artifact type of the object this schema refers to.
	ArtifactType string `json:"artifactType"`

	// Blobs is a collection of blobs referenced by this manifest.
	Blobs []Descriptor `json:"blobs"`

	// Subject is an optional reference to any existing manifest within the repository.
	// When specified, the artifact is said to be dependent upon the referenced subject.
	Subject Descriptor `json:"subject"`

	// Annotations contains arbitrary metadata for the artifact manifest.
	Annotations map[string]string `json:"annotations,omitempty"`
}
