package models

//NameStep 是 最基本的元素,定义了某一时刻一个人及其对应步数
type NameStep struct {
	Name string
	Step int
}

//Record 定义了信息采集设备一次提交的数据
type Record struct {
	NameStep  []NameStep
	TimeStamp int64
}

//PublicRecord 共享的主结构体
var PublicRecord []Record

