package services

import (
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

type ApiUserService struct {
	url string
}

func NewApiUserService(url string) *ApiUserService {
	return &ApiUserService{
		url: url,
	}
}

func (service *ApiUserService) Login(context *fiber.Ctx, request *models.LoginRequestModel) *models.LoginResponseModel {
	url := service.getUrl("/users/login")

	var responseBody models.LoginResponseModel
	client := shared.NewHttpClientGET(url)
	client.SetAuthorizationHeaderBasicAuth(request.Username, request.Password)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.LoginResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiUserService) getUrl(path string) string {
	return service.url + path
}
