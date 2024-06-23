package user

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("{ID: %d, Name: %s, Age: %d}", u.ID, u.Name, u.Age)
}
