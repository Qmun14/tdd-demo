package golearn

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

type AwsUploader interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

type Uploader struct {
	svc AwsUploader
}

func (u Uploader) Upload() error {

	input := &s3.PutObjectInput{
		Body:                 aws.ReadSeekCloser(strings.NewReader("filetoupload")),
		Bucket:               aws.String("examplebucket"),
		Key:                  aws.String("exampleobject"),
		ServerSideEncryption: aws.String("AES256"),
		Tagging:              aws.String("key1=value1&key2=value2"),
	}

	result, err := u.svc.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	fmt.Println(result)
	return nil
}
