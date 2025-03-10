package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanchi26/CanelonesGaming/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza registro order")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := " INSERT INTO orders (order_UserUUID, order_Total, order_AddId) VALUES ('"

	sentencia += o.Order_UserUUID + "'," + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"

	fmt.Println(sentencia)

	var result sql.Result

	result, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	for _, od := range o.OrderDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId)) + "," + strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ")"
		fmt.Println(sentencia)

		_, err = Db.Exec(sentencia)

		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("InsertOrder > Ejecucion exitosa")

	return LastInsertId, nil

}
