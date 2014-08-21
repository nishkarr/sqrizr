package main

import (
	"fmt"
	"flag"
	"os"

	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	_ "image/gif"

)

var (
	srcImagePath string
)

func determineOrientation(rect image.Rectangle) (orientation string, longestSide int, startingPoint image.Point) {
	if rect.Dx() > rect.Dy(){
		return "Landscape", rect.Dx(), image.Pt(0, (rect.Dx() - rect.Dy()) / -2) 
	} else { 
		if rect.Dx() < rect.Dy() {
			return "Portrait", rect.Dy(), image.Pt((rect.Dy() - rect.Dx()) / -2, 0)
		} else {
			return "Square", rect.Dx(), image.Pt(0,0)
		}
	}
}

func createDstSquareImg(sideLength int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0,0, sideLength, sideLength))
	col := color.RGBA{5, 5, 5, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.ZP, draw.Src)
	return img
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
		err 								error
		srcImageFile, outputFile 	*os.File
		src 								image.Image
		format							string
	)

	srcImageFile, err = os.Open(srcImagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcImageFile.Close();

	src, format, err = image.Decode(srcImageFile)

	orientation, longestSide, sp := determineOrientation(src.Bounds())

	dst := createDstSquareImg(longestSide);

	draw.Draw(dst, dst.Bounds(), src, sp, draw.Src)

	outputFile, err = os.Create("output.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, dst)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("processing file: %s\nformat: %s\ndimensions: %v (%s)\nsquare side: %d", srcImagePath, format, src.Bounds(), orientation, longestSide))
}

