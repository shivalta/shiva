package driver

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

var S3Uploader *s3manager.Uploader

func connectAws() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(viper.GetString(`aws_s3.region`)),
			Credentials: credentials.NewStaticCredentials(
				viper.GetString(`aws_s3.access_key_id`),
				viper.GetString(`aws_s3.secret_access_key`),
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func ConnectS3Bucket() {
	S3Uploader = s3manager.NewUploader(connectAws())
}
