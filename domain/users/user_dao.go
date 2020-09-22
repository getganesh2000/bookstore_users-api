package users

import (
	"fmt"
	"github.com/getganesh2000/bookstore_users-api/utils/erros"
)


var(
	usersDB = make(map[int64]*User)
)

func (user *User)Get() *erros.RestErr{
	result := usersDB[user.Id]

	if result == nil{
		return erros.NewNotFound(fmt.Sprintf("user %d is not found ",user.Id))
	}

	user.Id = result.Id
	user.FirstName= result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil

}

func (user *User)Save() *erros.RestErr{

	current := usersDB[user.Id]

	if current !=nil {
		if current.Email == user.Email{
			return erros.NewBadRequest(fmt.Sprintf("user %s already exists",user.Email))

		}
		return erros.NewBadRequest(fmt.Sprintf("user %d already exists",user.Id))
	}

	usersDB[user.Id] = user

	return nil
}