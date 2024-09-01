package storage

import (
	"context"
	"log"
	"object_storage/config/setting"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

var BucketName = setting.ConfigSetting.StorageBucketName
func InitStorage() {
	var err error

	MinioClient, err = minio.New(setting.StorageEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(setting.ConfigSetting.StorageAccessKey, setting.ConfigSetting.StorageSecretKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = MinioClient.MakeBucket(context.Background(), BucketName,minio.MakeBucketOptions{Region: ""})

	if err != nil {
		// Check to see if we already own this bucket.
		exists, err := MinioClient.BucketExists(context.Background(), BucketName)
		if err != nil {
			log.Fatalln(err)
		}
		if !exists {
			log.Fatalln(err)
		}
	}
}