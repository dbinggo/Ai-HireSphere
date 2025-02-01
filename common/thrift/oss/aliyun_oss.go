package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

// AliyunOSSClient represents the Aliyun OSS client
type AliyunOSSClient struct {
	client *oss.Client
	bucket *oss.Bucket
}

type AliyunOssConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

// NewAliyunOSSClient initializes a new Aliyun OSS client
func NewAliyunOSSClient(config AliyunOssConfig) (*AliyunOSSClient, error) {
	// Create OSSClient instance.
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create OSS client: %w", err)
	}

	// Get bucket.
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}

	return &AliyunOSSClient{
		client: client,
		bucket: bucket,
	}, nil
}
func MustNewAliyunOSSClient(config AliyunOssConfig) *AliyunOSSClient {
	client, err := NewAliyunOSSClient(config)
	if err != nil {
		panic(err)
	}
	return client
}

// UploadFile uploads bytes to the OSS bucket
func (c *AliyunOSSClient) UploadFile(objectName string, reader io.Reader) error {
	// Upload bytes.
	err := c.bucket.PutObject(objectName, reader)
	if err != nil {
		return fmt.Errorf("failed to upload bytes: %w", err)
	}
	return nil
}

// DownloadFile downloads a file from the OSS bucket
func (c *AliyunOSSClient) DownloadFile(objectName, filePath string) error {
	// Download file.
	err := c.bucket.GetObjectToFile(objectName, filePath)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	return nil
}

// DeleteFile deletes a file from the OSS bucket
func (c *AliyunOSSClient) DeleteFile(objectName string) error {
	// Delete file.
	err := c.bucket.DeleteObject(objectName)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
