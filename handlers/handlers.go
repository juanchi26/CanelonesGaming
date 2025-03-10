package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juanchi26/CanelonesGaming/auth"
	"github.com/juanchi26/CanelonesGaming/routers"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy a procesar " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOK, statusCode, user := ValidoAuthorization(path, method, headers)

	fmt.Println("Usuario validado:", user)
	fmt.Println("Headers recibidos:", headers)

	if !isOK {
		return statusCode, user
	}

	switch path[1:5] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)
	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)
	case "addr":
		return ProcesoAddress(body, path, method, user, idn, request)
	case "cate":
		fmt.Println("Entrando a ProcesoCategory con método:", method)
		return ProcesoCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)

	default:
		fmt.Println("Path no reconocido:", path[1:5])
	}

	return 400, "Method Invalido"
}

func ValidoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {

	if (path == "/product" && method == "GET") || (path == "/category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]

	if len(token) == 0 {
		return false, 401, "Token Requerido"
	}

	todoOK, err, msg := auth.ValidoToken(token)

	if !todoOK {
		if err != nil {
			fmt.Println("Error en el Token" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")

	return true, 200, msg

}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {

	if path == "/user/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)
		}
	}

	if path == "/users" {
		if method == "GET" {
			return routers.SelectUsers(body, user, request)
		}
	}

	return 400, "Method invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "PUT":
		return routers.UpdateProducts(body, user, id)
	case "DELETE":
		return routers.DeleteProducts(user, id)
	case "GET":
		return routers.SelectProducts(request)

	}

	return 400, "Method invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(body, user, id)
	case "GET":
		return routers.SelectCategories(body, request)
	}

	return 400, "Method invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return routers.UpdateStock(body, user, id)

}

func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertAddress(body, user)
	case "PUT":
		return routers.UpdateAddress(body, user, id)
	case "DELETE":
		return routers.DeleteAddress(user, id)
	case "GET":
		return routers.SelectAddress(user)
	}

	return 400, "Method invalid"
}

func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method invalid"
}
