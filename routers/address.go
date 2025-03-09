package routers

import (
	"encoding/json"

	"github.com/juanchi26/CanelonesGaming/bd"
	"github.com/juanchi26/CanelonesGaming/models"
)

func InsertAddress(body string, user string) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Debe especificar la dirección" + err.Error()
	}

	if t.AddName == "" {
		return 400, "Debe especificar el Nombre" + err.Error()
	}

	if t.AddTitle == "" {
		return 400, "Debe especificar el Titulo" + err.Error()
	}

	if t.AddCity == "" {
		return 400, "Debe especificar la Ciudad" + err.Error()
	}

	if t.AddPhone == "" {
		return 400, "Debe especificar el numero de Telefono" + err.Error()
	}

	if t.AddPostalCode == "" {
		return 400, "Debe especificar el codigo postal" + err.Error()
	}

	err = bd.InsertAddress(t, user)

	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la dirección para el ID de usuario" + user + " > " + err.Error()
	}

	return 200, "InsertAddress OK"

}
