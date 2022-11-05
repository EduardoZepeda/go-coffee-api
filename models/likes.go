package models

type LikesByUserRequest struct {
	UserId string
	Pagination
}

type LikeUnlikeCoffeeShopRequest struct {
	UserId string `db:"UserId" json:"userId,omitempty"`
	ShopId string `db:"ShopId" json:"shopId,omitempty"`
}
