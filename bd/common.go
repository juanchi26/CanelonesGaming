package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanchi26/CanelonesGaming/models"
	"github.com/juanchi26/CanelonesGaming/secretm"
)

var SecretModel models.SecretRDSjson
var err error

var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("Secretname"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa a la base de datos")
	return nil
}

func ConnStr(claves models.SecretRDSjson) string {
	var dbUser, authToken, dbEndPoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndPoint = claves.Host
	dbName = "CanelonesGaming"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	fmt.Println(dsn) //borrar en produccion
	return dsn
}

func IsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza IsAdmin")
	err := DbConnect()

	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"

	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)

	if err != nil {
		return false, err.Error()
	}

	var valor string

	rows.Next()
	rows.Scan(&valor)

	fmt.Println("IsAdmin > Ejecucion exitosa - valor devuelto" + valor)

	if valor == "1" {
		return true, ""
	}

	return false, "El Usuario NO es Administrador"

}
