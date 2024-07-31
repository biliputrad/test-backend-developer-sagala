package article

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"test-backend-developer-sagala/common/constants"
	responseMessage "test-backend-developer-sagala/common/response-message"
	"test-backend-developer-sagala/config/database/paginate"
	articleDto "test-backend-developer-sagala/contract/article-dto"
	"test-backend-developer-sagala/domain"
	repository "test-backend-developer-sagala/infrastructure/repository/article"
)

type ArticleService interface {
	Create(dto articleDto.Create) responseMessage.Response
	FindById(id int64) responseMessage.Response
	FindAll(pagination paginate.Pagination, search string, filter string) responseMessage.ResponsePaginate
	FindAllWithoutPagination() responseMessage.Response
	Delete(id int64) responseMessage.Response
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(articleRepository repository.ArticleRepository) *articleService {
	return &articleService{articleRepository}
}

func (s *articleService) Create(dto articleDto.Create) responseMessage.Response {
	data := dtoToModelsCreate(dto)
	result, err := s.articleRepository.Create(data)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    constants.ResponseCreated,
		Data:       result,
	}
}

func dtoToModelsCreate(dto articleDto.Create) domain.Article {
	return domain.Article{
		Author: dto.Author,
		Title:  dto.Title,
		Body:   dto.Body,
	}
}

func (s *articleService) FindById(id int64) responseMessage.Response {
	result, err := s.articleRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("article with id %d not found", id).Error(),
				Data:       nil,
			}

		} else {
			return responseMessage.Response{
				StatusCode: http.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
				Data:       nil,
			}
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}
}

func (s *articleService) FindAll(pagination paginate.Pagination, search string, filter string) responseMessage.ResponsePaginate {
	var result responseMessage.ResponsePaginate
	if search != "" {
		result = s.findAllWithFilterAndSearch(pagination, search, filter)
	} else if filter != "" {
		result = s.findAllWithFilter(pagination, filter)

	} else {
		result = s.findAllWithoutQuery(pagination)
	}

	return result
}

func (s *articleService) findAllWithFilterAndSearch(pagination paginate.Pagination, search string, filter string) responseMessage.ResponsePaginate {
	query := " (author ILIKE '%" + search + "%' OR title ILIKE '%" + search + "%' OR body ILIKE '%" + search + "%')"
	if filter != "" {
		query += " AND author = '" + filter + "'"
	}

	result, paginateResult, err := s.articleRepository.FindAllWithQuery(pagination, query)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}

func (s *articleService) findAllWithFilter(pagination paginate.Pagination, filter string) responseMessage.ResponsePaginate {
	query := "author = '" + filter + "'"

	result, paginateResult, err := s.articleRepository.FindAllWithQuery(pagination, query)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}

func (s *articleService) findAllWithoutQuery(pagination paginate.Pagination) responseMessage.ResponsePaginate {
	result, paginateResult, err := s.articleRepository.FindAllWithoutQuery(pagination)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}

func (s *articleService) FindAllWithoutPagination() responseMessage.Response {
	result, err := s.articleRepository.FindAllWithoutPagination()
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}
}

func (s *articleService) Delete(id int64) responseMessage.Response {
	_, err := s.articleRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("article with id %d not found", id).Error(),
				Data:       nil,
			}

		} else {
			return responseMessage.Response{
				StatusCode: http.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
				Data:       nil,
			}
		}
	}

	err = s.articleRepository.Delete(id)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       true,
	}
}
