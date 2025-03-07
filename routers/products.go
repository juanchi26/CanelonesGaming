package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
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

func SelectProducts(request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.Product
	var page, pageSize int
	var orderType, orderField string

	param := request.QueryStringParameters

	page, _ = strconv.Atoi(param["page"])

	pageSize, _ = strconv.Atoi(param["pageSize"])

	orderType = param["orderType"] // D = Descendente, A o Nil = Ascendente

	orderField = param["orderField"] // I id , T TITULO, P PRECIO, D DESCRIPCION, F CREATED AT, C CATEGID

	if !strings.Contains("ITDFPCS", orderField) {
		orderField = ""
	}

	var choice string

	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}

	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch = param["search"]
	}

	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}

	if len(param["slug"]) > 0 {
		choice = "U"
		t.ProdPath = param["slug"]
	}

	if len(param["slugCateg"]) > 0 {
		choice = "K"
		t.ProdCategPath = param["slugCateg"]
	}

	fmt.Println(param)

	result, err2 := bd.SelectProducts(t, choice, page, pageSize, orderType, orderField)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el SELECT de Productos >" + choice + err2.Error()
	}

	Product, err3 := json.Marshal(result)

	if err3 != nil {
		return 400, "Ocurrio un error al intentar Convertir en JSON la busqueda del producto" + err3.Error()
	}

	return 200, string(Product)
}
