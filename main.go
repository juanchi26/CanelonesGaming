package main

import (
	"os"
	"strings"

	"github.com/juanchi26/CanelonesGaming/awsgo"
	"github.com/juanchi26/CanelonesGaming/bd"
	"github.com/juanchi26/CanelonesGaming/handlers"

	"context"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAWS()
	if !ValidoParametros() {
		panic("Error en los Parametros debe enviar 'Secretname','UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse

	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	//

	status, message := handlers.Manejadores(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil
}

func ValidoParametros() bool {

	_, traeParametro := os.LookupEnv("Secretname")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro

}
