package auth

type User struct {
	ID			int64	`gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username	string	`gorm:"column:username"`
	Password	string	`gorm:"column:password"`
}
