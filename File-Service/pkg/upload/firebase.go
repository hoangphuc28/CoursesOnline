package upload

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go/v4"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/File-Service/config"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

type FireBaseProvider struct {
	Bucket *storage.BucketHandle
}

func NewFireBaseProvider(cf *config.Config) *FireBaseProvider {
	opt := option.WithCredentialsFile(cf.FireBase.PathCredentialsFile)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		StorageBucket: cf.FireBase.Bucket,
	}, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	return &FireBaseProvider{bucket}
}

func (u UploadFileProvider) UploadFileFireBase(data []byte, storageObjectPath string) (string, error) {
	obj := u.FireBaseProvider.Bucket.Object(storageObjectPath)
	wc := obj.NewWriter(context.Background())
	contentType := http.DetectContentType(data)
	// Set the content type of the uploaded data (e.g., "image/jpeg")
	wc.ContentType = contentType
	fmt.Println(contentType)

	if _, err := wc.Write(data); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	attrs, err := obj.Attrs(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Println(attrs.Bucket)

	fmt.Println(attrs.Name)

	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", attrs.Bucket, strings.Replace(attrs.Name, "/", "%2F", -1)), nil
}
