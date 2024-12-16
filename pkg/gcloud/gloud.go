package gcloud

import (
	"backend-bangkit/pkg/errs"
	"context"

	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

type GCSUploader struct {
	client     *storage.Client
	bucketName string
	projectID  string
}

func NewGCSUploader() (*GCSUploader, errs.MessageErr) {
	ctx := context.Background()

	// Dapatkan project ID dari environment
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	if projectID == "" {
		return nil, errs.NewBadRequest("missing Google Cloud Project ID in environment variables")
	}

	bucketName := os.Getenv("GCLOUD_BUCKET_ID")
	if bucketName == "" {
		return nil, errs.NewBadRequest("missing Google Cloud Bucket ID in environment variables")
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/app/json/original-advice-438105-i6-9ed330e0dc52.json")

	client, err := storage.NewClient(ctx)
	if err != nil {

		fmt.Println(err)
		return nil, errs.NewBadRequest("failed to create GCS client")
	}

	return &GCSUploader{
		client:     client,
		bucketName: bucketName,
		projectID:  projectID,
	}, nil
}

// func (uploader *GCSUploader) UploadFile(file io.Reader, fileName string, contentType string) (string, errs.MessageErr) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
// 	defer cancel()

// 	bucket := uploader.client.Bucket(uploader.bucketName)
// 	objectName := uuid.New().String() + "_" + fileName
// 	wc := bucket.Object(objectName).NewWriter(ctx)
// 	wc.ContentType = contentType

// 	if _, err := io.Copy(wc, file); err != nil {
// 		return "", errs.NewBadRequest("failed to upload file to GCS")
// 	}
// 	if err := wc.Close(); err != nil {
// 		return "", errs.NewBadRequest("failed to close GCS writer")
// 	}

// 	// Generate URL
// 	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", uploader.bucketName, objectName)
// 	return url, nil
// }

func (uploader *GCSUploader) UploadFile(file io.Reader, contentType string) (string, errs.MessageErr) {
	// Perpanjang timeout menjadi 10 menit untuk unggahan file besar
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	bucket := uploader.client.Bucket(uploader.bucketName)
	objectName := uuid.New().String()
	wc := bucket.Object(objectName).NewWriter(ctx)

	wc.ContentType = contentType
	wc.ChunkSize = 1024 * 1024 * 5 // Menggunakan chunks 5 MB untuk unggahan besar

	// Coba unggah file ke GCS
	if _, err := io.Copy(wc, file); err != nil {
		cancel() // Batalkan konteks jika terjadi kesalahan
		return "", errs.NewBadRequest("failed to upload file to GCS")
	}

	// Tutup writer setelah unggahan selesai
	if err := wc.Close(); err != nil {
		return "", errs.NewBadRequest(err.Error())
	}

	// Generate URL publik untuk file
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", uploader.bucketName, objectName)
	return url, nil
}

func (uploader *GCSUploader) DeleteFile(objectName string) errs.MessageErr {
	// Buat konteks dengan timeout untuk operasi penghapusan
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	// Dapatkan referensi objek dari bucket
	bucket := uploader.client.Bucket(uploader.bucketName)
	object := bucket.Object(objectName)

	// Lakukan penghapusan
	if err := object.Delete(ctx); err != nil {
		return errs.NewInternalServerError("failed to delete file from GCS")
	}

	return nil
}
