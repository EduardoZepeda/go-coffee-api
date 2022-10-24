package repository

import (
	"context"

	"github.com/EduardoZepeda/go-coffee-api/models"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetCafes(ctx context.Context, page uint64, size uint64) ([]*models.Shop, error)
	GetCafeById(ctx context.Context, id string) (*models.Shop, error)
	CreateCafe(ctx context.Context, shopRequest *models.CreateShop) error
	DeleteCafe(ctx context.Context, id string) error
	UpdateCafe(ctx context.Context, shopRequest *models.InsertShop) error
	SearchCafe(ctx context.Context, query string, page uint64, size uint64) ([]*models.Shop, error)
	GetNearestCafes(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.Shop, error)
	GetUser(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.GetUserResponse, error)
	RegisterUser(ctx context.Context, user *models.SignUpRequest) error
	UpdateUser(ctx context.Context, user *models.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetCafes(ctx context.Context, page uint64, size uint64) ([]*models.Shop, error) {
	return implementation.GetCafes(ctx, page, size)
}

func GetCafeById(ctx context.Context, id string) (*models.Shop, error) {
	return implementation.GetCafeById(ctx, id)
}

func CreateCafe(ctx context.Context, shopRequest *models.CreateShop) error {
	return implementation.CreateCafe(ctx, shopRequest)
}

func DeleteCafe(ctx context.Context, id string) error {
	return implementation.DeleteCafe(ctx, id)
}

func UpdateCafe(ctx context.Context, shopRequest *models.InsertShop) error {
	return implementation.UpdateCafe(ctx, shopRequest)
}

func SearchCafe(ctx context.Context, query string, page uint64, size uint64) ([]*models.Shop, error) {
	return implementation.SearchCafe(ctx, query, page, size)
}

func GetNearestCafes(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.Shop, error) {
	return implementation.GetNearestCafes(ctx, UserCoordinates)
}

func GetUser(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUser(ctx, email)
}

func GetUserById(ctx context.Context, id string) (*models.GetUserResponse, error) {
	return implementation.GetUserById(ctx, id)
}

func RegisterUser(ctx context.Context, user *models.SignUpRequest) error {
	return implementation.RegisterUser(ctx, user)
}

func UpdateUser(ctx context.Context, user *models.UpdateUserRequest) error {
	return implementation.UpdateUser(ctx, user)
}

func DeleteUser(ctx context.Context, id string) error {
	return implementation.DeleteUser(ctx, id)
}

func Close() error {
	return implementation.Close()
}
