package user

import (
	"context"
	"myAPI/database"
	"myAPI/pkg/entity"
	"myAPI/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	FindAll(ctx context.Context) ([]entity.UserDto, error)
	FindById(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	FindByEmail(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	Update(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	Delete(ctx context.Context, req *entity.UserDto) error
	FindByIdForUpdate(ctx context.Context, id string) (*entity.UserDto, error)
	UpdateBalance(ctx context.Context, id string, newBalance int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {

	if req.ID == "" {
		req.ID = uuid.New().String()
	}

	m := entity.NewUserModelFromDto(req)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return nil, err
	}

	return entity.NewUserDtoFromModel(m), nil
}

func (r *repository) FindAll(ctx context.Context) ([]entity.UserDto, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	var userDto []entity.UserDto
	for _, user := range users {
		userDto = append(userDto, *entity.NewUserDtoFromModel(&user))
	}

	return userDto, nil

}

func (r *repository) FindById(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", req.ID).First(&user).Error; err != nil {
		return nil, entity.ErrUserNotFound
	}

	return entity.NewUserDtoFromModel(&user), nil
}

func (r *repository) Update(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	m := entity.NewUserModelFromDto(req)

	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", req.ID).Updates(m).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *repository) Delete(ctx context.Context, req *entity.UserDto) error {

	if err := r.db.WithContext(ctx).Where("id = ?", req.ID).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByEmail(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, entity.ErrUserNotFound
	}

	return entity.NewUserDtoFromModel(&user), nil
}

func (r *repository) getDB(ctx context.Context) *gorm.DB {
	return database.ExtractTx(ctx, r.db)
}

func (r *repository) FindByIdForUpdate(ctx context.Context, id string) (*entity.UserDto, error) {

	var m model.User
	err := r.getDB(ctx).WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", id).
		First(&m).Error

	if err != nil {
		return nil, err
	}
	return entity.NewUserDtoFromModel(&m), nil
}

func (r *repository) UpdateBalance(ctx context.Context, id string, newBalance int) error {
	return r.getDB(ctx).WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Update("balance", newBalance).Error
}
