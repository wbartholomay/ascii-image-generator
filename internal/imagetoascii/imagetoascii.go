package imagetoascii

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func ConvertImageFileToASCII(pathToImage string, sizeMult float64) (string, error) {
	img, _, err := readImage(pathToImage)
	if err != nil {
		return "", fmt.Errorf("error reading image: %w", err)
	}
	img = resizeImage(img, sizeMult)
	return convertImageToASCII(img)
}

func readImage(pathToImage string) (image.Image, string, error) {
	curDirectory, err := os.Getwd()
	if err != nil {
		return nil, "", fmt.Errorf("error getting working directory: %w", err)
	}
	file, err := os.Open(filepath.Join(curDirectory, pathToImage))
	if err != nil {
		return nil, "", fmt.Errorf("error opening image file: %w", err)
	}
	defer file.Close()
	image, format, err := image.Decode(file)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding image file: %w", err)
	}
	return image, format, nil
}

func resizeImage(img image.Image, sizeMult float64) image.Image {
	bounds := img.Bounds()
	sizeX := int(float64(bounds.Dx()) * sizeMult)
	sizeY := int(float64(bounds.Dy()) * sizeMult)
	return resize.Resize(uint(sizeX), uint(sizeY), img, resize.Lanczos3)
}

func convertImageToASCII(img image.Image) (string, error) {
	// TODO could resize image here if needed
	img = resize.Resize(uint(img.Bounds().Dx()), uint(img.Bounds().Dy())/2, img, resize.Lanczos3)
	ramp := []rune{' ', '.', '-', '+', '#', '@'}
	rampSize := len(ramp)
	step := 255 / (rampSize - 1)

	var result strings.Builder
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			originalColor := img.At(x, y)
			grayColor, ok := color.GrayModel.Convert(originalColor).(color.Gray)
			if !ok {
				return "", fmt.Errorf("error converting image to grayscale")
			}
			intensity := grayColor.Y
			index := int(intensity) / step
			result.WriteRune(ramp[index])
		}
		result.WriteString("\n")
	}
	return result.String(), nil
}
