package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"registUserActivity/pkg/model/dao"
	"registUserActivity/pkg/model/dto"
	"registUserActivity/pkg/res"
)



func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
	var response dto.Response
	var req dto.RegistReq
	err :=json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		log.Println(err)
		return res.ReturnInternalServerErrorResponse(err)
	}

	if req.UserID=="" || req.Status==0 || req.Timestamp==0{
		response.Message = "request is wrong"
		responseBody, _ := json.Marshal(response)
		return res.ReturnBadRequestResponse(string(responseBody)), nil
	}

	client, err := dao.New()
	if err != nil {
		log.Println(err)
		return res.ReturnInternalServerErrorResponse(err)
	}

	err = client.RegistUserActivity.RegistUserActivityLogic.RegistActivity(req)
	if err != nil {
		return res.ReturnInternalServerErrorResponse(err)
	}

	responseBody, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(responseBody),
		StatusCode: 200,
	},nil
}

func main(){
	lambda.Start(handler)
}
