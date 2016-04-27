package main

import (
	"bufio"
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"path"
)

type Config struct {
	Width  int
	Height int
	Path   string
}

func main() {

	c := Config{Width: 250, Height: 0, Path: path.Join(os.Getenv("USERPROFILE"), "Desktop")}

	fmt.Printf("%s\n", os.Args[1])

	// open "test.jpg"
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(250, 0, img, resize.Lanczos3)

	out, err := os.Create(path.Join(c.Path, "test_resized.jpg"))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	fmt.Printf("File saved to %s\n", c.Path)
	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
