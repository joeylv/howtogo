package main

import (
	"bytes"
	"fmt"
	"github.com/chai2010/guetzli-go"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"strconv"

	//"math"
	"os"
	"sync"
)

func main() {
	//var ret *int
	scanDir("E:\\mzi", 0)
	//var wg sync.WaitGroup
	//var count, numCpu = 0, runtime.NumCPU()
	//list := make([]os.FileInfo, numCpu)
	//iDir, oDir := "E:\\mzi\\tuku_1", "E:\\mzi\\tuku"
	//dir, err := ioutil.ReadDir(iDir)
	//
	//if err != nil {
	//	//fmt.Println("WWW")
	//}
	//for _, fi := range dir {
	//	println(count)
	//	if fi.IsDir() { // 目录, 递归遍历
	//		//dbcon.Insert("Path", strconv.Itoa(count), fi.Name(), thumbPath, )
	//		//err := os.Rename(dirPth+PthSep+fi.Name(), newPath)
	//		//checkErr(err)
	//		//GetFilesAndDirs(newPath)
	//	}
	//
	//	//fmt.Println(fi.Name())
	//	list[count] = fi
	//	count++
	//	if math.Mod(float64(count), float64(numCpu)) == 0 {
	//		guetzliEncode4(iDir, oDir, &wg, list)
	//		count = 0
	//	}
	//
	//}

}

func scanDir(iDir string, count int) {
	dir, err := ioutil.ReadDir(iDir)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	numCpu := runtime.NumCPU()
	list := make([]os.FileInfo, numCpu)
	oDir := "E:\\mzigu\\gku_"

	for _, di := range dir {
		gCount := 0
		if di.IsDir() { // 目录, 递归遍历
			count += 1
			//scanDir(iDir + "\\" + fi.Name(), count)
			files, err := ioutil.ReadDir(iDir + "\\" + di.Name())
			if err != nil {
				panic(err)
			}
			for _, fi := range files {
				list[gCount] = fi
				gCount++
				if math.Mod(float64(gCount), float64(numCpu)) == 0 {
					guetzliEncode4(iDir+"\\"+di.Name(), oDir+strconv.Itoa(count), &wg, list)
					gCount = 0
				}

			}

		}
	}
}

func guetzliEncode4(iDir string, oDir string, wg *sync.WaitGroup, list []os.FileInfo) {
	PthSep := string(os.PathSeparator)
	wg.Add(len(list))
	for _, val := range list {
		println(val.Name())
		go func(val os.FileInfo) {
			defer wg.Done()
			//fmt.Println(iDir + PthSep + val.Name())
			m0 := loadImage(iDir + PthSep + val.Name())
			err := os.MkdirAll(oDir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("guetzliEncode")
			data2 := guetzliEncode(m0, 84)

			//fmt.Println(newPath + PthSep + val.Name())
			//fmt.Println("WriteFile::::::::::" + newPath + PthSep + val.Name())
			if err := ioutil.WriteFile(oDir+PthSep+val.Name(), data2, 0666); err != nil {
				log.Println(err)
			}
		}(val)
	}
	wg.Wait()
}
func loadImage(name string) image.Image {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func jpegEncode(m image.Image, quality int) []byte {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, m, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func guetzliEncode(m image.Image, quality int) []byte {
	var buf bytes.Buffer
	err := guetzli.Encode(&buf, m, &guetzli.Options{Quality: quality})
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
