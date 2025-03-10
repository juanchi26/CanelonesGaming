package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

	sentencia := "INSERT INTO addresses (add_UserId, add_Address, add_City, add_State, add_PostalCode, add_Phone, add_Title, add_Name)"

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

func AddressExist(user string, id int) (error, bool) {
	fmt.Println("Comienza AddressExist")

	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM addresses WHERE add_Id = " + strconv.Itoa(id) + " AND add_UserId = '" + user + "'"

	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err, false
	}

	defer rows.Close()

	var valor string

	rows.Next()

	rows.Scan(&valor)

	fmt.Println("AddressExist > Ejecucion Exitosa - Valor:", valor)

	if valor == "1" {
		return nil, true
	}

	return nil, false

}

func UpdateAddress(addr models.Address) error {
	fmt.Println("Comienza UpdateAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE addresses SET "

	if addr.AddAddress != "" {
		sentencia += "Add_Address = '" + addr.AddAddress + "', "
	}

	if addr.AddCity != "" {
		sentencia += "Add_City = '" + addr.AddCity + "', "
	}

	if addr.AddName != "" {
		sentencia += "Add_Name = '" + addr.AddName + "', "
	}

	if addr.AddPhone != "" {
		sentencia += "Add_Phone = '" + addr.AddPhone + "', "
	}

	if addr.AddPostalCode != "" {
		sentencia += "Add_PostalCode = '" + addr.AddPostalCode + "', "
	}

	if addr.AddState != "" {
		sentencia += "Add_State = '" + addr.AddState + "', "
	}

	if addr.AddTitle != "" {
		sentencia += "Add_Title = '" + addr.AddTitle + "', "
	}

	sentencia, _ = strings.CutSuffix(sentencia, ", ")

	sentencia += " WHERE Add_Id = " + strconv.Itoa(addr.AddId)

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("UpdateAddress > EJECUCION EXITOSA!")
	return nil

}

func DeleteAddress(id int) error {
	fmt.Println("Comienza DeleteAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("DeleteAddress > EJECUCION EXITOSA!")
	return nil

}

func SelectAddress(user string) ([]models.Address, error) {
	fmt.Println("Comienza SelectAddress")

	Addr := []models.Address{}

	err := DbConnect()
	if err != nil {
		return Addr, err
	}
	defer Db.Close()

	sentencia := "SELECT add_Id, add_Address, add_City, add_State, add_PostalCode, add_Phone, add_Title, add_Name FROM addresses WHERE add_UserId = '" + user + "'"

	var rows *sql.Rows

	rows, err = Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return Addr, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress, addCity, addState, addPostalCode, addPhone, addTitle, addName sql.NullString

		err = rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)

		if err != nil {
			fmt.Println(err.Error())
			return Addr, err
		}

		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String

		Addr = append(Addr, a)

	}

	fmt.Println("SelectAddress > Ejecucion Exitosa")

	return Addr, nil

}
