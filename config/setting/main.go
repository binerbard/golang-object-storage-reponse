package setting

import utilssetting "object_storage/utils/setting"

type Config struct {
	StorageAccessKey string
	StorageSecretKey string
	StorageBucketName string
	StorageHost string
	StoragePort string
}


func LoadConfig() Config{
	utilssetting.LoadEnv()
	return Config{
		StorageAccessKey: utilssetting.GetEnv("STORAGE_ACCESS_KEY", ""),
		StorageSecretKey: utilssetting.GetEnv("STORAGE_SECRET_KEY", ""),
		StorageBucketName: utilssetting.GetEnv("STORAGE_BUCKET_NAME", ""),
		StorageHost: utilssetting.GetEnv("STORAGE_HOST", ""),
		StoragePort: utilssetting.GetEnv("STORAGE_PORT", ""),
	}
}

var	ConfigSetting = LoadConfig()
var StorageEndpoint = ConfigSetting.StorageHost + ":" + ConfigSetting.StoragePort