package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alen/echo-framework/db"
	"github.com/alen/echo-framework/helper"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateConf()
	sqlStatment := "SELECT * FROM user WHERE username = ?"

	err := con.QueryRow(sqlStatment, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false, err
	}
	if err != nil {
		fmt.Println("Querry Error")
		return false, err
	}
	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("hash And Password Doesn't Math")
		return false, err
	}
	return true, nil
}

func RegisterUser(username, password string) (Response, error) {
	var res Response

	v := validator.New()
	usr := User{
		Username: username,
		Password: password,
	}
	err := v.Struct(usr)
	if err != nil {
		return res, err
	}

	h, i := helper.HashPassword(username, password)
	if i != nil {
		fmt.Println("gagal")
	}
	fmt.Println("Sukses", h)
	con := db.CreateConf()

	sqlStatment := "INSERT user (username, password ) VALUES (?, ? )"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(username, password)
	if err != nil {
		return res, err
	}

	lastinsertedID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"Last ID": lastinsertedID,
	}
	return res, nil
}
