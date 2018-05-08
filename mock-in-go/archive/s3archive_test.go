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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/miland1978/blog-go/archive/gomock_mocks"
	"github.com/miland1978/blog-go/archive/testify_mocks"
	"github.com/stretchr/testify/mock"
)

func TestGoMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockS3Api := gomock_mocks.NewMockS3API(mockCtrl)
	a, _ := NewS3Archive(mockS3Api)

	// call to DeleteObject will return no error and should be called once
	mockS3Api.EXPECT().DeleteObject(gomock.Any()).Return(nil, nil).Times(1)

	err := a.DeleteFile("test")

	if err != nil {
		t.Fail()
	}
}

func TestTestify(t *testing.T) {
	mockS3Api := testify_mocks.S3API{}
	a, _ := NewS3Archive(&mockS3Api)

	// call to DeleteObject will return no error and should be called once
	mockS3Api.On("DeleteObject", mock.Anything).Return(nil, nil).Times(1)

	err := a.DeleteFile("test")

	if err != nil {
		t.Fail()
	}

	mockS3Api.AssertExpectations(t)
}

func BenchmarkGoMock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t testing.T
		mockCtrl := gomock.NewController(&t)
		defer mockCtrl.Finish()
		mockS3Api := gomock_mocks.NewMockS3API(mockCtrl)
		a, _ := NewS3Archive(mockS3Api)

		// we expect DeleteObject to be called once with not nil argument
		mockS3Api.EXPECT().DeleteObject(gomock.Not(nil)).Return(nil, nil).Times(1)

		err := a.DeleteFile("test")

		if err != nil {
			t.Fail()
		}
	}
}

func BenchmarkTestify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t testing.T
		mockS3Api := testify_mocks.S3API{}
		a, _ := NewS3Archive(&mockS3Api)

		// call to DeleteObject will return our test error
		//mockS3Api.EXPECT().DeleteObject(gomock.Any()).Return(nil, testError).Times(1)
		mockS3Api.On("DeleteObject", mock.Anything).Return(nil, nil).Once()

		err := a.DeleteFile("test")

		if err != nil {
			t.Fail()
		}

		mockS3Api.AssertExpectations(&t)
	}
}
