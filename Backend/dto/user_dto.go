package dto

import (
	"Siu_shop/Backend/model"
	"encoding/base64"
)

type UserDto struct {
	Id        int64  `json:"id"`
	Account   string `json:"account"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

func ToUserDto(user model.User) UserDto {
	avatarBase64 := ""
	if len(user.Avatar) > 0 {
		avatarBase64 = "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(user.Avatar))
	}
	return UserDto{
		Id:        user.Id,
		Account:   user.Account,
		Name:      user.Name,
		Telephone: user.Telephone,
		Email:     user.Email,
		Avatar:    avatarBase64,
	}
}
