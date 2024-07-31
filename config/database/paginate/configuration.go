package paginate

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
	"test-backend-developer-sagala/common/constants"
)

type Pagination struct {
	Limit      *int   `json:"limit,omitempty;query:limit"`
	Page       *int   `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	SortValue  string `json:"sort_value,omitempty;query:sort_value"`
	TotalRows  *int64 `json:"total_rows"`
	TotalPages *int   `json:"total_pages"`
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (p *Pagination) GetOffset() *int {
	if p.GetPage() == nil || p.GetLimit() == nil {
		return nil
	} else {
		var i = (*p.GetPage() - 1) * *p.GetLimit()

		return &i
	}
}

func (p *Pagination) GetLimit() *int {
	return p.Limit
}

func (p *Pagination) GetPage() *int {
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = constants.UpdatedAt
	}
	return p.Sort
}

func (p *Pagination) GetSortValue() string {
	if p.SortValue == "" {
		p.SortValue = constants.Desc
	}
	return p.SortValue
}

func (p *Pagination) GetPagination(c *gin.Context) (Pagination, string, string) {
	limit := p.GetLimit()
	page := p.GetPage()
	sort := p.GetSort()
	sortValue := p.GetSortValue()
	var search, filter string

	query := c.Request.URL.Query()
	for key, val := range query {
		queryValue := val[len(val)-1]
		switch key {
		case "limit":
			temp, _ := strconv.Atoi(queryValue)
			limit = &temp
			break
		case "page":
			temp, _ := strconv.Atoi(queryValue)
			page = &temp
			break
		case "sort":
			sort = queryValue
			break
		case "sort_value":
			sortValue = queryValue
			break
		case "search":
			search = queryValue
			break
		case "filter":
			filter = queryValue
		}
	}

	return Pagination{Limit: limit, Page: page, Sort: sort, SortValue: sortValue}, search, filter
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if pagination.GetPage() == nil || pagination.GetLimit() == nil {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(constants.UpdatedAt + " " + constants.Desc)
		}
	} else {
		var totalRows int64
		db.Model(value).Count(&totalRows)

		pagination.TotalRows = &totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(*pagination.Limit)))
		pagination.TotalPages = &totalPages

		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(*pagination.GetOffset()).Limit(*pagination.GetLimit()).Order(pagination.GetSort() + " " + pagination.GetSortValue())
		}
	}
}
