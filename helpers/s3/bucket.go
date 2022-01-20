package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
)

func ImageUpload(uploader *s3manager.Uploader, file *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	fileUp, err := file.Open()
	if err != nil {
		return nil, err
	}
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("shivalta-bucket"),
		ACL:         aws.String("public-read"),
		Key:         aws.String(file.Filename),
		ContentType: aws.String("image/jpeg"),
		Body:        fileUp,
	})
	if err != nil {
		return nil, err
	}
	return up, nil
}
