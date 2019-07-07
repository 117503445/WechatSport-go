package models

//Record 定义了信息采集设备一次提交的数据
type Record struct {
	NameStep  map[string]string
	TimeStamp int64
}

//PublicRecord 共享的主结构体
var PublicRecord []Record