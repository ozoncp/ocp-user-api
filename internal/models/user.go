package models

import (
    "fmt"
)

type User struct {
    Id uint64
    Name string
    Surname string
    Patronymic string
    Email string
}

func (u *User) ToString() string {
    return fmt.Sprintf("{ Id:%d, Name:%s, Surname:%s, Patronymic:%s, Email:%s }", u.Id, u.Name, u.Surname, u.Patronymic, u.Email)
}

func (u *User) Equals(other *User) bool {
    if other == nil || u.Id != other.Id {
        return false
    }

    return true
}
