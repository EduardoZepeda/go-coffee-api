package handlers

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/types"
)

// UserFeed godoc
// @Summary      The active user's feed
// @Description  This route returns the user's last ten feed items. Each item consists of a subject, an action and a destinatary
// @Tags         feed
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started."
// @Success      200 {array}  models.Feed
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /feed [get]
func GetUserFeed(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId := ctx.Value("userId")
		feed, err := app.Repo.GetUserFeed(ctx, userId.(string))
		if err != nil {
			app.Logger.Println(err)
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		if len(feed) == 0 {
			app.Respond(w, []int{}, http.StatusOK)
			return
		}
		app.Respond(w, feed, http.StatusOK)
		return
	}
}
