package common

type ValueResponseModel struct {
	Message string `json:"message"`
	Value   any    `json:"value"`
}

func ValueResponseOk(value any) ValueResponseModel {
	return ValueResponseModel{
		Message: "200 - OK",
		Value:   value,
	}
}
