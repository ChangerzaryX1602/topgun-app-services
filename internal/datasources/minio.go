package datasources

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func InitMinioClient() (*minio.Client, error) {
	// Replace with your MinIO server details
	endpoint := viper.GetString("minio.endpoint")
	accessKeyID := viper.GetString("minio.access_key")
	secretAccessKey := viper.GetString("minio.secret_key")
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("Failed to connect to MinIO", err)
		return nil, err
	}
	fmt.Println("Successfully connected to MinIO")
	return client, err
}
