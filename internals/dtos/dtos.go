package dtos

type UserLoginDTO struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type VideoGroupDTO struct {
	AuthorId  int    `json:"AuthorId"`
	GroupName string `json:"GroupName"`
}
