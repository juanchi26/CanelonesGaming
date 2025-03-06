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
