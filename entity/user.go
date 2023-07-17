package entity

type User struct {
	Base      `bson:",inline"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	UserTopic string `json:"user_topic"`
}

func (User) TableName() string {
	return "user"
}

func (User) IDField() string {
	return "id"
}
