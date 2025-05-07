package services

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

type ApiTodoListService struct {
	url string
}

func NewApiTodoListService(url string) *ApiTodoListService {
	return &ApiTodoListService{
		url: url,
	}
}

func (service *ApiTodoListService) GetNonDeletedById(context *fiber.Ctx, id string) *models.TodoListGetResponseModel {
	url := service.getUrl("/todo-lists")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoListGetResponseModel

	client := shared.NewHttpClientGET(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddQueryParam("id", id)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoListGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoListService) GetAllNonDeleted(context *fiber.Ctx) *models.TodoListGetAllResponseModel {
	url := service.getUrl("/todo-lists")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoListGetAllResponseModel

	client := shared.NewHttpClientGET(url)
	client.SetAuthorizationHeaderBearerToken(token)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoListGetAllResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoListService) GetAllNonDeletedByUserId(context *fiber.Ctx, userId string) *models.TodoListGetAllResponseModel {
	url := service.getUrl("/todo-lists")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoListGetAllResponseModel

	client := shared.NewHttpClientGET(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddQueryParam("user_id", userId)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoListGetAllResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoListService) AddWithUserIdAndName(context *fiber.Ctx, userId string, name string) *models.TodoListGetResponseModel {
	url := service.getUrl("/todo-lists")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoListGetResponseModel

	client := shared.NewHttpClientPOST(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("user_id", userId)
	client.AddFormValue("name", name)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoListGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoListService) UpdateNameById(context *fiber.Ctx, id string, name string) *models.TodoListGetResponseModel {
	url := service.getUrl("/todo-lists")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoListGetResponseModel

	client := shared.NewHttpClientPATCH(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("id", id)
	client.AddFormValue("name", name)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoListGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoListService) DeleteById(context *fiber.Ctx, id string) *models.TodoListGetResponseModel {
	url := service.getUrl("/todo-lists")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoListGetResponseModel

	client := shared.NewHttpClientDELETE(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("id", id)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoListGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoListService) getUrl(path string) string {
	return service.url + path
}
