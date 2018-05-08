/*
Copyright DI Miliy Andrew 2018 All Rights Reserved.

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

package archive

import "github.com/aws/aws-sdk-go/service/s3/s3iface"

// Archive is an abstraction of file archive implementation
type Archive interface {
	ArchiveFile(path string) error

	DeleteFile(path string) error

	// Real implementation should contain more methods
}

// NewS3Archive returns AWS S3 based implementation of Archive
func NewS3Archive(svc s3iface.S3API) (a Archive, err error) {
	return S3Archive{svc: svc}, nil
}
