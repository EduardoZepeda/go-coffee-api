package repository

import (
	"context"

	"github.com/EduardoZepeda/go-coffee-api/models"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetCoffeeShops(ctx context.Context, page uint64, size uint64) ([]*models.CoffeeShop, error)
	GetCoffeeShopById(ctx context.Context, id string) (*models.CoffeeShop, error)
	CreateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) (string, error)
	DeleteCoffeeShop(ctx context.Context, id string) error
	UpdateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) error
	SearchCoffeeShops(ctx context.Context, query string, page uint64, size uint64) ([]*models.CoffeeShop, error)
	GetNearestCoffeeShop(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.CoffeeShop, error)
	GetUser(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.GetUserResponse, error)
	RegisterUser(ctx context.Context, user *models.SignUpRequest) error
	UpdateUser(ctx context.Context, user *models.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error
	UnfollowUser(ctx context.Context, followUnfollowUserRequest *models.FollowUnfollowRequest) error
	FollowUser(ctx context.Context, followUnfollowUserRequest *models.FollowUnfollowRequest) error
	GetUserFollowing(ctx context.Context, userId string) ([]*models.GetUserResponse, error)
	GetUserFollowers(ctx context.Context, userId string) ([]*models.GetUserResponse, error)
	GetLikedCoffeeShops(ctx context.Context, likes *models.LikesByUserRequest) ([]*models.CoffeeShop, error)
	LikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error
	UnlikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error
	GetUserFeed(ctx context.Context, id string) ([]*models.Feed, error)
	GetCoffeeBag(ctx context.Context, CoffeeBagsList models.CoffeeBagsList) ([]*models.CoffeeBag, error)
	GetCoffeeBagById(ctx context.Context, coffeeBagId string) (*models.CoffeeBag, error)
	CreateCoffeeBag(ctx context.Context, coffeeBag *models.CoffeeBag) (*models.CoffeeBag, error)
	UpdateCoffeeBag(ctx context.Context, coffeeBag *models.CoffeeBag) (*models.CoffeeBag, error)
	DeleteCoffeeBag(ctx context.Context, coffeeShopId string) error
	GetCoffeeBagByCoffeeShop(ctx context.Context, coffeeShopId *models.CoffeeBagByShopId) ([]*models.CoffeeBag, error)
	AddCoffeeBagToCoffeeShop(ctx context.Context, coffeeBagId string, coffeeShopId string) error
	RemoveCoffeeBagFromCoffeeShop(ctx context.Context, coffeeBagId string, coffeeShopId string) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetCoffeeShops(ctx context.Context, page uint64, size uint64) ([]*models.CoffeeShop, error) {
	return implementation.GetCoffeeShops(ctx, page, size)
}

func GetCoffeeShopById(ctx context.Context, id string) (*models.CoffeeShop, error) {
	return implementation.GetCoffeeShopById(ctx, id)
}

func CreateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) (string, error) {
	return implementation.CreateCoffeeShop(ctx, shopRequest)
}

func DeleteCoffeeShop(ctx context.Context, id string) error {
	return implementation.DeleteCoffeeShop(ctx, id)
}

func UpdateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) error {
	return implementation.UpdateCoffeeShop(ctx, shopRequest)
}

func SearchCoffeeShops(ctx context.Context, query string, page uint64, size uint64) ([]*models.CoffeeShop, error) {
	return implementation.SearchCoffeeShops(ctx, query, page, size)
}

func GetNearestCoffeeShop(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.CoffeeShop, error) {
	return implementation.GetNearestCoffeeShop(ctx, UserCoordinates)
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

func GetUserFollowing(ctx context.Context, userId string) ([]*models.GetUserResponse, error) {
	return implementation.GetUserFollowing(ctx, userId)
}

func GetUserFollowers(ctx context.Context, userId string) ([]*models.GetUserResponse, error) {
	return implementation.GetUserFollowers(ctx, userId)
}

func UnfollowUser(ctx context.Context, followUnfollowUserRequest *models.FollowUnfollowRequest) error {
	return implementation.UnfollowUser(ctx, followUnfollowUserRequest)
}

func FollowUser(ctx context.Context, followUnfollowUserRequest *models.FollowUnfollowRequest) error {
	return implementation.FollowUser(ctx, followUnfollowUserRequest)
}

func GetLikedCoffeeShops(ctx context.Context, likes *models.LikesByUserRequest) ([]*models.CoffeeShop, error) {
	return implementation.GetLikedCoffeeShops(ctx, likes)
}

func LikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error {
	return implementation.LikeCoffeeShop(ctx, like)
}

func UnlikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error {
	return implementation.UnlikeCoffeeShop(ctx, like)
}

func GetUserFeed(ctx context.Context, id string) ([]*models.Feed, error) {
	return implementation.GetUserFeed(ctx, id)
}

func GetCoffeeBag(ctx context.Context, CoffeeBagsList models.CoffeeBagsList) ([]*models.CoffeeBag, error) {
	return implementation.GetCoffeeBag(ctx, CoffeeBagsList)
}

func GetCoffeeBagById(ctx context.Context, coffeeBagId string) (*models.CoffeeBag, error) {
	return implementation.GetCoffeeBagById(ctx, coffeeBagId)
}

func CreateCoffeeBag(ctx context.Context, coffeeBag *models.CoffeeBag) (*models.CoffeeBag, error) {
	return implementation.CreateCoffeeBag(ctx, coffeeBag)
}

func UpdateCoffeeBag(ctx context.Context, coffeeBag *models.CoffeeBag) (*models.CoffeeBag, error) {
	return implementation.UpdateCoffeeBag(ctx, coffeeBag)
}

func DeleteCoffeeBag(ctx context.Context, coffeeBagId string) error {
	return implementation.DeleteCoffeeBag(ctx, coffeeBagId)
}

func GetCoffeeBagByCoffeeShop(ctx context.Context, coffeeShopId *models.CoffeeBagByShopId) ([]*models.CoffeeBag, error) {
	return implementation.GetCoffeeBagByCoffeeShop(ctx, coffeeShopId)
}

func AddCoffeeBagToCoffeeShop(ctx context.Context, coffeeBagId string, coffeeShopId string) error {
	return implementation.AddCoffeeBagToCoffeeShop(ctx, coffeeBagId, coffeeShopId)
}

func RemoveCoffeeBagFromCoffeeShop(ctx context.Context, coffeeBagId string, coffeeShopId string) error {
	return implementation.RemoveCoffeeBagFromCoffeeShop(ctx, coffeeBagId, coffeeShopId)
}

func Close() error {
	return implementation.Close()
}
