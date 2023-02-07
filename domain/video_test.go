package domain_test

import (
	"testing"
	"time"

	"github.com/diegolopes98/video-encoder-go/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "not uuid"
	video.ResourceID = "random id"
	video.FilePath = "random path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIsOk(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "random id"
	video.FilePath = "random path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
