package routers

import (
	"encoding/json"
	"strconv"

	"github.com/juanchi26/CanelonesGaming/bd"
	"github.com/juanchi26/CanelonesGaming/models"
)

func InsertProduct(body string, user string) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especifiar el nombre (Title) del producto"
	}

	isAdmin, msg := bd.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar InsertProduct > " + err2.Error()
	}

	return 200, " { ProductID: " + strconv.Itoa(int(result)) + "}"

}

func UpdateProducts(body string, user string, id int) (int, string) {
	var p models.Product

	err := json.Unmarshal([]byte(body), &p)

	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	isAdmin, msg := bd.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	p.ProdId = id

	err2 := bd.UpdateProducts(p)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del Producto " + strconv.Itoa(id) + " >" + err2.Error()
	}

	return 200, "Update OK"
}

func DeleteProducts(user string, id int) (int, string) {
	isAdmin, msg := bd.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	err2 := bd.DeleteProducts(id)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el DELETE del Producto " + strconv.Itoa(id) + " >" + err2.Error()
	}

	return 200, "DELETE OK"

}
