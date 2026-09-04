package user

import (
	context "context"
	entity "myAPI/pkg/entity"
	"myAPI/pkg/shared/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceContext struct {
	svc      Service
	mockRepo *MockRepository
}

func (c *serviceContext) BeforeEach(t *testing.T) {
	c.mockRepo = &MockRepository{}
	c.svc = NewService(c.mockRepo, nil, nil)
}

func TestService_Create(t *testing.T) {

	t.Run("Given valid request when create then return user dto",
		func(t *testing.T) {
			c := serviceContext{}
			c.BeforeEach(t)
			req := &entity.UserDto{
				Name:  "Budi",
				Email: "budi@gmail.com",
			}
			expected := &entity.UserDto{
				CommonEntity: common.CommonEntity{
					ID:        "1",
					CreatedAt: time.Now(),
				},
				Name:  "Budi",
				Email: "budi@gmail.com",
			}

			c.mockRepo.On("FindByEmail", mock.Anything, req).Return(nil, entity.ErrUserNotFound).Once()
			c.mockRepo.On("Create", mock.Anything, req).Return(expected, nil).Once()
			res, err := c.svc.Create(context.Background(), req)

			assert.NoError(t, err)
			assert.Equal(t, expected, res)
			c.mockRepo.AssertExpectations(t)

		})
}
