package storage

import (
	"errors"
	"fmt"
	"pedigree/internal/storage/postgreSQL"
	"pedigree/internal/usecase"
	"strconv"
)

const (
	USER_TABLE_NAME = "CLIENT "

	U_CREATED_DATE_COLUMN = "createdDate"
	U_LAST_UPDATED_COLUMN = "lastUpdatedDate"
	U_ROLE_COLUMN         = "Role"
	U_FIRST_NAME_COLUMN   = "FirstName"
	U_LAST_NAME_COLUMN    = "LastName"
	U_LOGIN_COLUMN        = "Login"
	U_PASS_COLUMN         = "Password"
	U_HAS_PEDIGREE        = "HasPedigree"
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
	usr, errRead := us.ReadUserByUserName(user.Login)
	if errRead != nil {
		return 0, errors.New("user was created, but it was not possible to get ID")
	}
	id, errId := strconv.Atoi(usr.ID)
	if errId != nil {
		return 0, errors.New("not valid user ID")
	}
	return id, nil
}

func (us *UserStorage) getCreateUserQuery(ud *postgreSQL.UserData) string {
	return postgreSQL.INSERT_INTO + USER_TABLE_NAME +
		"(" + U_CREATED_DATE_COLUMN + ", " +
		U_LAST_UPDATED_COLUMN + ", " +
		U_ROLE_COLUMN + ", " +
		U_FIRST_NAME_COLUMN + " ," +
		U_LAST_NAME_COLUMN + ", " +
		U_LOGIN_COLUMN + ", " +
		U_PASS_COLUMN + ", " +
		U_HAS_PEDIGREE + ") " +
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
	err := res.Scan(&ud.ID,
		&ud.CreatedDate,
		&ud.LastUpdatedDate,
		&ud.Role,
		&ud.FirstName,
		&ud.LastName,
		&ud.Login,
		&ud.Password,
		&ud.HasPedigree)
	if err != nil {
		return usecase.User{}, nil
	}
	return *ud.ToModel(), nil
}

func (us *UserStorage) getReadUserByUserNameQuery(username string) string {
	return postgreSQL.SELECT + postgreSQL.ALL + postgreSQL.FROM + USER_TABLE_NAME + postgreSQL.WHERE +
		"Login = '" + username + "'"
}
