package dto

type DeviceInfo struct{
	DeviceID string `dynamodbav:"deviceID"`
}

type Response struct{
	Message string `json:"Message"`
}

type RegistReq struct{
	UserID string `dynamodbav:"userID" json:"userID"`
	Timestamp int `dynamodbav:"timestamp" json:"timestamp"`
	status  int  `dynamodbav:"status" json:"status"`
}