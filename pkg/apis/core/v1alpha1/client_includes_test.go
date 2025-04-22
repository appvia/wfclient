/**
 * Copyright 2021 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludedFiles_Get(t *testing.T) {
	tests := []struct {
		name     string
		files    IncludedFiles
		fileName string
		want     bool
		wantFile *IncludedFile
	}{
		{
			name:     "nil files returns false",
			files:    nil,
			fileName: "test.txt",
			want:     false,
			wantFile: nil,
		},
		{
			name: "existing file returns true and file",
			files: IncludedFiles{
				{Name: "test.txt", Content: strPtr("content")},
			},
			fileName: "test.txt",
			want:     true,
			wantFile: &IncludedFile{Name: "test.txt", Content: strPtr("content")},
		},
		{
			name: "non-existing file returns false",
			files: IncludedFiles{
				{Name: "test.txt", Content: strPtr("content")},
			},
			fileName: "other.txt",
			want:     false,
			wantFile: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotFile := tt.files.Get(tt.fileName)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantFile, gotFile)
		})
	}
}

func TestIncludedFiles_GetNotIncludedFiles(t *testing.T) {
	tests := []struct {
		name  string
		files IncludedFiles
		want  []string
	}{
		{
			name:  "nil files returns nil",
			files: nil,
			want:  nil,
		},
		{
			name: "returns files with nil content",
			files: IncludedFiles{
				{Name: "test1.txt", Content: nil},
				{Name: "test2.txt", Content: strPtr("content")},
				{Name: "test3.txt", Content: nil},
			},
			want: []string{"test1.txt", "test3.txt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.files.GetNotIncludedFiles()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetIncludeFileName(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		defaultName string
		want        string
		wantOk      bool
	}{
		{
			name:        "empty content returns false",
			content:     "",
			defaultName: "default.txt",
			want:        "",
			wantOk:      false,
		},
		{
			name:        "content without include prefix returns false",
			content:     "not an include",
			defaultName: "default.txt",
			want:        "",
			wantOk:      false,
		},
		{
			name:        "include directive with no name returns default",
			content:     "!file",
			defaultName: "default.txt",
			want:        "default.txt",
			wantOk:      true,
		},
		{
			name:        "include directive with name returns name",
			content:     "!file test.txt",
			defaultName: "default.txt",
			want:        "test.txt",
			wantOk:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := GetIncludeFileName(tt.content, tt.defaultName)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantOk, gotOk)
		})
	}
}

// Helper function to create string pointer
func strPtr(s string) *string {
	return &s
}
