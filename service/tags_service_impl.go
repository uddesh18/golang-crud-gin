package service

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagRequest) {
	err := t.Validate.Struct(tags)
	helper.Errorpanic(err)

	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse

	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagsdata, err := t.TagsRepository.FindById(tagsId)
	helper.Errorpanic(err)

	tagResp := response.TagsResponse{
		Id:   tagsdata.Id,
		Name: tagsdata.Name,
	}
	return tagResp
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdateTagRequest) {
	tagsData, err := t.TagsRepository.Update(tags.Id)
	helper.Errorpanic(err)

	tagsData.Name = tags.Name
	t.TagsRepository.Update(tagsData.Name)
}
