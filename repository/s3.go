package repository

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//S3Repository interface
type S3Repository interface {
	PutImage(directory string) string
}

//NewS3Repository constructor
func NewS3Repository() S3Repository {
	return &s3RepositoryImplementation{
		client: s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1")),
	}
}

type s3RepositoryImplementation struct {
	client *s3.S3
}

//PutImage put
func (repository *s3RepositoryImplementation) PutImage(directory string) string {
	req, _ := repository.client.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String("qiup-image"),
		Key:         aws.String(directory),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/jpeg"),
	})
	url, _ := req.Presign(15 * time.Minute)
	return url
}
