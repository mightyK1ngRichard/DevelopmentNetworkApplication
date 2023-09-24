package handler

import (
	"VikingsServer/internal/app/kingMinio"
	"VikingsServer/internal/utils"
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"mime/multipart"
)

func (h *Handler) createImageInMinio(file *multipart.File, header *multipart.FileHeader) (string, error) {
	objectName := header.Filename
	data, err := io.ReadAll(*file)
	if err != nil {
		return "", err
	}

	if errName := utils.GenerateUniqueName(data, &objectName); errName != nil {
		return "", errName
	}

	_, err = h.Minio.PutObject("vikings-server", objectName, *file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://127.0.0.1:9000/%s/%s", kingMinio.BucketName, objectName), nil
}
