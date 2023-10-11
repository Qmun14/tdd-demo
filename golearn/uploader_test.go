package golearn

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

type MockS3 struct {
	mock.Mock
}

func (m MockS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)

}

func TestUploader_Upload(t *testing.T) {
	m := new(MockS3)
	u := Uploader{
		svc: m,
	}

	m.On("PutObject", &s3.PutObjectInput{
		Body:                 aws.ReadSeekCloser(strings.NewReader("filetoupload")),
		Bucket:               aws.String("examplebucket"),
		Key:                  aws.String("exampleobject"),
		ServerSideEncryption: aws.String("AES256"),
		Tagging:              aws.String("key1=value1&key2=value2"),
	}).Return(&s3.PutObjectOutput{}, nil)

	err := u.Upload()
	assert.NoError(t, err)

	m.AssertExpectations(t)
}
