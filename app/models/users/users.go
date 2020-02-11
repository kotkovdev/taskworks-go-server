package tasks

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Token    string
}

func SaveUser(user *User) *User {
	return user
}
