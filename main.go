package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"math"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "File size converter")
)

func hrSize(fsize int) string {
	if fsize < 1048576 {
		finalFsize := int(math.Ceil(float64(fsize)/1024))
		return strconv.Itoa(finalFsize) + "KB"
	} else {
		finalFsize := int(math.Ceil(float64(fsize)/(1048576)))
		return strconv.Itoa(finalFsize) + "MB"
	}
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize := int(file.Size())
	if *h {
		fmt.Printf("%s %s %s \n", hrSize(fSize), time, file.Name())
	} else {
		fmt.Printf("%s %s %s \n", fSize, time, file.Name())
	}
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
