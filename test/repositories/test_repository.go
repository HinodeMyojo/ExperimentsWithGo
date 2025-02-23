package repositories

import (
	"test/models"

	"github.com/gin-gonic/gin"
)

type TestRepository interface{
	GetTests(c *gin.Context) ([]models.Test, error)
}

type TestRepositoryImpl struct{}

func (t *TestRepositoryImpl) GetTests(c *gin.Context) ([]models.Test, error) {
	tests := []models.Test{
		{
			Id:   1,
			Name: "Test 1",
		},
		{
			Id:   2,
			Name: "Test 2",
		}};
	return tests, nil
}


func (r TestRepositoryImpl) GetTest(c Test)
