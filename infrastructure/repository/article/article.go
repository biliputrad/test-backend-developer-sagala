package article

import (
	"gorm.io/gorm"
	"test-backend-developer-sagala/common/constants"
	"test-backend-developer-sagala/config/database/paginate"
	"test-backend-developer-sagala/domain"
)

type ArticleRepository interface {
	Create(article domain.Article) (domain.Article, error)
	Update(article domain.Article) (domain.Article, error)
	Delete(id int64) error
	FindAllWithQuery(pagination paginate.Pagination, query string) ([]domain.Article, paginate.Pagination, error)
	FindById(id int64) (domain.Article, error)
	FindAllWithoutPagination() ([]domain.Article, error)
	FindAllWithoutQuery(pagination paginate.Pagination) ([]domain.Article, paginate.Pagination, error)
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepo {
	return &articleRepo{db}
}

func (r *articleRepo) Create(article domain.Article) (domain.Article, error) {
	err := r.db.Create(&article).Error

	return article, err
}

func (r *articleRepo) Update(article domain.Article) (domain.Article, error) {
	err := r.db.Save(&article).Error

	return article, err
}

func (r *articleRepo) Delete(id int64) error {
	err := r.db.Where(constants.ById, id).Delete(&domain.Article{}).Error

	return err
}

func (r *articleRepo) FindAllWithQuery(pagination paginate.Pagination, query string) ([]domain.Article, paginate.Pagination, error) {
	var countries []domain.Article

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Where(query).Find(&countries).Error

	return countries, pagination, err
}

func (r *articleRepo) FindById(id int64) (domain.Article, error) {
	var result domain.Article

	err := r.db.Where(constants.ById, id).First(&result).Error

	return result, err
}

func (r *articleRepo) FindAllWithoutPagination() ([]domain.Article, error) {
	var countries []domain.Article

	err := r.db.Find(&countries).Error

	return countries, err
}

func (r *articleRepo) FindAllWithoutQuery(pagination paginate.Pagination) ([]domain.Article, paginate.Pagination, error) {
	var countries []domain.Article

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Find(&countries).Error

	return countries, pagination, err
}
