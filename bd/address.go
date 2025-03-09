package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanchi26/CanelonesGaming/models"
)

func InsertAddress(t models.Address, user string) error {
	fmt.Println(" Comienza InsertAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO address (add_UserId, add_Address, add_City, add_State, add_PostalCode, add_Phone, add_Title, add_Name)"

	sentencia += " VALUES ('" + user + "','" + t.AddAddress + "','" + t.AddCity + "','" + t.AddState + "','" + t.AddPostalCode + "','" + t.AddPhone + "','" + t.AddTitle + "','" + t.AddName + "')"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("InsertAddress > EJECUCION EXITOSA!")
	return nil

}
