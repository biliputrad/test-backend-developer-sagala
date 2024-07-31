package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"test-backend-developer-sagala/application/article"
	responseMessage "test-backend-developer-sagala/common/response-message"
	"test-backend-developer-sagala/config/database/paginate"
	articleDto "test-backend-developer-sagala/contract/article-dto"
)

type articleController struct {
	articleService article.ArticleService
	pagination     paginate.Pagination
}

func NewArticleController(articleService article.ArticleService, pagination paginate.Pagination) *articleController {
	return &articleController{articleService, pagination}
}

func (h *articleController) Create(c *gin.Context) {
	var input articleDto.Create
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	input.CreatedBy = c.GetInt64("id")
	result := h.articleService.Create(input)
	c.JSON(result.StatusCode, result)
}

func (h *articleController) FindById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.articleService.FindById(id)
	c.JSON(result.StatusCode, result)
}

func (h *articleController) FindAll(c *gin.Context) {
	pagination, search, filter := h.pagination.GetPagination(c)
	result := h.articleService.FindAll(pagination, search, filter)

	c.JSON(result.StatusCode, result)
}

func (h *articleController) FindAllWithoutPagination(c *gin.Context) {
	result := h.articleService.FindAllWithoutPagination()
	c.JSON(result.StatusCode, result)
}

func (h *articleController) Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.articleService.Delete(id)
	c.JSON(result.StatusCode, result)
}
