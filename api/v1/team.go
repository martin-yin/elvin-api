package v1

import (
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"dancin-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

// 获取团队列表
func GetTeamList(context *gin.Context) {
	_, exists := context.Get("claims")
	if exists {
		responses, _ := services.GetTeamList()
		response.OkWithDetailed(responses, "获取成功", context)
	}
}

func DelTeam(context *gin.Context) {
	id, isExist := context.GetQuery("id")
	if isExist {
		paId := StrToUInt(id)
		responses, err := services.DelTeam(paId)
		if err != nil {
			response.FailWithDetailed(responses, err.Error(), context)
			return
		}
		response.OkWithDetailed(responses, "删除成功", context)
	}
}

func CreateTeam(context *gin.Context) {
	customClaims := getCustomClaims(context)
	if customClaims != nil {
		var teamParams request.AddTeamParams
		_ = context.ShouldBind(&teamParams)
		admins, err := services.FindAdmins(customClaims.ID)
		if err == nil {
			team := &model.Team{
				Name:     teamParams.Name,
				AdminId:  customClaims.ID,
				NickName: customClaims.NickName,
				Admins:   admins,
			}
			if err := services.CreateTeam(*team); err != nil {
				response.FailWithMessage("创建团队失败！！", context)
				return
			}
			response.OkWithMessage("创建团队成功！", context)
		}
	}
}

// 判断团队是否存在
func TeamIsExist(context *gin.Context) {
	var teamParams request.AddTeamParams
	_ = context.ShouldBind(&teamParams)
	teamIsExist := services.FindProject(teamParams.Name)
	if teamIsExist {
		response.FailWithMessage("团队已存在", context)
		return
	}
	response.OkWithMessage("", context)
}

func CreateProject(context *gin.Context) {
	customClaims := getCustomClaims(context)
	if customClaims != nil {
		var addTeamProjectParams request.AddTeamProjectParams
		_ = context.ShouldBind(&addTeamProjectParams)
		team, err := services.FindTeam(addTeamProjectParams.TeamId)
		if err != nil {
			response.FailWithMessage("没有查询到团队！", context)
		} else {
			monitorId := "monitor_id" + fmt.Sprintf("%d", int32(time.Now().Unix()))
			projectModel := model.Project{
				ProjectName: addTeamProjectParams.ProjectName,
				ProjectType: addTeamProjectParams.ProjectType,
				Logo:        addTeamProjectParams.Logo,
				MonitorId:   monitorId,
				AdminID:     customClaims.ID,
				TeamID:      team.ID,
			}
			err := services.CreateProject(projectModel)
			if err != nil {
				response.FailWithMessage(err.Error(), context)
				return
			}
			response.OkWithDetailed("project", "项目创建成功！", context)
		}
	}
}

func getCustomClaims(context *gin.Context) (customClaims *request.CustomClaims) {
	claims, exists := context.Get("claims")
	if exists {
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		return customClaims
	}
	return nil
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// 给团队绑定管理员
func BindTeamAdmins(context *gin.Context) {
	var teamAdminsParams request.BindTeamAdminsParams
	_ = context.ShouldBind(&teamAdminsParams)
	team, err := services.FindTeam(teamAdminsParams.TeamId)
	if err != nil {
		response.FailWithMessage("没有查询到团队！", context)
	} else {
		// 接受需要被绑定的用户ids
		admins, err := services.FindAdmins(teamAdminIds(&team, strings.Split(teamAdminsParams.AdminIds, ","))...)
		if err == nil {
			team.Admins = admins
			if err := services.UpdateTeamAdminIds(&team); err != nil {
				response.FailWithMessage("绑定团队成员失败！", context)
				return
			}
			response.OkWithMessage("绑定团队成员成功！", context)
		}
	}
}

func teamAdminIds(team *model.Team, paramAdminIds []string) (adminIds []uint) {
	adminIds = []uint{team.AdminId}
	for _, value := range paramAdminIds {
		adminIds = append(adminIds, StrToUInt(value))
	}
	return
}
