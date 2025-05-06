package services

import (
	"encoding/json"
	"io"
	"net/http"
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"
)

type ApiUserService struct {
	Url string
}

func NewApiUserService(url string) *ApiUserService {
	return &ApiUserService{
		Url: url,
	}
}

func (service *ApiUserService) Login(request *models.LoginRequestModel) *models.LoginResponseModel {
	url := service.getUrl("/users/login")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &models.LoginResponseModel{
			StatusModel: shared.StatusInternalServerError(),
			Message:     "Something gone wrong, try again",
		}
	}

	req.SetBasicAuth(request.Username, request.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &models.LoginResponseModel{
			StatusModel: shared.StatusInternalServerError(),
			Message:     "Something gone wrong, try again",
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &models.LoginResponseModel{
			StatusModel: shared.StatusInternalServerError(),
			Message:     "Something gone wrong, try again",
		}
	}

	var responseBody models.LoginResponseModel
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return &models.LoginResponseModel{
			StatusModel: shared.StatusInternalServerError(),
			Message:     "Something gone wrong, try again",
		}
	}

	responseBody.StatusModel = shared.StatusFromCode(resp.StatusCode)

	return &responseBody
}

func (service *ApiUserService) getUrl(path string) string {
	return service.Url + path
}
