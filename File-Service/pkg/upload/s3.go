package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hoangphuc28/CoursesOnline/File-Service/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type s3Provider struct {
	bucket  string
	region  string
	apiKey  string
	secret  string
	domain  string
	session *session.Session
}

func NewS3Provider(cfg *config.Config) *s3Provider {
	provider := &s3Provider{
		bucket: cfg.AWS.S3Bucket,
		region: cfg.AWS.Region,
		apiKey: cfg.AWS.APIKey,
		secret: cfg.AWS.SecretKey,
		domain: cfg.AWS.S3Domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			"",              // Token có thể bỏ qua
		),
	})

	if err != nil {

		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}
func (uf UploadFileProvider) DeleteFile(objectUrl string) error {
	u, err := url.Parse(objectUrl)
	if err != nil {
		return err
	}
	objectKey := strings.TrimPrefix(u.Path, "/")
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(uf.s3Provider.bucket),
		Key:    aws.String(objectKey),
	}

	_, err = s3.New(uf.s3Provider.session).DeleteObject(input)
	if err != nil {
		return err
	}
	return nil
}
func (uf UploadFileProvider) UploadFile(ctx context.Context, data []byte, dst string) (string, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	_, err := s3.New(uf.s3Provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(uf.s3Provider.bucket),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("%s/%s", uf.s3Provider.domain, dst), nil
}
