package models

type FollowUnfollowRequest struct {
	UserFromId string `db:"UserFromId" json:"userFromId"`
	UserToId   string `db:"UserToId" json:"userToId"`
}
