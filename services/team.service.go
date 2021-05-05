package services

import (
	"danci-api/global"
	"danci-api/model"
)

func GetTeamList() (team []model.Team, err error) {
	err = global.GVA_DB.Preload("Admins").Preload("Projects").Model(&model.Team{}).Find(&team).Error
	return
}

func CreateTeam(team model.Team) (err error) {
	err = global.GVA_DB.Model(&model.Team{}).Create(&team).Error
	return
}

func FindTeam(id uint) (team model.Team, err error) {
	err = global.GVA_DB.Model(&model.Team{}).Where("id = ? ", id).Find(&team).Error
	return
}

func UpdateTeamAdminIds(team *model.Team) (err error) {
	err = global.GVA_DB.Model(&team).Association("Admins").Replace(&team.Admins)
	return
}
