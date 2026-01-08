package po

func init() {
	Register(&User{})
}

type User struct {
	Base
	Username   string  `gorm:"type:varchar(255);not null"`
	Password   string  `gorm:"type:varchar(255);not null"`
	Email      string  `gorm:"type:varchar(255);not null"`
	Roles      []*Role `gorm:"many2many:go_db_user_role;"`
	IsVerified bool    `gorm:"type:boolean;not null"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
