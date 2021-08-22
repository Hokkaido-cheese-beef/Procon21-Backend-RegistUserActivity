package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"registUserActivity/pkg/model/dto"
	"strconv"
)

type registUserActivityClientMethods struct {
	Client *dynamodb.DynamoDB
}

func newRegistUserActivityClientMethodsClient(client *dynamodb.DynamoDB) methods {
	return &registUserActivityClientMethods{Client: client}
}

func (r *registUserActivityClientMethods)RegistActivity(req dto.RegistReq)error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("userActivities"),
		Item: map[string]*dynamodb.AttributeValue {
			"userID": {
				S: aws.String(req.UserID),
			},
			"timestamp": {
				N: aws.String(strconv.Itoa( req.Timestamp )),
			},
			"status": {
				N: aws.String(strconv.Itoa( req.Status )),
			},
		},
	}

	_, err := r.Client.PutItem(input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}