package repository

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
// interface for account repository
type S3Repository interface{
	PutFileURL(fileName string)(string)
}

// Constructor for AccountRepository
func NewS3Repository()(S3Repository){
	return &S3RepositoryImplementation{
		// configuration for s3
		client: s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1")),
	}
}

// Class for account repository implementation
type S3RepositoryImplementation struct{
	client *s3.S3
}

func (repository *S3RepositoryImplementation) PutFileURL(fileName string)(string){
	req, _ := repository.client.PutObjectRequest(&s3.PutObjectInput{
        Bucket: aws.String("qiup-image"),
		Key: aws.String(fileName + ".png"),
		ACL: aws.String("public-read"),
    })
    url,_ := req.Presign(15 * time.Minute)

    return url
}