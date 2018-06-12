package buz

import (
	"time"
	"fmt"
	"os"
	"eyecool.com/node-retrieval/utils"
	"encoding/base64"
	"io/ioutil"
	"eyecool.com/node-retrieval/algorithm"
	. "github.com/polaris1119/config"
)

type VerifyReponse struct {
	Rtn int `json:"rtn"`
	Message string `json:"message"`
	Similarity float64 `json:"similarity"`
}

type VerifyRequest struct {
	ImageBase64_1 string   `json:"image_base64_1"`//图片base64编码 
	ImageBase64_2 string	`json:"image_base64_2"`
	ImageType_1 int `json:"image_type_1"`
	ImageType_2 int  `json:"image_type_2"`//图片类型
}

func FaceVerify(verify *VerifyRequest)*VerifyReponse{
	result:=&VerifyReponse{}
	path:=time.Now().Format("20060102")
	imagePath, err := ConfigFile.GetSection("path")
	if err!=nil{
		fmt.Println("VerifyService read file err:",err)
		result.Rtn=-1
		result.Message="比对失败"
		return result
	}
	targetPath:=imagePath["people_target_path"]+"/verify/"+path
	if _, err := os.Stat(targetPath); err != nil {
		if os.IsNotExist(err){
			os.MkdirAll(targetPath, os.ModePerm)
		}
	}
	imageName1 := utils.UUID() + ".jpg"
	imageName2:=utils.UUID()+".jpg"

	realPath1:=targetPath+"/"+imageName1
	realPath2:=targetPath+"/"+imageName2

	data1, _ := base64.StdEncoding.DecodeString(verify.ImageBase64_1)
	err= ioutil.WriteFile(realPath1, data1, os.ModePerm)
	data2, _ := base64.StdEncoding.DecodeString(verify.ImageBase64_2)
	err= ioutil.WriteFile(realPath2, data2, os.ModePerm)
	if err != nil {
		fmt.Println("verify insert write error", err, "path:", realPath1)
		result.Rtn = -1
		result.Message = "文件写入失败"
		return result
	}

	//检测是否有人脸
	_, width1, height1, rgb24Data1 := algorithm.NewChlFaceX().ReadImageFile(realPath1, 0, 0)
	hasFace1, faceResult1 := algorithm.NewChlFaceX().ChlFaceSdkDetectFace(-1, rgb24Data1, width1, height1, true)
	_, width2, height2, rgb24Data2 := algorithm.NewChlFaceX().ReadImageFile(realPath2, 0, 0)
	hasFace2, faceResult2 := algorithm.NewChlFaceX().ChlFaceSdkDetectFace(-1, rgb24Data2, width2, height2, true)
	if hasFace1 == 0 || faceResult1 == nil|| hasFace2 == 0 || faceResult2 == nil {
		//检测不到人脸
		result.Rtn = -1
		result.Message = "检测不到人脸"
		return result
	}
	//提取特征值
	_,feature1:=algorithm.NewChlFaceX().ChlFaceSdkFeatureGet(0,rgb24Data1, width1, height1,faceResult1)
	_,feature2:=algorithm.NewChlFaceX().ChlFaceSdkFeatureGet(0,rgb24Data2, width2, height2,faceResult2)
	//比较
	similarity:=algorithm.NewChlFaceX().ChlFaceSdkFeatureCompare(0,feature1,feature2)
	result.Similarity=float64(similarity)
	result.Rtn=0
	result.Message="比对成功！"
	return result
}
