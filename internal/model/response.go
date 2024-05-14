package model

type Response struct {
	Result interface{} `json:"result"`
	Status string      `json:"status"`
}

type Content struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
