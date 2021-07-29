package models

type Model interface {

	//获取一条记录
	GetOne(interface{},[]string) interface{}

}
