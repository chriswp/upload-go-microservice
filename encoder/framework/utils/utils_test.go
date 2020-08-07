package utils_test

import (
	"encoder/framework/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsJson(t *testing.T) {
	json := `{
			  	"id": "cf7dafbd-ba3b-4d7f-8481-5086634d7ade",
			  	"file_path": "teste.mp4",
				"status": "pending"
			  }`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `chr`
	err = utils.IsJson(json)
	require.Error(t, err)
}
