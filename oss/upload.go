package oss

import (
	"mime/multipart"

	"github.com/xyy277/gallery/global"

	"github.com/gin-gonic/gin"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(ctx *gin.Context, file *multipart.FileHeader) (string, string, error)
	DeleteFile(ctx *gin.Context, key string) error
	CreateBucket(ctx *gin.Context, bucketName string, location string) error
	DownloadFile(ctx *gin.Context, key string) error
}

// NewOss OSS的实例化方法
func NewOss() OSS {
	switch global.G_CONFIG.System.OssType {
	case "minio":
		return &Minio{}
	default:
		return nil
	}
}
