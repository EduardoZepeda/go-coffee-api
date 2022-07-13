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

func (repo *PostgresRepository) GetCafes(ctx context.Context, page uint64, size uint64) ([]*models.Shop, error) {
	var shops []*models.Shop
	err := repo.db.SelectContext(ctx, &shops, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY name DESC LIMIT $1 OFFSET $2;", size, page*size)
	return shops, err
}

func (repo *PostgresRepository) GetCafeById(ctx context.Context, id string) (*models.Shop, error) {
	var shop models.Shop
	err := repo.db.GetContext(ctx, &shop, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE id = $1;", id)
	return &shop, err
}

func (repo *PostgresRepository) CreateCafe(ctx context.Context, shopRequest *models.CreateShop) error {
	_, err := repo.db.NamedExecContext(ctx, "INSERT INTO shops_shop (name, location, address, rating) VALUES (:Name, :Location, :Address, :Rating);", shopRequest)
	return err
}

func (repo *PostgresRepository) UpdateCafe(ctx context.Context, shopRequest *models.InsertShop) error {
	_, err := repo.db.NamedExecContext(ctx, "UPDATE shops_shop SET name = :name, location = :Location, address = :Address, rating = :Rating WHERE id = :ID;", shopRequest)
	return err
}

func (repo *PostgresRepository) DeleteCafe(ctx context.Context, id string) error {
	_, err := repo.db.NamedExecContext(ctx, "DELETE FROM shops_shop WHERE id = :id;", map[string]interface{}{"id": id})
	return err
}

func (repo *PostgresRepository) SearchCafe(ctx context.Context, query string, page uint64, size uint64) ([]*models.Shop, error) {
	var cafes []*models.Shop
	err := repo.db.SelectContext(ctx, &cafes, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE to_tsvector(COALESCE(name, '') || COALESCE(address, '')) @@ plainto_tsquery($1) LIMIT $2 OFFSET $3;", query, size, page*size)
	return cafes, err
}

func (repo *PostgresRepository) GetNearestCafes(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.Shop, error) {
	var cafes []*models.Shop
	err := repo.db.SelectContext(ctx, &cafes, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY location <-> ST_SetSRID(ST_MakePoint(:Latitude, :Longitude), 4326) LIMIT 10;", UserCoordinates)
	return cafes, err
}

func (repo *PostgresRepository) GetUser(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := repo.db.GetContext(ctx, &user, "SELECT id, email, password FROM auth_user WHERE email = $1;", email)
	return &user, err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
