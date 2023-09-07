package model

type User struct {
	UserId    string `json:"userId"`
	Password  string `json:"password"`
	Alias     string `json:"alias"`
	GmtCreate string `json:"gmtCreate" gorm:"column:gmt_create"`
}

/*
{
    "userId":"1",
    "password":"1",
    "alias":"y",
    "gmtCreate":"2022-01-02"
}
*/

func (*User) TableName() string {
	return "user"
}
