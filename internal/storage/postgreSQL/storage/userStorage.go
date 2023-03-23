package storage

import (
	"fmt"
	"pedigree/internal/storage/postgreSQL"
	"pedigree/internal/usecase"
)

const (
	TABLE_NAME = "USER"
)

type UserStorage struct {
	Psql postgreSQL.PostgreSQL
}

func (us *UserStorage) CreateUser(user *usecase.User) (int, error) {
	//TODO
	ud, err := postgreSQL.ToUserData(user)
	if err != nil {
		return 0, err
	}

	_, err = us.Psql.SendQuery(postgreSQL.INSERT_INTO + TABLE_NAME + postgreSQL.VALUES + "(" +
		ud.CreatedDate.String() + ", " +
		ud.LastUpdatedDate.String() + ", " +
		ud.Role + ", " +
		ud.FirstName + ", " +
		ud.LastName + ", " +
		ud.Login + ", " +
		ud.Password + ", " +
		ud.HasPedigree + ")")

	if err != nil {
		return 0, err
	}
	//TODO

	return 0, nil
}
func (us *UserStorage) ReadUserByUserName(username string) (usecase.User, error) {
	res := us.Psql.GetRow(postgreSQL.SELECT + postgreSQL.ALL + postgreSQL.FROM + TABLE_NAME + postgreSQL.WHERE +
		"Login = '" + username + "'")
	//TODO
	fmt.Print(res)
	return usecase.User{}, nil
	//TODO
}
