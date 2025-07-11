package main

import (
	"flag"
	"fmt"

	"github.com/wbartholomay/ascii-image-generator/internal/imagetoascii"
)

func main() {
	sizeFlag := flag.Float64("size", 1.0, "size mulitplier for the result image. Capped at 4")
	imageDir := flag.String("imageDir", "", "directory for image")
	flag.Parse()
	if *imageDir == "" {
		fmt.Println("Must provide image directory in the form of a flag: --imageDir={PATH_TO_IMAGE}")
		return
	}
	sizeMult := *sizeFlag
	if sizeMult > 4 {
		sizeMult = 4
	} else if sizeMult < 0.25 {
		sizeMult = 0.25
	}

	result, err := imagetoascii.ConvertImageFileToASCII(*imageDir, sizeMult)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
