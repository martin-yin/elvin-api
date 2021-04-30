package middleware

import (
	"bytes"
	"danci-api/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
		publicFiles.IP = "58.243.220.37"
		fmt.Println(publicFiles.IP, "publicFiles.IP")
		context.Set("public_files", publicFiles)
		context.Next()
	}
}
