package model

type User struct {
	ID       uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"-"`
	UserName string `gorm:"column:user_name;type:varchar(100);not null" json:"userName"`
	Password string `gorm:"column:password;type:char(32);not null" json:"password"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID       string
	UserName string
	Password string
}{
	ID:       "id",
	UserName: "user_name",
	Password: "password",
}
