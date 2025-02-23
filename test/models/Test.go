package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Tests = []Test{
	{Id: 1, Name: "Test 1"},
	{Id: 2, Name: "Test 2"},
}
