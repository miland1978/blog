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

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3Archive is some lame implementation over AWS S3
//go:generate mockgen -destination=gomock_mocks/mock_s3api.go -package=gomock_mocks github.com/aws/aws-sdk-go/service/s3/s3iface S3API
//go:generate mockery -dir ../../../aws/aws-sdk-go/service/s3/s3iface/ -name S3API -output testify_mocks -outpkg testify_mocks
type S3Archive struct {
	svc s3iface.S3API
}

// ArchiveFile supposed to implement archiving of the local file to AWS S3
func (S3Archive) ArchiveFile(path string) error {
	return nil
}

// DeleteFile supposed to implement archiving of the local file to AWS S3
func (a S3Archive) DeleteFile(path string) error {
	i := s3.DeleteObjectInput{}
	_, err := a.svc.DeleteObject(&i)
	return err
}
