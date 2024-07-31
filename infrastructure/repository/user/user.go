package user

import (
	"gorm.io/gorm"
	"test-backend-developer-sagala/common/constants"
	"test-backend-developer-sagala/config/database/paginate"
	"test-backend-developer-sagala/domain"
)

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int64) error
	FindAll(pagination paginate.Pagination, search string) ([]domain.User, paginate.Pagination, error)
	FindById(id int64) (domain.User, error)
	FindAllWithoutPagination() ([]domain.User, error)
	FindByUserName(userName string) (domain.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) Create(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepo) Update(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *userRepo) Delete(id int64) error {
	err := r.db.Where(constants.ById, id).Delete(&domain.User{}).Error

	return err
}

func (r *userRepo) FindAll(pagination paginate.Pagination, search string) ([]domain.User, paginate.Pagination, error) {
	var countries []domain.User

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Where(search).Find(&countries).Error

	return countries, pagination, err
}

func (r *userRepo) FindById(id int64) (domain.User, error) {
	var result domain.User

	err := r.db.Where(constants.ById, id).First(&result).Error

	return result, err
}

func (r *userRepo) FindAllWithoutPagination() ([]domain.User, error) {
	var countries []domain.User

	err := r.db.Find(&countries).Error

	return countries, err
}

func (r *userRepo) FindByUserName(userName string) (domain.User, error) {
	var result domain.User

	err := r.db.Where(constants.ByUserName, userName).First(&result).Error

	return result, err
}
