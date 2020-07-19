package domain_test

import (
	"encoder/domain"
	"testing"
	"github.com/stretchr/testify/require"
	"time"
	uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmpty(t *testing.T){
	video:= domain.NewVideo()
	err:= video.Validate()

	require.Error(t,err)
}

func TestVideoIdIsNotAUuid(t *testing.T){
	video := domain.NewVideo()
	video.ID = "ABC"
	video.ResourceID = "A"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T){
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "A"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}

