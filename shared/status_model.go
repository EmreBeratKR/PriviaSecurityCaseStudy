package shared

type StatusModel struct {
	Status string `json:"status"`
}

func StatusFromCode(code int) StatusModel {
	if code == 200 {
		return StatusSuccess()
	}

	if code == 400 {
		return StatusBadRequest()
	}

	if code == 401 {
		return StatusUnauthorized()
	}

	if code == 403 {
		return StatusForbidden()
	}

	if code == 404 {
		return StatusNotFound()
	}

	return StatusInternalServerError()
}

func StatusSuccess() StatusModel {
	return StatusModel{Status: "success"}
}

func StatusBadRequest() StatusModel {
	return StatusModel{Status: "bad_request"}
}

func StatusUnauthorized() StatusModel {
	return StatusModel{Status: "unauthorized"}
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
