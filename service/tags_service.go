package service

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagRequest)
	Update(tags request.UpdateTagRequest)
	Delete(tagsId int) response.TagsResponse
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
