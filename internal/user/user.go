package user

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) String() string {
	return u.Name
}
