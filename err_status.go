package gacos

func errStatusTostring(errcode int) string {
	switch errcode {
	case 400:
		return "客户端请求中的语法错误"
	case 403:
		return "没有权限"
	case 404:
		return "无法找到资源"
	case 500:
		return "服务器内部错误"
	default:
		return "未知错误"
	}
}