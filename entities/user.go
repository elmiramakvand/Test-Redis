package entities

type User struct {
	ID       int64
	Name     string
	Family   string
	UserName string `gorm:"column:UserName"`
	Password string
}

func (User) TableName() string {
	return "User"
}
