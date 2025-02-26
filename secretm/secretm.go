package secretm

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/juanchi26/CanelonesGaming/awsgo"
	"github.com/juanchi26/CanelonesGaming/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSjson, error) { //funcion de SECRET MANAGER
	var datosSecret models.SecretRDSjson
	fmt.Println("> Pudio secreto" + nombreSecret)

	nombreSecret = strings.TrimSpace(nombreSecret) // Elimina espacios en blanco
	if nombreSecret == "" {
		return datosSecret, errors.New("el nombre del secreto está vacío")
	}

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error()) //si hay algun error retorna datos secret sin cambios y el error
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret) //estructura para procesar lo que me devuelta secret value, primero es lo que recibe y lo segundo donde lo graba
	fmt.Sprintln(" > lectura secret OK" + nombreSecret)
	return datosSecret, nil
}
