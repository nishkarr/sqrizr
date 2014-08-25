package sqrizrlib

import (
	"io"

	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	_ "image/gif"
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

func Sqrize(srcImageFile io.Reader, outputFile io.Writer) (format, orientation string, err error) {

	var (
		src 		image.Image
	)

	src, format, err = image.Decode(srcImageFile)
	if err != nil {
		return
	}

	orientation, longestSide, sp := determineOrientation(src.Bounds())

	dst := createDstSquareImg(longestSide);

	draw.Draw(dst, dst.Bounds(), src, sp, draw.Src)

	err = png.Encode(outputFile, dst)

	return format, orientation, err
}
