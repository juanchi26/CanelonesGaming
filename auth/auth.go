package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("El Token no es Valido!")
		return false, nil, "El token no es Valido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		fmt.Println("No se pudo decodificar la parte del token :", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON

	err = json.Unmarshal(userInfo, &tkj)

	if err != nil {
		fmt.Println("No se puede decodificar en la estrutcura Json", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Token Expirado, Fecha de expiracion del Token" + tm.String())

		return false, err, "el Token a Expirado!!"
	}

	return true, nil, string(tkj.Username)

}
