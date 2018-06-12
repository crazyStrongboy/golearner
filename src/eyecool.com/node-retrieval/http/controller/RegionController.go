package controller

import (
	"github.com/emicklei/go-restful"
	"log"
	"eyecool.com/node-retrieval/http/buz"
	"encoding/json"
	"eyecool.com/node-retrieval/utils"
	"io/ioutil"
	"eyecool.com/node-retrieval/model"
	"fmt"
)

type RegionController struct{}

func (this *RegionController) QueryRegion(req *restful.Request, res *restful.Response) {
	log.Print("Received RegionController.QueryRegion API request : ", req.Request.RemoteAddr)
	sessionId := req.HeaderParameter("session_id")
	cacheMap := utils.CacheMap{}
	flag := cacheMap.CheckSession(sessionId)
	flag = true
	if flag {
		//查询数据库
		result := buz.QueryRegion()
		fmt.Println(req.Request.Method)
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
		res.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
		res.Header().Set("Access-Control-Max-Age", "1800"); //30 min
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	} else {
		result := &buz.RegionResponse{}
		result.Rtn = -1
		result.Message = "用户未登录"
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	}
}

func (this *RegionController) InsertRegion(req *restful.Request, res *restful.Response) {
	log.Print("Received RegionController.InsertRegion API request : ", req.Request.RemoteAddr)
	sessionId := req.HeaderParameter("session_id")
	cacheMap := utils.CacheMap{}
	flag := cacheMap.CheckSession(sessionId)
	flag = true
	if flag {
		region := buz.RegionRequest{}
		body, _ := ioutil.ReadAll(req.Request.Body)
		err := json.Unmarshal(body, &region)
		if err != nil {
			log.Println("InsertRegion err:", err)
			result := &buz.InsertRegionResponse{}
			result.Rtn = -1
			result.Message = "参数错误！"
			responseBytes, _ := json.Marshal(result)
			res.ResponseWriter.Write(responseBytes)
			return
		}
		//入库
		result := buz.InsertRegion(&region)
		fmt.Println(req.Request.Method)
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
		res.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
		res.Header().Set("Access-Control-Max-Age", "1800"); //30 min
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	} else {
		result := &buz.InsertRegionResponse{}
		result.Rtn = -1
		result.Message = "用户未登录"
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	}
}

func (this *RegionController) UpdateRegion(req *restful.Request, res *restful.Response) {
	log.Print("Received RegionController.UpdateRegion API request : ", req.Request.RemoteAddr)
	sessionId := req.HeaderParameter("session_id")
	cacheMap := utils.CacheMap{}
	flag := cacheMap.CheckSession(sessionId)
	flag = true
	if flag {
		region := buz.RegionRequest{}
		body, _ := ioutil.ReadAll(req.Request.Body)
		err := json.Unmarshal(body, &region)
		if err != nil {
			log.Println("InsertRegion err:", err)
			result := &model.RespMsg{}
			result.Rtn = -1
			result.Message = "参数错误！"
			responseBytes, _ := json.Marshal(result)
			res.ResponseWriter.Write(responseBytes)
			return
		}
		//更新
		result := buz.UpdateRegion(&region)
		fmt.Println(req.Request.Method)
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
		res.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
		res.Header().Set("Access-Control-Max-Age", "1800"); //30 min
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	} else {
		result := &model.RespMsg{}
		result.Rtn = -1
		result.Message = "用户未登录"
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	}
}

func (this *RegionController) DeleteRegion(req *restful.Request, res *restful.Response) {
	log.Print("Received RegionController.DeleteRegion API request : ", req.Request.RemoteAddr)
	sessionId := req.HeaderParameter("session_id")
	cacheMap := utils.CacheMap{}
	flag := cacheMap.CheckSession(sessionId)
	flag = true
	if flag {
		m := req.Request.URL.Query()
		id := m.Get("id")
		//删除
		result := buz.DeleteRegion(id)
		fmt.Println(req.Request.Method)
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
		res.Header().Set("Access-Control-Allow-Headers", "x-requested-with");
		res.Header().Set("Access-Control-Max-Age", "1800"); //30 min
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	} else {
		result := &model.RespMsg{}
		result.Rtn = -1
		result.Message = "用户未登录"
		responseBytes, _ := json.Marshal(result)
		res.ResponseWriter.Write(responseBytes)
	}
}