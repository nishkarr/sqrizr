package main

import (
	"fmt"
	"flag"
	"os"

	"image"
//	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"

)

var (
	srcImagePath string
)

type Orientation int

const (
	Portrait Orientation = iota
	Landscape
	Square
)

func (o Orientation) String() string {
	switch(o){
		case Portrait:
			return "Portrait"
		case Landscape:
			return "Landscape"
		case Square:
			return "Square"
		default:
			return "UNKNOWN"
	}
}


func determineOrientation(rect image.Rectangle) Orientation {
	if rect.Dx() > rect.Dy(){
		return Landscape
	} else if rect.Dx() < rect.Dy() {
				return Portrait
			} else {
				return Square
			}
}
 
func main(){
	flag.StringVar(&srcImagePath, "src", "", "Source image path")
	flag.Parse()

	if srcImagePath == "" {
		fmt.Println("Missing parameters")
		flag.PrintDefaults()
		return
	}

	var (
		err 				error
		srcImageFile 	*os.File
		src 				image.Image
		format			string
	)

	srcImageFile, err = os.Open(srcImagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcImageFile.Close();

	src, format, err = image.Decode(srcImageFile)

	srcBounds := src.Bounds()
	orientation := determineOrientation(srcBounds)

	fmt.Println(fmt.Sprintf("processing file: %s\nformat: %s\ndimensions: %v (%v)\n", srcImagePath, format, srcBounds, orientation))
}

