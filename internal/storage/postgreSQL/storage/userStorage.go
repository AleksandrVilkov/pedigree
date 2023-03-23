package storage

import (
	"fmt"
	"pedigree/internal/storage/postgreSQL"
	"pedigree/internal/usecase"
)

const (
	TABLE_NAME = "CLIENT "

	CREATED_DATE_COLUMN = "createdDate"
	LAST_UPDATED_COLUMN = "lastUpdatedDate"
	ROLE_COLUMN         = "Role"
	FIRST_NAME_COLUMN   = "FirstName"
	LAST_NAME_COLUMN    = "LastName"
	LOGIN_COLUMN        = "Login"
	PASS_COLUMN         = "Password"
	HAS_PEDIGREE        = "HasPedigree"
)

type UserStorage struct {
	Psql postgreSQL.PostgreSQL
}

func (us *UserStorage) CreateUser(user *usecase.User) (int, error) {
	ud, err := postgreSQL.ToUserData(user)
	if err != nil {
		return 0, err
	}

	res, err := us.Psql.SendQuery(us.getCreateUserQuery(ud))
	li, _ := res.LastInsertId()
	fmt.Print(li)
	if err != nil {
		return 0, err
	}
	//TODO return id from psql
	return 0, nil
}

func (us *UserStorage) getCreateUserQuery(ud *postgreSQL.UserData) string {
	return postgreSQL.INSERT_INTO + TABLE_NAME +
		"(" + CREATED_DATE_COLUMN + ", " +
		LAST_UPDATED_COLUMN + ", " +
		ROLE_COLUMN + ", " +
		FIRST_NAME_COLUMN + " ," +
		LAST_NAME_COLUMN + ", " +
		LOGIN_COLUMN + ", " +
		PASS_COLUMN + ", " +
		HAS_PEDIGREE + ") " +
		postgreSQL.VALUES + "(" +
		" '" + ud.CreatedDate + "', " +
		" '" + ud.LastUpdatedDate + "', " +
		" '" + ud.Role + "', " +
		" '" + ud.FirstName + "', " +
		" '" + ud.LastName + "', " +
		" '" + ud.Login + "', " +
		" '" + ud.Password + "', " +
		" '" + ud.HasPedigree + "')"
}
func (us *UserStorage) ReadUserByUserName(username string) (usecase.User, error) {
	q := us.getReadUserByUserNameQuery(username)
	res := us.Psql.GetRow(q)
	var ud postgreSQL.UserData
	err := res.Scan(&ud.ID, &ud.CreatedDate, &ud.LastUpdatedDate, &ud.Role, &ud.FirstName, &ud.LastName, &ud.Login, &ud.Password, &ud.HasPedigree)
	if err != nil {
		return usecase.User{}, nil
	}
	return *ud.ToModel(), nil
}

func (us *UserStorage) getReadUserByUserNameQuery(username string) string {
	return postgreSQL.SELECT + postgreSQL.ALL + postgreSQL.FROM + TABLE_NAME + postgreSQL.WHERE +
		"Login = '" + username + "'"
}
