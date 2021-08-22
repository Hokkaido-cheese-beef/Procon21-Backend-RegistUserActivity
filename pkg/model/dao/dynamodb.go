package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"registUserActivity/pkg/model/dto"
)

type  DynamoDB struct{
	Dynamo  *dynamodb.DynamoDB
	RegistUserActivity Methods
}

type Methods struct {
	RegistUserActivityLogic methods
}

type methods interface {
	RegistActivity(req dto.RegistReq)error
}

func New()(*DynamoDB,error){
	//DB接続
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// init methods
	registUserActivityMethod := newRegistUserActivityClientMethodsClient(svc)

	return &DynamoDB{
		Dynamo:  svc,
		RegistUserActivity: Methods{registUserActivityMethod},
	},nil
}
