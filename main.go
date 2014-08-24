package main

import (
	"fmt"
	"flag"
	"os"

	"sqrizr/sqrizrlib"
)

var (
	srcImagePath string
)

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
	)

	srcImageFile, err = os.Open(srcImagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcImageFile.Close();

	outputFile, err = os.Create("output.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	format, orientation, err := sqrizrlib.Sqrize(srcImageFile, outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("processed file: %s\nformat: %s\norientation: %s", srcImagePath, format, orientation))
}

