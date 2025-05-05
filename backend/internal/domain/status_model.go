package domain

type StatusModel struct {
	Status string `json:"status"`
}

func StatusSuccess() StatusModel {
	return StatusModel{Status: "success"}
}

func StatusBadRequest() StatusModel {
	return StatusModel{Status: "bad_request"}
}

func StatusForbidden() StatusModel {
	return StatusModel{Status: "forbidden"}
}

func StatusNotFound() StatusModel {
	return StatusModel{Status: "not_found"}
}

func StatusInternalServerError() StatusModel {
	return StatusModel{Status: "internal_server_error"}
}

func (model *StatusModel) IsNotSuccess() bool {
	return model.Status != StatusSuccess().Status
}
