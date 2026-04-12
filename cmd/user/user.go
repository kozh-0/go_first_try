package user

import "fmt"

// User структура пользователя
type User struct {
    Name string
    Age  int
}

// NewUser создаёт нового пользователя
func NewUser(name string, age int) *User {
    return &User{
        Name: name,
        Age:  age,
    }
}

// Greet приветствие от пользователя
func (u *User) Greet() {
    fmt.Printf("Привет! Я %s, мне %d лет\n", u.Name, u.Age)
}

// SetAge устанавливает возраст
func (u *User) SetAge(age int) {
    u.Age = age
}