package utilsstorage

import (
	"bytes"
	"context"
	"mime/multipart"
	"object_storage/config/setting"
	"object_storage/internal/infrastructure/storage"
	"strings"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadFile(file *multipart.FileHeader)(string, error) {
	objectName := uuid.New().String() + "." + strings.Split(file.Filename, ".")[1]
	bucketName := setting.ConfigSetting.StorageBucketName
	
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = storage.MinioClient.PutObject(context.Background(), bucketName, objectName, f, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func ShowFile(objectReader string) (*bytes.Reader, error) {
	bucketName := setting.ConfigSetting.StorageBucketName
	objectName := objectReader
	object, err := storage.MinioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(object)
	if err != nil {
		return nil, err
	}

	objectResponse := bytes.NewReader(buf.Bytes())
	return objectResponse, nil
}