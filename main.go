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
	sorted = flag.String("sort", "", "Output through sorting. Use 'size' or 'date' parameter")
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


type SortedByDate struct {
	files []os.FileInfo
}

func (sd *SortedByDate) Len() int {
	return len(sd.files)
}

func (sd *SortedByDate) Less(i int, j int) bool {
	return sd.files[i].ModTime().Unix() < sd.files[j].ModTime().Unix()
}

func (sd *SortedByDate) Swap(i int, j int) {
	sd.files[i], sd.files[j] = sd.files[j], sd.files[i]
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
	files, err := ioutil.ReadDir(*d)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *sorted == "size" {
		sort.Sort(&SortedBySize{files})
	} else if *sorted == "date" {
		sort.Sort(&SortedByDate{files})
	}

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
