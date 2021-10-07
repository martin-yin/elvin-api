package services

import (
	"dancin-api/global"
	"dancin-api/model"
)

func GetTeamList() (team []model.Team, err error) {
	err = global.GORMDB.Preload("Admins").Preload("Projects").Model(&model.Team{}).Find(&team).Error
	return
}

func CreateTeam(team model.Team) (err error) {
	err = global.GORMDB.Model(&model.Team{}).Create(&team).Error
	return
}

func FindTeam(id uint) (team model.Team, err error) {
	err = global.GORMDB.Model(&model.Team{}).Where("id = ? ", id).Find(&team).Error
	return
}

func UpdateTeamAdminIds(team *model.Team) (err error) {
	err = global.GORMDB.Model(&team).Association("Admins").Replace(&team.Admins)
	return
}
