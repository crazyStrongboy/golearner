package controller

import (
	"github.com/emicklei/go-restful"
	"eyecool.com/node-retrieval/http/buz"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type OrigImageController struct {
}

var origImageLogic = new(buz.OrigImageService)
//获取摄像头抓拍数据
func (this *OrigImageController) GetCaptureImage(req *restful.Request, rsp *restful.Response) {
	response := new(buz.OrigImageResponse)
	request := new(buz.OrigImageRequest)
	body, _ := ioutil.ReadAll(req.Request.Body)
	err := json.Unmarshal(body, request)
	if err != nil {
		fmt.Println("PictureSynchronized Unmarshal  err : ", err, ":", request)
		response.Rtn = -1
		response.Message = err.Error()
		responseBytes, _ := json.Marshal(response)
		rsp.Header().Set("Access-Control-Allow-Origin", "*")
		rsp.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
		rsp.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
		rsp.Header().Set("Access-Control-Max-Age", "1800"); //30 min
		rsp.ResponseWriter.Write(responseBytes)
		return
	}
	sessionId := req.HeaderParameter("session_id")
	user := cacheMap.GetUserSession(sessionId)
	if user != nil {
		response = origImageLogic.GetCaptureImage(request)
	} else {
		response.Rtn = -1
		response.Message = "用户未登录!"
	}
	rsp.Header().Set("Access-Control-Allow-Origin", "*")
	rsp.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
	rsp.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
	rsp.Header().Set("Access-Control-Max-Age", "1800"); //30 min
	responseBytes, _ := json.Marshal(response)
	rsp.ResponseWriter.Write(responseBytes)
}

//获取单个图片
func (this *OrigImageController) GetSingleImage(req *restful.Request, rsp *restful.Response) {
	response := new(buz.OrigImageResponse)
	request := new(buz.OrigImageRequest)
	body, _ := ioutil.ReadAll(req.Request.Body)
	err := json.Unmarshal(body, request)
	if err != nil {
		fmt.Println("PictureSynchronized Unmarshal  err : ", err, ":", request)
		return
	}
	sessionId := req.HeaderParameter("session_id")
	user := cacheMap.GetUserSession(sessionId)
	if user != nil {
		response = origImageLogic.GetSingleImage(request)
	} else {
		response.Rtn = -1
		response.Message = "用户未登录!"
	}
	rsp.Header().Set("Access-Control-Allow-Origin", "*")
	rsp.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
	rsp.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
	rsp.Header().Set("Access-Control-Max-Age", "1800"); //30 min
	responseBytes, _ := json.Marshal(response)
	rsp.ResponseWriter.Write(responseBytes)
}
