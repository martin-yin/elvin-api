package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetTeamList(context *gin.Context) {
	//claims, exists := context.Get("claims")
	//if exists {
	//	var customClaims request.CustomClaims
	//utils.InterfaceToJsonToStruct(claims, &customClaims)
	responses, _ :=services.GetTeamList()
	response.OkWithDetailed(responses, "获取成功", context)

}

func CreateTeam(context *gin.Context) {
	services.CreateTeam()
}

