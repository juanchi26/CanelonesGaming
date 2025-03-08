package bd

import (
	"fmt"

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
