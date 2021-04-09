package middleware

import (
	"bytes"
	"danci-api/model"
	"danci-api/model/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func PublicFields() gin.HandlerFunc {
	return func(context *gin.Context) {
		var body []byte
		if context.Request.Body != nil {
			body, _ = ioutil.ReadAll(context.Request.Body)
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		var publicFiles model.PublicFiles
		_ = json.Unmarshal(body, &publicFiles)
		ip := context.ClientIP()
		resp, _ := http.Get("https://apis.map.qq.com/ws/location/v1/ip?ip="+ip+"&key=TFNBZ-STIKX-JQ242-TNUNK-4NWCT-CLF7S")
		txMapbody, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var txMapResponse response.TxMapResponse
		err = json.Unmarshal(txMapbody, &txMapResponse)
		fmt.Println(err)
		publicFiles.IP = ip
		publicFiles.Nation = txMapResponse.Result.AdInfo.Nation
		publicFiles.Province = txMapResponse.Result.AdInfo.Province
		publicFiles.City = txMapResponse.Result.AdInfo.City
		publicFiles.District = txMapResponse.Result.AdInfo.District
		context.Set("public_files", publicFiles)
		context.Next()
	}
}
