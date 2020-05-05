package uploader

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

const (
	bucketURL          = "https://storage.googleapis.com/barcelona-zoo-backet"
	credentialFileName = "barcelona-zoo-282d16c6bf66.json"
)

func ContentUploadToGCS(ctx context.Context, content multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFileName))
	if err != nil {
		log.Fatal(err)
	}

	// オブジェクトのReaderを作成
	bucketName := "barcelona-zoo-backet"
	objectName, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	objectPath := fmt.Sprintf("content/%s%s", objectName.String(), fileHeader.Filename)
	writer := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	writer.ContentType = fileHeader.Header.Get("Content-Type")
	if _, err := io.Copy(writer, content); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	// 成功したらアップロードしたファイルのURLを返却
	return fmt.Sprintf("%s/%s", bucketURL, objectPath), nil
}
