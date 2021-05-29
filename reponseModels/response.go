package reponseModels

type ResponseModel struct {
	Status StatusModel `json:"status"`
	Data   interface{} `json:"data"`
}

type StatusModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
