package storage

import (
	"context"
	"fmt"
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"strings"
)

type MinioClient struct {
	Config *config.Config
	Client *minio.Client
}

func NewMinioClient(config *config.Config) *MinioClient {
	return &MinioClient{Config: config}
}

func (m *MinioClient) connect() error {
	minioUrl := fmt.Sprintf("%s:%d", m.Config.Minio.Host, m.Config.Minio.Port)

	// Initialize MinIO client
	client, err := minio.New(minioUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(m.Config.Minio.AccessKey, m.Config.Minio.SecretKey, ""),
		Secure: m.Config.Minio.UseSSL,
	})
	if err != nil {
		return err
	}

	m.Client = client
	return nil
}

func (m *MinioClient) objectExists(bucket, object string) error {
	_, err := m.Client.StatObject(context.Background(), bucket, object, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return err
		}
		return err
	}

	return nil
}

func (m *MinioClient) FileExists(path string) (bool, error) {
	// Extract bucket name and object name
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		log.Fatalln("Invalid object path format")
	}
	bucketName := parts[0]
	objectName := parts[1]

	if err := m.connect(); err != nil {
		return false, err
	}

	if err := m.objectExists(bucketName, objectName); err != nil {
		return false, err
	}

	return true, nil
}
