package common

// Captcha 验证码响应
type Captcha struct {
	Img interface{} `json:"img"` //数据内容
	Key string      `json:"key"` //验证码ID
}

// TreeData 通用的树形结构
type TreeData struct {
	Id    string `json:"id"`    //节点ID
	Pid   string `json:"pId"`   //节点父ID
	Name  string `json:"name"`  //节点名称
	Title string `json:"title"` //节点标题
}
