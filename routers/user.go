package routers

import (
	"encoding/json"
	"fmt"

	"github.com/juanchi26/CanelonesGaming/bd"
	"github.com/juanchi26/CanelonesGaming/models"
)

func UpdateUser(body string, user string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error al Convertir el JSON"
	}

	if len(t.UserFirstName) == 0 && len(t.UserLastName) == 0 {
		return 400, "Debe especificar el nombre o apellido del Usuario"

	}

	_, encontrado := bd.UserExist(user)

	if !encontrado {
		return 400, "Usuario no encontrado con ese UUID" + user
	}

	err = bd.UpdateUser(t, user)

	if err != nil {
		return 400, "Ocurrio un error al intentar la actualizacion del usuario" + err.Error()
	}

	return 200, "Usuario Actualizado Correctamente"
}

func SelectUser(body string, user string) (int, string) {
	_, encontrado := bd.UserExist(user)

	if !encontrado {
		return 400, "Usuario no encontrado con ese UUID" + user
	}

	row, err := bd.SelectUser(user)

	fmt.Println(row)

	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el select del usuario" + user + " > " + err.Error()
	}

	respJson, err := json.Marshal(row)

	if err != nil {
		return 500, "Error al convertir el JSON"
	}

	return 200, string(respJson)

}
