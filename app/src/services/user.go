package services

import "app/models"

func GetUser(userid string) (models.User, error) {
	return models.GetUser(userid)
}

func GetDiscordUser(discordid string) (models.User, error) {
	return models.GetFromDiscordId(discordid)
}