package oss

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/xyy277/gallery/global"

	"github.com/gin-gonic/gin"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type Minio struct{}

func (*Minio) UploadFile(ctx *gin.Context, file *multipart.FileHeader) (string, string, error) {

	minioClient, initError := initialize()
	if initError != nil {
		return "", "", errors.New("minio initialize Filed, err:" + initError.Error())
	}

	f, openError := file.Open()
	if openError != nil {
		global.G_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	objectname := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	bucketName := global.G_CONFIG.Minio.BucketName

	contentType := file.Header.Get("Content-Type")

	// Upload the zip file with PutObject
	info, putErr := minioClient.PutObject(ctx, bucketName, objectname, f, file.Size, minio.PutObjectOptions{ContentType: contentType})

	if putErr != nil {
		global.G_LOG.Error(" Create an object in a bucket failed ", zap.Any("putErr", putErr))
	} else {
		global.G_LOG.Info("create an object success ", zap.Any("info", info))
	}
	return global.G_CONFIG.Minio.Endpoint + "/" + bucketName + "/" + objectname, objectname, nil
}

func (*Minio) DeleteFile(ctx *gin.Context, key string) error {
	minioClient, initError := initialize()

	if initError != nil {
		return errors.New("minio initialize Filed, err:" + initError.Error())
	}
	err := minioClient.RemoveObject(ctx, global.G_CONFIG.Minio.BucketName, key, minio.RemoveObjectOptions{ForceDelete: true})

	return err
}

func initialize() (*minio.Client, error) {
	minioClient, err := minio.New(global.G_CONFIG.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.G_CONFIG.Minio.AccessKeyID, global.G_CONFIG.Minio.SecretAccessKey, ""),
		Secure: global.G_CONFIG.Minio.UseSSL,
	})
	if err != nil {
		global.G_LOG.Error("minio initialize Filed", zap.Any("err", err.Error()))
		return nil, err
	}
	return minioClient, nil
}

func (*Minio) CreateBucket(ctx *gin.Context, bucketName string, location string) error {

	minioClient, initError := initialize()

	if initError != nil {
		return errors.New("minio initialize Filed, err:" + initError.Error())
	}

	bucketErr := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})

	if bucketErr != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			global.G_LOG.Info("We already own ", zap.String("bucketName", bucketName))
		} else {
			global.G_LOG.Error("MakeBucket Filed", zap.Any("err", bucketErr.Error()))
			return bucketErr
		}
	} else {
		global.G_LOG.Info("Successfully create bucket ", zap.String("bucketName", bucketName))
	}

	return nil
}

func (*Minio) DownloadFile(ctx *gin.Context, key string) error {
	minioClient, initError := initialize()

	if initError != nil {
		global.G_LOG.Error("minio initialize Filed, err:" + initError.Error())
		return initError
	}

	object, err := minioClient.GetObject(ctx, global.G_CONFIG.Minio.BucketName, key, minio.GetObjectOptions{})

	if err != nil {
		global.G_LOG.Error("minio GetObject Filed, err:" + err.Error())
		return err
	}

	file, err := os.Create(key)
	if err != nil {
		global.G_LOG.Error("Create file Filed, err:" + err.Error())
		return err
	}

	if _, err = io.Copy(file, object); err != nil {
		global.G_LOG.Error("Copy file Filed, err:" + err.Error())
		return err
	}

	return nil
}
