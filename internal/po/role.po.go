package po

func init() {
	Register(&Role{})
}

type Role struct {
	Base
	RoleName string `gorm:"type:varchar(255);not null"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
