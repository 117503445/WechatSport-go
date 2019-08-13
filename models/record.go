package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	//mysql
	_ "github.com/go-sql-driver/mysql"
)

//Record 为一次记录
type Record struct {
	TimeStamp int64
	Name      string
	Step      int
}

var db *sql.DB

//InitDatabase 初始化数据库连接 默认 port 3306
func InitDatabase(host string, port int, username string, password string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port) //Data Source Name
	fmt.Println(dsn)
	db, _ = sql.Open("mysql", dsn)
	db.Exec("Create Database If Not Exists TestDB Character Set UTF8;")
	db.Exec("create table If Not Exists TestDB.testTB(timestamp BIGINT UNSIGNED, name varchar(10), step int);")
}

//CloseDatabase 关闭数据库连接
func CloseDatabase() {
	db.Close()
}

//PostJSONToRecords 将POST得到的JSON数据转换为结构体数组
func PostJSONToRecords(js []byte) []Record {
	type temp struct {
		NameStep  map[string]string
		TimeStamp int64
	}
	var t temp
	json.Unmarshal(js, &t)
	var result []Record
	for name, step := range t.NameStep {
		//s, _ := strconv.ParseInt(step, 10, 64)
		s, _ := strconv.Atoi(step)
		r := Record{t.TimeStamp, name, s}
		result = append(result, r)
	}
	return result
}

//SubmitData 向数据库插入新的记录
func SubmitData(records []Record) {
	//"(1562503795, '林辰希', '20874'), (1562503795, '钟保明', '10127');"
	sql := "insert into TestDB.testTB values "
	for i, record := range records {
		s := fmt.Sprintf("(%d,'%s','%d')", record.TimeStamp, record.Name, record.Step)
		if i != len(records)-1 {
			s += ", "
		}
		sql += s
	}
	sql += ";"
	fmt.Println(sql)
	db.Exec(sql)
}

//GetRecords 获取记录,根据姓名和起止时间戳进行筛选
func GetRecords(name string, beginTimeStamp int64, endTimeStamp int64) []Record {
	//todo 继续优化代码
	fmt.Println(name, beginTimeStamp, endTimeStamp)
	sql := ""
	if name == "" && beginTimeStamp == 0 {
		sql = "SELECT * FROM `TestDB`.`testTB` ORDER BY `timestamp`"
	}
	if name != "" && beginTimeStamp == 0 {
		sql = "SELECT * FROM `TestDB`.`testTB` where `name`='" + name + "' ORDER BY timestamp"
	}
	if name == "" && beginTimeStamp != 0 {
		sql = "SELECT * FROM `TestDB`.`testTB` where `timestamp` between " + strconv.FormatInt(beginTimeStamp, 10) + " AND " + strconv.FormatInt(endTimeStamp, 10) + " ORDER BY `timestamp`"
	}
	if name != "" && beginTimeStamp != 0 {
		sql = "SELECT * FROM `TestDB`.`testTB` where `timestamp` between " + strconv.FormatInt(beginTimeStamp, 10) + " AND " + strconv.FormatInt(endTimeStamp, 10) + " AND `name`='" + name + "'  ORDER BY `timestamp`"
	}
	fmt.Println(sql)
	rows, _ := db.Query(sql)
	result := getRecordsFromRows(rows)
	fmt.Println(result)
	return result
}

//将返回的Rows转换为结构体数组
func getRecordsFromRows(query *sql.Rows) []Record {

	var records []Record

	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return records
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	for _, mp := range results { //查询出来的数组
		timestamp, _ := strconv.ParseInt(mp["timestamp"], 10, 64)
		step, _ := strconv.Atoi(mp["step"])
		r := Record{timestamp, mp["name"], step}
		records = append(records, r)
	}
	return records
}
