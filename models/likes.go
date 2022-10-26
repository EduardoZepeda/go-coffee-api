package models

type LikesByUserRequest struct {
	Page   uint64
	Size   uint64
	UserId string
}

type LikeUnlikeCoffeeShopRequest struct {
	UserId string `db:"UserId" json:"userId,omitempty"`
	ShopId string `db:"ShopId" json:"shopId,omitempty"`
}
