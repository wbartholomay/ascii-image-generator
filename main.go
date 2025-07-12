package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// openai "github.com/openai/openai-go"
	"github.com/wbartholomay/ascii-image-generator/internal/imagetoascii"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables.")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Missing API_KEY environment variablee")
	}

	sizeFlag := flag.Float64("size", 1.0, "size mulitplier for the result image. Capped at 4")
	imageDir := flag.String("imageDir", "", "directory for image")
	outFilePath := flag.String("out", "", "output destination for text. If not provided, text is printed to stdout and not persisted.")
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
	if *outFilePath == "" {
		fmt.Println(result)
		os.Exit(0)
	}

	outFile, err := os.OpenFile(*outFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatalf("Error opening output file at path: %v\n", *outFilePath)
	}
	_, err = outFile.Write([]byte(result))
	if err != nil {
		log.Fatal("Error writing result to output file.")
	}
}
