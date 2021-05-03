package services

import (
	"danci-api/global"
	"danci-api/model"
	"fmt"
)

func GetTeamList() (team []model.Team, err error) {
	err = global.GVA_DB.Preload("Admins").Preload("Projects").Model(&model.Team{}).Find(&team).Error

	return
}

func CreateTeam() {
	team := model.Team{
		Name:     "测试团队111111",
		LeaderId: 1,
		Projects: []model.Project{
			{
				ProjectName: "测试项目2",
			},
		},
	}
	err := global.GVA_DB.Omit("Admins").Model(&model.Team{}).Create(&team).Error
	fmt.Println(err, "err!!!!!!!!!!!!!")
	return
}
