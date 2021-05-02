package response

import "github.com/gogf/gf/net/ghttp"

type JsonResponse struct {
	Code    int         `json:"code"` // (0:成功；1-失败；>1：错误码
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(r *ghttp.Request, errCode int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(JsonResponse{
		Code:    errCode,
		Message: message,
		Data:    responseData,
	})
}

// 返回json数据并终止http执行
func JsonExit(r *ghttp.Request, errCode int, msg string, data ...interface{}) {
	Json(r, errCode, msg, data)
	r.Exit()
}
