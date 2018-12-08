package dbcon

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/go-xorm/xorm"
	"module"
	"os"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

//var db, _ = sql.Open("sqlite3", pwd+"/foo.db")
var pwd, _ = os.Getwd()
var engine *xorm.Engine

func init() {
	var err error
	//与数据库建立链接
	engine, err = xorm.NewEngine("sqlite3", pwd+"/foo.db", )
	checkErr(err)
	//在控制台打印sql语句，默认为false
	engine.ShowSQL(true)
	err = engine.Sync2(new(module.Path))
	err = engine.Sync2(new(module.Top))

	// set default database
	//orm.RegisterDataBase("default", "sqlite3", pwd+"/foo.db", 30)
	//orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)

	// register model
	//orm.RegisterModel(new(module.DirInfo))

	// create table
	//orm.RunSyncdb("default", false, true)
	//pwd, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

}
func Insert(arg ... string, ) {

	if arg[0] == "Path" {
		p := module.Path{
			Name: arg[1], Path: arg[2], Thumb: arg[3],
		}
		engine.Insert(&p)
	} else if arg[0] == "Top" {
		fmt.Println(arg)
		var ret module.Path
		engine.Where("path = ?", arg[2]).Get(&ret)

		//path = engine.Id(1).Get(&module.Path{})
		//fmt.Println(ret)
		t := module.Top{
			File: arg[1],
			Path: ret,
		}
		engine.Insert(&t)
	}

	//o := orm.NewOrm()
	//dir := module.DirInfo{Id: id, Name: strconv.Itoa(int(id)), Dir: dirName, Thumb: thumb}
	//id, err := o.Insert(&dir)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

func Update(dir *module.Path) {
	o := orm.NewOrm()
	//dir.Name = "astaxie"
	num, err := o.Update(&dir)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//更新数据
	//stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//
	//res, err := stmt.Exec("username", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
}

func Get(dir *module.Path) {
	o := orm.NewOrm()
	u := module.Path{Id: dir.Id}
	err := o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
}

func Search(table string, limit int, offset int) []*module.Path {
	o := orm.NewOrm()
	var posts []*module.Path
	qs := o.QueryTable("post")
	num, _ := qs.Filter("User__Name", "slene").All(&posts)
	fmt.Println(num)
	return posts
	//db, _ = sql.Open("sqlite3", pwd+"/foo.db")
	////查询数据
	//var str = fmt.Sprintf("SELECT name,dirName FROM %s LIMIT %d OFFSET %d", table, limit, offset)
	//rows, err := db.Query(str)
	//checkErr(err)
	//return rows

	//for rows.Next() {
	//	var name string
	//	var dirName string
	//	err = rows.Scan(&name, &dirName, )
	//	checkErr(err)
	//	fmt.Println(name)
	//	fmt.Println(dirName)
	//}
}

func Remove(dir *module.Path) {
	o := orm.NewOrm()

	num, err := o.Delete(dir)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//删除数据
	//stmt, err := db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err := stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	//db.Close()
}

//CreateTable
func Create() {
	//sql_table := `
	//CREATE TABLE IF NOT EXISTS dir_info(
	//    uid INTEGER PRIMARY KEY AUTOINCREMENT,
	//	name VARCHAR(64) NULL,
	//    dirName VARCHAR(256) NULL,
	//    created DATE NULL
	//);
	//`

	//db.Exec(sql_table)
	//db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
