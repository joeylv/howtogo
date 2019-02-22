package main

import (
	"./dbcon"
	"fmt"
	_ "github.com/chai2010/guetzli-go"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/yanyiwu/gojieba"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"thumbnail"
)

//var count, newPath, thumbPath = 0, "", ""
var wg sync.WaitGroup
var count = 0

func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	checkErr(err)
	//gCount, numCpu := 0, runtime.NumCPU()
	//list := make([]os.FileInfo, numCpu)
	//if err != nil {
	//	return nil, nil, err
	//}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		dirs = append(dirs, dirPth+PthSep+fi.Name())

		//gPath := "E:/mzig/tuku_" + strconv.Itoa(count)

		//fmt.Println(dirs)
		if fi.IsDir() { // 目录, 递归遍历

			count++
			newPath := dirPth + "/tuku_" + strconv.Itoa(count)
			//thumbPath := dirPth + "/thumb_" + strconv.Itoa(count)
			//err := os.MkdirAll(newPath, os.ModePerm)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//err = os.MkdirAll(thumbPath, os.ModePerm)

			//if err != nil {
			//	fmt.Println(err)
			//}

			fmt.Println(newPath)
			err = os.Rename(dirPth+PthSep+fi.Name(), newPath)
			checkErr(err)
			dbcon.Insert("Path", strconv.Itoa(count), fi.Name(), "tuku_"+strconv.Itoa(count))
			GetFilesAndDirs(newPath)
		} else {
			tuku := "tuku_" + strconv.Itoa(count)
			//tuku := "tuku_" + strconv.Itoa(count)
			//newPath := "F:/mzi/tuku_" + strconv.Itoa(count)

			//thumbPath := dirPth + thumb
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".jpg")
			//fmt.Println(count)
			//fmt.Println(thumbPath)
			//fmt.Println(fi.Name())
			if ok {
				//fmt.Println(fi.Sys())
				//list[gCount] = fi
				//gCount++
				//if math.Mod(float64(gCount), float64(numCpu)) == 0 {
				//	guetzliEncode4(newPath, gPath, &wg, list)
				//	gCount = 0
				//}
				//fmt.Println(dirPth)
				//fmt.Println(dirPth[strings.LastIndex(dirPth, "\\")+1:])

				dbcon.Insert("Top", fi.Name(), tuku)

				// thumbnail 缩略图
				_, err := thumbnail.ImageFile(dirPth + PthSep + fi.Name()) //, thumbPath+PthSep+fi.Name()
				if err != nil {
					log.Print(err)
					continue
				}
				//fmt.Println(thumb)

				//time.Sleep(time.Second)

				//fmt.Println("Mkdir")
				//fmt.Println("guetzliSave")

			}
		}
	}
	return files, dirs, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//start("E:\\mzitu")
	GetFilesAndDirs("F:/mzitu")
	//list := []string{""}
	//words := []string{""}
	//x := gojieba.NewJieba()
	//defer x.Free()
	////var words []string
	//words := x.Cut("大胸长腿清新可人 养眼萌妹子夏美酱清纯又不失性感", false)
	//word2 := x.Cut("风骚欲女周于希奶大屁股翘 脱衣玩自摸豪放大胆", false)
	//n_w := append(words, word2...)
	//fmt.Print(n_w)

	//result := dbcon.Search("dir_info", 3, 0)
	//for result.Next() {
	//	name, dirName := "", ""
	//	result.Scan(&name, &dirName)
	//	var dirPth = "d:/mzitu/"
	//	//println(dirPth + "tuku_" + name)
	//	files, dirs, _ := GetFilesAndDirs(dirPth + "tuku_" + name)
	//	fmt.Println(files, dirs)
	//	//fmt.Println(name, dirName)
	//}
	//dir := module.DirInfo{Id: 1}
	//
	//dbcon.Get(&dir)

	//i := append(list, words)
	//fmt.Println(i)

	//pwd, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//fmt.Println(pwd)
	//dbcon.Insert(1, "大胸长腿清新可人 养眼萌妹子夏美酱清纯又不失性感", )
	//fmt.Println(time.Now().Local())
	//db, err := sql.Open("sqlite3", pwd+"/foo.db")
	//checkErr(err)
	//println(db)
	//db.Exec(`
	//CREATE TABLE IF NOT EXISTS dir_info(
	//    uid INTEGER PRIMARY KEY AUTOINCREMENT,
	//	name VARCHAR(64) NULL,
	//    dirName VARCHAR(256) NULL,
	//    created DATE NULL
	//);
	//`)
	//db.Close()

	//stmt, err := db.Prepare("INSERT INTO dir_info(uid, name, dirName, created) values(?,?,?,?)")
	//checkErr(err)
	//
	//_, err = stmt.Exec(1, strconv.Itoa(1), "暧昧气氛撩人心弦 寂寞熟女雪瑞Lisa透明情趣睡衣秀出魔鬼身材", time.Now().Local())
	//checkErr(err)
	//fmt.Println(time.Now().Local())
}
