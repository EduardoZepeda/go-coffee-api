package database

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func GenerateConnectionString() url.URL {
	q := make(url.Values)
	q.Set("sslmode", "require")
	q.Set("timezone", "utc")
	u := url.URL{
		// Remember to update environmental variables at vercel
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Host:     os.Getenv("DB_HOST"),
		Path:     os.Getenv("DB_PATH"),
		RawQuery: q.Encode(),
	}
	return u
}

func NewPostgresRepository() (*PostgresRepository, error) {
	u := GenerateConnectionString()
	db, err := sqlx.Connect("postgres", u.String())
	if err != nil {
		log.Println(u.String(), err)
		return nil, err
	}
	db.SetMaxOpenConns(25)                 // The default is 0 (unlimited)
	db.SetMaxIdleConns(25)                 // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(1 * time.Minute) // 0, connections are reused forever.
	// Return an error if opening the database takes longer than 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) GetCoffeeShops(ctx context.Context, page uint64, size uint64) ([]*models.CoffeeShop, error) {
	var shops []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &shops, "SELECT id, name, location, address, roaster, city, rating, created_date, modified_date FROM shops_shop ORDER BY created_date DESC LIMIT $1 OFFSET $2;", size, page*size)
	return shops, err
}

func (repo *PostgresRepository) GetCoffeeShopById(ctx context.Context, id string) (*models.CoffeeShop, error) {
	var shop models.CoffeeShop
	err := repo.db.GetContext(ctx, &shop, "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE id = $1;", id)
	return &shop, err
}

func (repo *PostgresRepository) CreateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) (string, error) {
	var id string
	err := repo.db.QueryRowContext(ctx, "INSERT INTO shops_shop (name, location, city, roaster, address, rating, created_date, modified_date) VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp) RETURNING id;", shopRequest.Name, shopRequest.Location, shopRequest.City, shopRequest.Roaster, shopRequest.Address, shopRequest.Rating).Scan(&id)
	return id, err
}

func (repo *PostgresRepository) UpdateCoffeeShop(ctx context.Context, shopRequest *models.CoffeeShop) error {
	_, err := repo.db.NamedExecContext(ctx, "UPDATE shops_shop SET name = :name, location = :location, address = :address, rating = :rating, roaster = :roaster, modified_date = current_timestamp WHERE id = :id;", shopRequest)
	return err
}

func (repo *PostgresRepository) DeleteCoffeeShop(ctx context.Context, id string) error {
	_, err := repo.db.NamedExecContext(ctx, "DELETE FROM shops_shop WHERE id = :id;", map[string]interface{}{"id": id})
	return err
}

func (repo *PostgresRepository) SearchCoffeeShops(ctx context.Context, query string, page uint64, size uint64) ([]*models.CoffeeShop, error) {
	var cafes []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &cafes, "SELECT id, name, location, address, roaster, rating, created_date, modified_date FROM shops_shop WHERE to_tsvector(COALESCE(LOWER(name), '') || COALESCE(LOWER(address), '') || COALESCE(LOWER(content), '')) @@ plainto_tsquery($1) LIMIT $2 OFFSET $3;", query, size, page*size)
	return cafes, err
}

func (repo *PostgresRepository) GetNearestCoffeeShop(ctx context.Context, UserCoordinates *models.UserCoordinates) ([]*models.CoffeeShop, error) {
	var cafes []*models.CoffeeShop
	err := repo.db.SelectContext(ctx, &cafes, "SELECT id, name, location, address, roaster, rating, created_date, modified_date FROM shops_shop ORDER BY location <-> ST_SetSRID(ST_MakePoint($1, $2), 4326) LIMIT 10;", UserCoordinates.Latitude, UserCoordinates.Longitude)
	return cafes, err
}

func (repo *PostgresRepository) GetUser(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := repo.db.GetContext(ctx, &user, "SELECT id, email, password, is_staff FROM accounts_user WHERE email = $1;", email)
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
	if err != nil {
		// Check for user constraints on database
		if strings.Contains(err.Error(), "accounts_user_username_key") || strings.Contains(err.Error(), "unique-email") {
			return errors.New("An user with that username or email address already exists")
		}
	}
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
	if err != nil {
		// Check for user constraints on database
		if strings.Contains(err.Error(), "userTo-userFrom") {
			return errors.New("You are already following this user")
		}
	}
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
	err := repo.db.SelectContext(ctx, &coffeeShops, "SELECT shops_shop.id, name, location, address, rating, city, roaster, created_date, modified_date FROM shops_shop INNER JOIN shops_shop_likes ON shops_shop_likes.shop_id = shops_shop.id WHERE shops_shop_likes.user_id = $1 LIMIT $2 OFFSET $3;", likes.UserId, likes.Size, likes.Size*likes.Page)
	return coffeeShops, err
}

func (repo *PostgresRepository) LikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "INSERT INTO shops_shop_likes (shop_id, user_id) VALUES (:ShopId, :UserId);", like)
	if err != nil {
		// Check for user constraints on database
		if strings.Contains(err.Error(), "shops_shop_likes_shop_id_user_id_09e87394_uniq") {
			return errors.New("You already like that coffee shop. You can't like it twice.")
		}
	}
	return err
}

func (repo *PostgresRepository) UnlikeCoffeeShop(ctx context.Context, like *models.LikeUnlikeCoffeeShopRequest) error {
	_, err := repo.db.NamedExecContext(ctx, "DELETE FROM shops_shop_likes WHERE shop_id = :ShopId AND user_id = :UserId;", like)
	return err
}

func (repo *PostgresRepository) GetUserFeed(ctx context.Context, id string) ([]*models.Feed, error) {
	var feed []*models.Feed
	// target_ct_id makes reference to a table that keeps a register of the used models
	// Since feeds consists only of shops and users (7 and 9), the target_ct_id are hardcoded
	err := repo.db.SelectContext(ctx, &feed, `SELECT accounts_user.username, feeds_action.action, 
	CASE WHEN feeds_action.target_ct_id = 7 THEN shops_shop.name 
	WHEN feeds_action.target_ct_id = 9 THEN (SELECT username FROM accounts_user WHERE feeds_action.target_id = id) END AS target 
	FROM feeds_action JOIN accounts_user ON feeds_action.user_id = accounts_user.id  JOIN shops_shop on feeds_action.target_id = shops_shop.id WHERE accounts_user.id = $1 ORDER BY feeds_action.created DESC LIMIT 20;`, id)
	return feed, err
}

func (repo *PostgresRepository) GetCoffeeBags(ctx context.Context, CoffeeBagsList models.CoffeeBagsList) ([]*models.CoffeeBag, error) {
	var coffeeBags []*models.CoffeeBag
	rows, err := repo.db.QueryxContext(ctx, "SELECT id, brand, species, origin FROM shops_coffeebag LIMIT $1 OFFSET $2;", CoffeeBagsList.Size, CoffeeBagsList.Page*CoffeeBagsList.Size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item models.CoffeeBag
		err = rows.StructScan(&item)
		item.Species = COFFEE_SPECIES[item.Species]
		item.Origin = STATE_CHOICES[item.Origin]
		coffeeBags = append(coffeeBags, &item)
	}
	err = rows.Err()
	return coffeeBags, err
}

func (repo *PostgresRepository) CreateCoffeeBag(ctx context.Context, coffeeBag *models.CoffeeBag) (*models.CoffeeBag, error) {
	var coffeBagId string
	err := repo.db.QueryRowContext(ctx, "INSERT INTO shops_coffeebag (brand, species, origin) VALUES($1, $2, $3) RETURNING id;", coffeeBag.Brand, coffeeBag.Species, coffeeBag.Origin).Scan(&coffeBagId)
	coffeeBag.ID = coffeBagId
	return coffeeBag, err
}

func (repo *PostgresRepository) GetCoffeeBagById(ctx context.Context, coffeeBagId string) (*models.CoffeeBag, error) {
	var coffeeShopBag models.CoffeeBag
	err := repo.db.GetContext(ctx, &coffeeShopBag, "SELECT id, brand, species, origin FROM shops_coffeebag WHERE id = $1;", coffeeBagId)
	return &coffeeShopBag, err
}

func (repo *PostgresRepository) UpdateCoffeeBag(ctx context.Context, coffeeBag *models.CoffeeBag) (*models.CoffeeBag, error) {
	_, err := repo.db.NamedExecContext(ctx, "UPDATE shops_coffeebag SET brand = brand, species = species, origin = origin WHERE id = id;", coffeeBag)
	return coffeeBag, err
}

func (repo *PostgresRepository) DeleteCoffeeBag(ctx context.Context, coffeeBagId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM shops_coffeebag WHERE id = $1;", coffeeBagId)
	return err
}

func (repo *PostgresRepository) GetCoffeeBagByCoffeeShop(ctx context.Context, coffeeShopId *models.CoffeeBagByShopId) ([]*models.CoffeeBag, error) {
	var coffeeBags []*models.CoffeeBag
	rows, err := repo.db.QueryxContext(ctx, "SELECT shops_coffeebag.id, brand, species, origin FROM shops_coffeebag INNER JOIN shops_coffeebag_coffee_shop ON shops_coffeebag.id = shops_coffeebag_coffee_shop.coffeebag_id WHERE shops_coffeebag_coffee_shop.shop_id = $1 LIMIT $2 OFFSET $3;", coffeeShopId.CoffeeShopId, coffeeShopId.Size, coffeeShopId.Page*coffeeShopId.Size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item models.CoffeeBag
		err = rows.StructScan(&item)
		item.Species = COFFEE_SPECIES[item.Species]
		item.Origin = STATE_CHOICES[item.Origin]
		coffeeBags = append(coffeeBags, &item)
	}
	err = rows.Err()
	return coffeeBags, err
}

func (repo *PostgresRepository) AddCoffeeBagToCoffeeShop(ctx context.Context, coffeeBagId string, coffeeShopId string) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO shops_coffeebag_coffee_shop (coffeebag_id, shop_id) VALUES($1, $2);", coffeeBagId, coffeeShopId)
	if err != nil {
		if strings.Contains(err.Error(), "shops_coffeebag_coffee_shop_coffeebag_id_shop_id_2d92af17_uniq") {
			return errors.New("That coffee bag is already registered as a product of that coffee shop")
		}
	}
	return err
}

func (repo *PostgresRepository) RemoveCoffeeBagFromCoffeeShop(ctx context.Context, coffeeBagId string, coffeeShopId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM shops_coffeebag_coffee_shop WHERE coffeebag_id = $1 and shop_id = $2;", coffeeBagId, coffeeShopId)
	return err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
