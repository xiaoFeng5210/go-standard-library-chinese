package imagecompression

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"testing"

	"bytes"

	"time"

	"github.com/disintegration/imaging"
)

func TestCompressImage(t *testing.T) {
	start := time.Now()
	fmt.Println("Start time:", start)
	currentPath := os.Getenv("PWD")
	imagePath := filepath.Join(currentPath, "..", "assets", "images", "test-compression.jpg")
	imageFile, err := os.Open(imagePath)

	defer imageFile.Close()

	if err != nil {
		t.Fatalf("Failed to open image file: %v", err)
	}
	img, _, err := image.Decode(imageFile)
	newWidth := img.Bounds().Dx() / 2
	newHeight := img.Bounds().Dy() / 2
	dstImage := CompressImage(img, newWidth, newHeight)

	savePath := filepath.Join(currentPath, "images", "dst.jpg")

	// err = imaging.Save(dstImage, savePath, imaging.JPEGQuality(80))
	// if err != nil {
	// 	t.Fatalf("Failed to save image: %v", err)
	// }

	var buf bytes.Buffer

	err = imaging.Encode(&buf, dstImage, imaging.JPEG)
	if err != nil {
		t.Fatalf("Failed to encode image: %v", err)
	}

	outFile, err := os.Create(savePath)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}

	defer outFile.Close()
	writer := bufio.NewWriter(outFile)

	_, err = writer.Write(buf.Bytes())
	if err != nil {
		t.Fatalf("Failed to write image data: %v", err)
	}
	err = writer.Flush()
	if err != nil {
		t.Fatalf("Failed to flush writer: %v", err)
	}

	fmt.Println("Image saved to", savePath)
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
}
