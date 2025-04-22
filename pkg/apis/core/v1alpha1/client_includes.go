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

import "strings"

type IncludedFile struct {
	// Name is the name of the file to include it should be a reference to a relative path
	Name string `json:"name"`
	// Content is a pointer to the data that should be included from the file
	// if the content is nil then the file has not been included
	Content *string `json:"content"`
}
type IncludedFiles []IncludedFile

func (f *IncludedFiles) GetFromDirective(content, defaultName string) (bool, *IncludedFile) {
	name, ok := GetIncludeFileName(content, defaultName)
	if !ok {
		return false, nil
	}
	return f.Get(name)
}

func (f *IncludedFiles) Get(name string) (bool, *IncludedFile) {
	if f == nil {
		return false, nil
	}
	for _, inc := range *f {
		if inc.Name == name {
			return true, &inc
		}
	}

	return false, nil
}

func (f *IncludedFiles) Set(name string, content string) {
	if f == nil {
		return
	}
	if found, inc := f.Get(name); found {
		inc.Content = &content
	} else {
		*f = append(*f, IncludedFile{
			Name:    name,
			Content: &content,
		})
	}
}

func (f IncludedFiles) GetNotIncludedFiles() []string {
	if f == nil {
		return nil
	}
	var notIncluded []string
	for _, inc := range f {
		if inc.Content == nil {
			notIncluded = append(notIncluded, inc.Name)
		}
	}
	return notIncluded
}

// SupportsIncludes is the interface implemented by types which can include fields from files
// +kubebuilder:object:generate=false
type SupportsIncludes interface {
	// HasIncludes will return true if the object has includes
	HasIncludes() bool
	// GetFilesToInclude returns a structure that can be used to include files
	GetFilesToInclude() IncludedFiles
	// UpdateFileContent updates the content of the included files
	// must pass in the original structure returned by GetFilesToInclude
	UpdateFileContent(IncludedFiles)
}

// +kubebuilder:object:generate=false
type ObjectWithIncludes interface {
	Object
	SupportsIncludes
}

// IsObjectWithIncludes returns true if the provided object supports client side includes
func IsObjectWithIncludes(obj Object) (ObjectWithIncludes, bool) {
	incObj, ok := obj.(ObjectWithIncludes)
	if !ok {
		return nil, false
	}
	return incObj, incObj.HasIncludes()
}

func HasIncludeReference(content string) bool {
	_, ok := GetIncludeFileName(content, "")
	return ok
}

const includeFilePrefix = "!file"

// GetIncludeFileName returns the name of the file to include and if the content is a include file
// if an include directive is found but there is no file name it returns the default name
// otherwise it returns and empty string and false
func GetIncludeFileName(content, defaultName string) (string, bool) {

	if len(content) < len(includeFilePrefix) {
		return "", false
	}
	if content[:len(includeFilePrefix)] != includeFilePrefix {
		return "", false
	}

	name := content[len(includeFilePrefix):]

	if name == "" {
		return defaultName, true
	}

	return strings.TrimLeft(name, " "), true
}
