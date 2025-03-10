package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juanchi26/CanelonesGaming/bd"
	"github.com/juanchi26/CanelonesGaming/models"
)

func InsertOrder(body string, user string) (int, string) {
	var o models.Orders

	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		fmt.Println("Error al unmarshal body:", err.Error())
		return 400, "Error en los datos recibidos"
	}

	o.Order_UserUUID = user

	OK, msg := ValidOrder(o)

	if !OK {
		return 400, msg
	}

	result, err2 := bd.InsertOrder(o)
	if err2 != nil {
		fmt.Println("Error al insertar Orden:", err2.Error())
		return 400, "Error al insertar Orden"
	}

	return 200, "{OrderID : " + strconv.Itoa(int(result)) + "}"

}

func ValidOrder(o models.Orders) (bool, string) {

	if o.Order_Total == 0 {
		return false, "Debe indica el total de la orden"
	}

	count := 0

	for _, od := range o.OrderDetails {
		if od.OD_ProdId == 0 {
			return false, "Debe indicar el ID del producto en el detalle de la Orden"
		}

		if od.OD_Quantity == 0 {
			return false, "Debe indicar la cantidad del producto en el detalle de la Orden"
		}
		count++

	}
	if count == 0 {
		return false, "Debe indicar al menos un producto en la orden"
	}

	return true, ""
}

func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {

	var fechaDesde, fechaHasta string
	var orderId int
	var page int

	if len(request.QueryStringParameters["fechaDesde"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechaDesde"]
	}

	if len(request.QueryStringParameters["fechaHasta"]) > 0 {
		fechaHasta = request.QueryStringParameters["fechaHasta"]
	}

	if len(request.QueryStringParameters["page"]) > 0 {
		page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	if len(request.QueryStringParameters["orderId"]) > 0 {
		orderId, _ = strconv.Atoi(request.QueryStringParameters["orderId"])
	}

	result, err := bd.SelectOrders(user, fechaDesde, fechaHasta, page, orderId)

	if err != nil {
		return 400, "Ocurrio un error al intentar capturar los registro de ordenes del" + fechaDesde + "al" + fechaHasta + " > " + err.Error()
	}

	Orders, err2 := json.Marshal(result)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar convertir el JSON el registro de orden"
	}

	return 200, string(Orders)
}
