package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Job struct {
	ID               string    `valid:"uuid"`
	OutputBucketPath string    `valid:"notnull"`
	Status           string    `valid:"notnull"`
	Video            *Video    `valid:"-"`
	VideoID          string    `valid:"-"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `valid:"-"`
	UpdatedAT        time.Time `valid:"-"`
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	j := &Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
		VideoID:          video.ID,
	}

	j.prepare()

	err := j.Validate()
	if err != nil {
		return nil, err
	}

	return j, nil
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (j *Job) prepare() {
	j.ID = uuid.NewV4().String()
	j.CreatedAt = time.Now()
	j.UpdatedAT = time.Now()
}

func (j *Job) Validate() error {
	_, err := govalidator.ValidateStruct(j)
	if err != nil {
		return err
	}
	return nil
}
