package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"sort"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "File size converter")
	sorted = flag.Bool("sort", false, "Output through sorting")
)


type SortedBySize struct {
	files []os.FileInfo
}

func (ss *SortedBySize) Len() int {
	return len(ss.files)
}

func (ss *SortedBySize) Less(i int, j int) bool {
	return ss.files[i].Size() < ss.files[j].Size()
}

func (ss *SortedBySize) Swap(i int, j int) {
	ss.files[i], ss.files[j] = ss.files[j], ss.files[i]
}


func hrSize(fsize int) string {
	if fsize < 1048576 {
		fsize = int(math.Ceil(float64(fsize) / 1024))
		return strconv.Itoa(fsize) + "KB"
	} else {
		fsize = int(math.Ceil(float64(fsize) / (1048576)))
		return strconv.Itoa(fsize) + "MB"
	}
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize := strconv.Itoa(int(file.Size()))
	if *h {
		fSize = hrSize(int(file.Size()))
	}
	fmt.Printf("%s %s %s \n", fSize, time, file.Name())
}


func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)
	sortedFiles := SortedBySize{files}
	if *sorted == true {
		sort.Sort(&sortedFiles)
	}

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
