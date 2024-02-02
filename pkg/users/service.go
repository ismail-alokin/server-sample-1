package users

type User struct {
	name string
	id   int
}

func (s User) GetUserById(id int) (*User, error) {
	user1 := User{"ismail", 1}
	return &user1, nil
}

func GetUserImage() {

}
