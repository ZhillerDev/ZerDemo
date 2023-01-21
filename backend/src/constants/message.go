package constants

var MessageDefine = map[int]string{
	SUCCESS:  "请求成功！",
	REDIRECT: "重定向错误",
	FAILED:   "文件丢失或请求失败",
	ERROR:    "后端处理错误",

	ERROR_UPLOAD_SAVE_IMAGE: "图片上传失败",
	ERROR_UPLOAD_TYPE_IMAGE: "图片类型错误",
}

func GetMessage(codeId int) string {
	message, ok := MessageDefine[codeId]
	if ok {
		return message
	}
	return MessageDefine[FAILED]
}
