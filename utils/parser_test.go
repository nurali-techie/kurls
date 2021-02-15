package utils_test

import (
	"testing"

	"github.com/nurali-techie/kurls/commands"
	"github.com/nurali-techie/kurls/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetCommandList(t *testing.T) {
	args := []string{"list"}

	// test
	cmd := utils.GetCommand(args)

	_, ok := cmd.(*commands.List)
	assert.True(t, ok)
}

func TestGetCommandGet(t *testing.T) {
	args := []string{"CUST_POST"}

	// test
	cmd := utils.GetCommand(args)

	get, ok := cmd.(*commands.Get)
	assert.True(t, ok)
	assert.NotNil(t, get)
	assert.Equal(t, "CUST_POST", get.Key)
}

func TestGetCommandAdd(t *testing.T) {
	wantCurl := `curl localhost:8080/api/customer -H 'Content-Type: application/json' -d {"name": "ali", "city": "bangalore"}`
	args := []string{"CUST_POST", wantCurl}

	// test
	cmd := utils.GetCommand(args)

	add, ok := cmd.(*commands.Add)
	assert.True(t, ok)
	assert.NotNil(t, add)
	assert.Equal(t, "CUST_POST", add.Key)
	assert.Equal(t, wantCurl, add.Curl)
}

func TestCommandNil(t *testing.T) {
	wrongCurl := `localhost:8080/api/customer -H 'Content-Type: application/json' -d {"name": "ali", "city": "bangalore"}`
	args := []string{"CUST_POST", wrongCurl}

	// test
	cmd := utils.GetCommand(args)

	assert.Nil(t, cmd)
}
