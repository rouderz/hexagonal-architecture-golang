package domain

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

func NewGame(name string, surname string, age string) User {
	return User{
		Name:  name,
		Surname: surname,
		Age: age,
	}
}

type Users[] string