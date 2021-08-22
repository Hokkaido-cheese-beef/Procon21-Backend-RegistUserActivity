package dto

type Response struct{
	Message string `json:"Message"`
}

type RegistReq struct{
	UserID string `dynamodbav:"userID" json:"userID"`
	Timestamp int `dynamodbav:"timestamp" json:"timestamp"`
	Status  int  `dynamodbav:"status" json:"status"`
}