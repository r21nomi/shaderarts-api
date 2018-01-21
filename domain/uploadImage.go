package domain

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/xid"
)

type UploadImage struct{}

func (u *UploadImage) Execute(base64Thumb string) (string, error) {
	region := "ap-northeast-1"
	bucket := "arto-image"

	file, err := toPngImage(base64Thumb)
	if err != nil {
		return "", err
	}
	filename := file.Name()
	filePath := "https://s3-" + region + ".amazonaws.com/" + bucket + "/" + filename

	defer file.Close()

	// Initialize a session.
	// the SDK will use to load credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	err = os.Remove(filename)
	if err != nil {
		log.Println("Failed to remove file.")
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	return filePath, err
}

/**
 * Save base64 image as png.
 */
func toPngImage(base64Thumb string) (*os.File, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Thumb))
	img, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	bounds := img.Bounds()
	fmt.Println(bounds, formatString)

	guid := xid.New()
	filePath := ""
	fileName := guid.String() + ".png"
	file, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return file, err
	}

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
		return file, err
	}
	fmt.Println("Png file", fileName, "created")

	return file, err
}
