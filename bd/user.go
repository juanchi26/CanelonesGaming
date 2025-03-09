package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanchi26/CanelonesGaming/models"
	"github.com/juanchi26/CanelonesGaming/tools"
)

func UpdateUser(UField models.User, user string) error {
	fmt.Println("Comienza UpdateUser")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE users SET "

	coma := ""

	if len(UField.UserFirstName) > 0 {
		coma = ","
		sentencia += "User_FirstName = '" + UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0 {
		sentencia += coma + "User_LastName = '" + UField.UserLastName + "'"
	}

	sentencia += ", User_DateUpg = '" + tools.FechaMySQL() + "' WHERE User_UUID = '" + user + "'"

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("UpdateUsers > Ejecucion Exitosa")

	return nil

}

func SelectUser(userId string) (models.User, error) {
	fmt.Println("Comienza SelectUser")

	User := models.User{}

	err := DbConnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	sentencia := "SELECT * FROM users WHERE User_UUID = '" + userId + "'"

	var rows *sql.Rows

	rows, err = Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return User, err
	}
	defer rows.Close()

	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpd = dateUpg.Time.String()

	fmt.Println("SelectUser > Ejecucion Exitosa")

	return User, nil
}

func SelectUsers(Page int) (models.ListUsers, error) {
	fmt.Println("Comienza SelectUsers")

	var lu models.ListUsers

	User := []models.User{}

	err := DbConnect()
	if err != nil {
		return lu, err
	}
	defer Db.Close()

	var offset = (Page * 10) - 10
	var sentencia string
	var sentenciaCount string

	sentenciaCount = "SELECT COUNT(*) as registros FROM users"

	sentencia = "SELECT * FROM users LIMIT 10"

	if offset > 0 {
		sentencia += " OFFSET " + strconv.Itoa(offset)
	}

	var rowsCount *sql.Rows

	rowsCount, err = Db.Query(sentenciaCount)

	if err != nil {
		fmt.Println(err.Error())
		return lu, err
	}

	defer rowsCount.Close()

	rowsCount.Next()

	var registros int

	rowsCount.Scan(&registros)

	lu.TotalItems = registros

	var rows *sql.Rows

	rows, err = Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return lu, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User

		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime

		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpd = dateUpg.Time.String()

		User = append(User, u)

	}

	fmt.Println("SelectUsers > Ejecucion Exitosa")
	lu.Data = User
	return lu, nil
}
