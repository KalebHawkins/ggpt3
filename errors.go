package ggpt3

type Error struct {
	Message string  `json:"message"`
	Type    string  `json:"type"`
	Param   *string `json:"param"`
	Code    *string `json:"code"`
}

type ErrorResponse struct {
	*Error `json:"error"`
}
