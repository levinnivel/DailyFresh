package response

type LoginResponse struct {
	Status  int         `form:"status" json:"status"`
	Message string      `form:"message" json:"message"`
	Data    interface{} `form:"data" json:"data"`
	Token   string      `form:"token" json:"token"`
}
