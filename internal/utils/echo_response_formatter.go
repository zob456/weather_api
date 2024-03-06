package utils

type EchoMap struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func EchoResponse(message string, data any) EchoMap {
	return EchoMap{
		Message: message,
		Data:    data,
	}
}
