package database

import (
	"context"
	"net/url"
	"os"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository() (*PostgresRepository, error) {
	q := make(url.Values)
	q.Set("sslmode", "require")
	q.Set("timezone", "utc")
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Host:     os.Getenv("DB_HOST"),
		Path:     os.Getenv("DB_PATH"),
		RawQuery: q.Encode(),
	}

	db, err := sqlx.Connect("postgres", u.String())
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) GetCoffeeShops(ctx context.Context, page uint64, size uint64) ([]*models.CoffeeShop, error) {
	var shops []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &shops, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY name DESC LIMIT $1 OFFSET $2;", size, page*size)
	return shops, err
}

func (repo *PostgresRepository) GetCoffeeShopById(ctx context.Context, id string) (*models.CoffeeShop, error) {
	var shop models.CoffeeShop
	err := repo.db.GetContext(ctx, &shop, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE id = $1;", id)
	return &shop, err
}

func (repo *PostgresRepository) CreateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) error {
	_, err := repo.db.NamedExecContext(ctx, "INSERT INTO shops_shop (name, location, address, rating) VALUES (:Name, :Location, :Address, :Rating);", shopRequest)
	return err
}

func (repo *PostgresRepository) UpdateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) error {
	_, err := repo.db.NamedExecContext(ctx, "UPDATE shops_shop SET name = :Name, location = :Location, address = :Address, rating = :Rating WHERE id = :ID;", shopRequest)
	return err
}

func (repo *PostgresRepository) DeleteCoffeeShop(ctx context.Context, id string) error {
	_, err := repo.db.NamedExecContext(ctx, "DELETE FROM shops_shop WHERE id = :id;", map[string]interface{}{"id": id})
	return err
}

func (repo *PostgresRepository) SearchCoffeeShops(ctx context.Context, query string, page uint64, size uint64) ([]*models.CoffeeShop, error) {
	var cafes []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &cafes, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE to_tsvector(COALESCE(LOWER(name), '') || COALESCE(LOWER(address), '')) @@ plainto_tsquery($1) LIMIT $2 OFFSET $3;", query, size, page*size)
	return cafes, err
}

func (repo *PostgresRepository) GetNearestCoffeeShop(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.CoffeeShop, error) {
	var cafes []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &cafes, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY location <-> ST_SetSRID(ST_MakePoint($1, $2), 4326) LIMIT 10;", UserCoordinates.Latitude, UserCoordinates.Longitude)
	return cafes, err
}

func (repo *PostgresRepository) GetUser(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := repo.db.GetContext(ctx, &user, "SELECT id, email, password FROM accounts_user WHERE email = $1;", email)
	return &user, err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.GetUserResponse, error) {
	var user models.GetUserResponse
	// Null values cannot be converted to string automatically, thus, we need to handle null values from db
	// COALESCE will return the first not null value, and it must be used together with as <field>, otherwise it will fail
	err := repo.db.GetContext(ctx, &user, "SELECT id, email, username, first_name, last_name, COALESCE(bio, '') as bio FROM accounts_user WHERE id = $1;", id)
	return &user, err
}

func (repo *PostgresRepository) RegisterUser(ctx context.Context, user *models.SignUpRequest) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO accounts_user (is_superuser, password, username, email, is_staff, is_active, first_name, last_name, date_joined) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, current_timestamp);", false, user.HashedPassword, user.Username, user.Email, false, true, "", "")
	return err
}

func (repo *PostgresRepository) UpdateUser(ctx context.Context, user *models.UpdateUserRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "UPDATE accounts_user SET username = :Username, bio = :Bio, first_name = :FirstName, last_name = :LastName WHERE id = :Id;", user)
	return err
}

func (repo *PostgresRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM accounts_user WHERE id = $1;", id)
	return err
}

func (repo *PostgresRepository) FollowUser(ctx context.Context, followUnfollowUserRequest *models.FollowUnfollowRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "INSERT INTO accounts_contact (created, user_from_id, user_to_id) VALUES (current_timestamp, :UserFromId, :UserToId);", followUnfollowUserRequest)
	return err
}

func (repo *PostgresRepository) UnfollowUser(ctx context.Context, followUnfollowUserRequest *models.FollowUnfollowRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "DELETE FROM accounts_contact WHERE user_from_id = :UserFromId AND user_to_id = :UserToId;", followUnfollowUserRequest)
	return err
}

func (repo *PostgresRepository) GetUserFollowing(ctx context.Context, userId string) ([]*models.GetUserResponse, error) {
	var users []*models.GetUserResponse
	err := repo.db.SelectContext(ctx, &users, "SELECT accounts_user.id, username, first_name, last_name, COALESCE(bio, '') as bio, email FROM accounts_user INNER JOIN accounts_contact ON accounts_contact.user_to_id = accounts_user.id WHERE accounts_contact.user_from_id = $1;", userId)
	return users, err
}

func (repo *PostgresRepository) GetUserFollowers(ctx context.Context, userId string) ([]*models.GetUserResponse, error) {
	var users []*models.GetUserResponse
	err := repo.db.SelectContext(ctx, &users, "SELECT accounts_user.id, username, first_name, last_name, COALESCE(bio, '') as bio, email FROM accounts_user INNER JOIN accounts_contact ON accounts_contact.user_from_id = accounts_user.id WHERE accounts_contact.user_to_id = $1;", userId)
	return users, err
}

func (repo *PostgresRepository) GetLikedCoffeeShops(ctx context.Context, likes *models.LikesByUserRequest) ([]*models.CoffeeShop, error) {
	var coffeeShops []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &coffeeShops, "SELECT shops_shop.id, name, location, address, rating, created_date, modified_date FROM shops_shop INNER JOIN shops_shop_likes ON shops_shop_likes.shop_id = shops_shop.id WHERE shops_shop_likes.user_id = $1 LIMIT $2 OFFSET $3;", likes.UserId, likes.Size, likes.Size*likes.Page)
	return coffeeShops, err
}

func (repo *PostgresRepository) LikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "INSERT INTO shops_shop_likes (shop_id, user_id) VALUES (:ShopId, :UserId);", like)
	return err
}

func (repo *PostgresRepository) UnlikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "DELETE FROM shops_shop_likes WHERE shop_id = :ShopId AND user_id = :UserId;", like)
	return err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
