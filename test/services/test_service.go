package services

import "github.com/gin-gonic/gin"
import "test/repositories"
import "test/models"


type TestService interface {
	GetTests(c *gin.Context) (any, error)
}

type TestServiceImpl struct{
	repo repositories.TestRepository 
}

func (t TestServiceImpl) GetTests(c *gin.Context) (any, error) {
	result, err := t.repo.GetTests(c)
	if err != nil {
		return nil, err
	}
	return result, nil
}
