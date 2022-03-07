package v1

import (
	"dancin-api/global"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/utils"
	"go.uber.org/zap"

	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var handles *utils.Handles

func init() {
	handles = utils.NewHandles()
	routerHandles := map[string]utils.RouterFunc{
		"PERFORMANCE": func(context *gin.Context) {
			var performanceBody request.PerformanceBody
			err := context.ShouldBindJSON(&performanceBody)
			performanceBody.IP = context.ClientIP()
			reportProducer(context, performanceBody, err)
			return
		},
		"HTTPLOG": func(context *gin.Context) {
			var httpBody request.HttpBody
			err := context.ShouldBindJSON(&httpBody)
			httpBody.IP = context.ClientIP()
			reportProducer(context, httpBody, err)
			return
		},
		"PAGEVIEW": func(context *gin.Context) {
			var pageViewBody request.PageViewBody
			err := context.ShouldBindJSON(&pageViewBody)
			pageViewBody.IP = context.ClientIP()
			reportProducer(context, pageViewBody, err)
			return
		},
		"OPERATION": func(context *gin.Context) {
			var operationBody request.OperationBody
			err := context.ShouldBindJSON(&operationBody)
			operationBody.IP = context.ClientIP()
			reportProducer(context, operationBody, err)
			return
		},
		"RESOURCE": func(context *gin.Context) {
			var resourceBody request.ResourceErrorBody
			err := context.ShouldBindJSON(&resourceBody)
			resourceBody.IP = context.ClientIP()
			reportProducer(context, resourceBody, err)
			return
		},
		"JSERROR": func(context *gin.Context) {
			var issuesBody request.IssuesBody
			err := context.ShouldBindJSON(&issuesBody)
			issuesBody.IP = context.ClientIP()
			reportProducer(context, issuesBody, err)
			return
		},
	}
	handles.RoutersHandlerRegister(routerHandles)
}

func Report(context *gin.Context) {
	actionType := context.Query("action_type")
	handles.RouterHandlers[actionType](context)
}

// 数据消费
func reportProducer(context *gin.Context, body interface{}, err error) {
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	sessionId := context.Query("session_id")
	report, _ := json.Marshal(body)
	if global.KAFKA != nil { // 判断下 kafka 是否初始化成功，失败则走redis
		msg := kafka.Message{
			Key:   []byte(fmt.Sprint(sessionId)),
			Value: report,
		}
		_, err = global.KAFKA.WriteMessages(msg)
		if err != nil {
			global.LOGGER.Error("kafka 写入数据失败:", zap.Any("err", err))
		}
	} else {
		global.REDIS.LPush("reportData", report)
	}
	response.Ok(context)
	return
}
