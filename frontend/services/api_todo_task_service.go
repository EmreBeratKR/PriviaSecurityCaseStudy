package services

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

type ApiTodoTaskService struct {
	url string
}

func NewApiTodoTaskService(url string) *ApiTodoTaskService {
	return &ApiTodoTaskService{
		url: url,
	}
}

func (service *ApiTodoTaskService) GetAllNonDeletedByTodoListId(context *fiber.Ctx, todoListId string) *models.TodoTaskGetAllResponseModel {
	url := service.getUrl("/todo-tasks")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoTaskGetAllResponseModel

	client := shared.NewHttpClientGET(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddQueryParam("list_id", todoListId)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoTaskGetAllResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoTaskService) AddWithListIdAndContent(context *fiber.Ctx, todoListId string, content string) *models.TodoTaskGetResponseModel {
	url := service.getUrl("/todo-tasks")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoTaskGetResponseModel

	client := shared.NewHttpClientPOST(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("list_id", todoListId)
	client.AddFormValue("content", content)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoTaskGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoTaskService) ToggleIsCompletedById(context *fiber.Ctx, id string) *models.TodoTaskGetResponseModel {
	url := service.getUrl("/todo-tasks")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoTaskGetResponseModel

	client := shared.NewHttpClientPATCH(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("id", id)
	client.AddFormValue("action", "toggle")
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoTaskGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoTaskService) UpdateContentById(context *fiber.Ctx, id string, content string) *models.TodoTaskGetResponseModel {
	url := service.getUrl("/todo-tasks")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoTaskGetResponseModel

	client := shared.NewHttpClientPATCH(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("id", id)
	client.AddFormValue("action", "edit")
	client.AddFormValue("content", content)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoTaskGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoTaskService) DeleteById(context *fiber.Ctx, id string) *models.TodoTaskGetResponseModel {
	url := service.getUrl("/todo-tasks")
	token := common.GetJWTFromCookies(context)

	var responseBody models.TodoTaskGetResponseModel

	client := shared.NewHttpClientDELETE(url)
	client.SetAuthorizationHeaderBearerToken(token)
	client.AddFormValue("id", id)
	err := client.SendAndParseBody(&responseBody)
	if err != nil {
		return &models.TodoTaskGetResponseModel{
			StatusModel: shared.StatusInternalServerError(),
		}
	}

	response := client.GetResponse()
	responseBody.StatusModel = shared.StatusFromCode(response.StatusCode)

	return &responseBody
}

func (service *ApiTodoTaskService) getUrl(path string) string {
	return service.url + path
}
