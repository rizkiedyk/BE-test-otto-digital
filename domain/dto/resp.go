package dto

type Resp struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Success     bool        `json:"success" bson:"success"`
	Code        int         `json:"code" bson:"code"`
	Message     string      `json:"message" bson:"message"`
	ErrorDetail interface{} `json:"error_detail" bson:"error_detail"`
}
