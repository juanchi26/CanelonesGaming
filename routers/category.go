package routers

import (
	"encoding/json"
	"strconv"

	"github.com/juanchi26/CanelonesGaming/bd"
	"github.com/juanchi26/CanelonesGaming/models"
)

func InsertCategory(body string, user string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Titulo de la categoria" + err.Error()
	}

	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el path" + err.Error()
	}

	isAdmin, msg := bd.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la categoria" + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategId: " + strconv.Itoa(int(result)) + "}"

}

func UpdateCategory(body string, user string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar el CategName y el CategPath" + err.Error()
	}

	isAdmin, msg := bd.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	t.CategID = id

	err2 := bd.UpdateCategory(t)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE de la categoria " + strconv.Itoa(id) + " >" + err2.Error()
	}

	return 200, "Update OK"
}

func DeleteCategory(body string, user string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar el ID de la categoria a borrar"
	}

	isAdmin, msg := bd.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteCategory(id)

	if err != nil {
		return 400, "Error al intentar hacer el DELETE de la categoria" + strconv.Itoa(id) + " > " + err.Error()
	}
	return 200, "DELETE OK"
}
