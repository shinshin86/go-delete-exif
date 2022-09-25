package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func usage() {
	fmt.Println("Usage: go-delete-exif <input image path(JPEG only)>")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: You must specify the image (JPEG) from which you want to delete Exif.")
		usage()
		os.Exit(0)
	}

	imgPath := os.Args[1]
	_, err := os.Stat(imgPath)
	if err != nil {
		fmt.Println("Error: Invalid specify filepath")
		os.Exit(1)
	}

	// Read jpeg
	inputImg, err := os.Open(imgPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer inputImg.Close()

	img, err := jpeg.Decode(inputImg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Create tmp image(png)
	tmpPng := filepath.Base(imgPath) + ".png"
	tmpPngPath := filepath.Join(filepath.Dir(imgPath), tmpPng)
	tmpImg, err := os.Create(tmpPngPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer tmpImg.Close()

	// convert png and delete exif
	png.Encode(tmpImg, img)

	outImgName := strings.Split(filepath.Base(imgPath), ".")[0] + "_exif-deleted.JPG"
	outputPath := filepath.Join(filepath.Dir(imgPath), outImgName)
	outImg, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer outImg.Close()

	jpeg.Encode(outImg, img, nil)

	// Delete tmp file
	if err := os.Remove(tmpPngPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Success: delete exif ===> %s\n", outputPath)
}
