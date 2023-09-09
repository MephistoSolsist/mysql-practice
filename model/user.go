package model

type User struct {
	UserId    string `json:"userId" gorm:"PRIMARY_KEY"`
	Password  string `json:"password"`
	Alias     string `json:"alias"`
	Role      string `json:"role"`
	GmtCreate string `json:"gmtCreate" gorm:"column:gmt_create"`
}

func (*User) TableName() string {
	return "user"
}
