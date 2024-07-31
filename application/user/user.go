package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"net/http"
	"test-backend-developer-sagala/common/constants"
	"test-backend-developer-sagala/common/password"
	responseMessage "test-backend-developer-sagala/common/response-message"
	jwtToken "test-backend-developer-sagala/config/jwt-token"
	authenticationDto "test-backend-developer-sagala/contract/authentication-dto"
	"test-backend-developer-sagala/domain"
	"test-backend-developer-sagala/infrastructure/repository/user"
)

type UserService interface {
	Register(dto authenticationDto.Register) responseMessage.Response
	Login(dto authenticationDto.Login) responseMessage.Response
}

type userService struct {
	userRepository user.UserRepository
}

func NewUserService(
	userRepository user.UserRepository,
) *userService {
	return &userService{
		userRepository,
	}
}

func (s *userService) Register(dto authenticationDto.Register) responseMessage.Response {
	response := s.validateRegisterUser(dto)
	if (response != responseMessage.Response{}) {
		return response
	}

	hashedPassword, err := password.HashPassword(dto.Password)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	data := dtoToModelsCreate(dto, hashedPassword)
	result, err := s.userRepository.Create(data)
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

func (s *userService) validateRegisterUser(dto authenticationDto.Register) responseMessage.Response {
	findByUserName, _ := s.userRepository.FindByUserName(dto.Username)
	if findByUserName.ID > 0 {
		return responseMessage.Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    fmt.Errorf("user name %s already taken", dto.Username).Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{}
}

func dtoToModelsCreate(dto authenticationDto.Register, hashedPassword string) domain.User {
	return domain.User{
		Username: dto.Username,
		Password: hashedPassword,
	}
}

func (s *userService) Login(dto authenticationDto.Login) responseMessage.Response {
	userData, err := s.userRepository.FindByUserName(dto.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusUnauthorized,
				Success:    false,
				Message:    fmt.Errorf("invalid username or password").Error(),
				Data:       nil,
			}

		}
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	compare := password.ComparePassword(userData.Password, dto.Password)
	if !compare {
		return responseMessage.Response{
			StatusCode: http.StatusUnauthorized,
			Success:    false,
			Message:    "invalid username or password",
			Data:       nil,
		}
	}

	token, expiredAt, err := jwtToken.GenerateToken(jwtToken.JwtClaim{
		ID:             userData.ID,
		Username:       userData.Username,
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	result := authenticationDto.LoginResult{
		Type:        constants.Bearer,
		AccessToken: token,
		ExpiredAt:   expiredAt,
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}
}
