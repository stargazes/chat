package controller

import "chat/tools"

func TestRedis() interface{} {

	tools.SetString("name","lyh",30)

	return tools.GetString("name")

}
