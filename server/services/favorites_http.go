package services

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func GetFavorites(c echo.Context) error {
	slog.Info("Getting favorites")
	userID := c.Param("userId")
	if userID == "" {
		return c.String(400, "userId is required")
	}

	favs, err := selectAllFavorites(userID)
	if err != nil {
		slog.Error("Error getting favorites", "selectAllFavorites", err)
		return c.String(500, "Error getting favorites")
	}

	return c.JSON(200, favs)
}

func AddFavorite(c echo.Context) error {
	slog.Info("Adding favorite")
	userID := c.Param("userId")
	if userID == "" {
		return c.String(400, "userId is required")
	}

	var requestBody Anime
	if err := c.Bind(&requestBody); err != nil {
		return c.String(400, "Invalid request body")
	}
	if requestBody.Title == "" || requestBody.Image == "" || requestBody.AnimeId == 0 {
		return c.String(400, "Title, image, and animeId are required")
	}

	err := upsertAnime(requestBody)
	if err != nil {
		slog.Error("Error upserting anime", "upsertAnime", err)
		return c.String(500, "Error upserting anime")
	}
	err = addFavorite(userID, requestBody.AnimeId)
	if err != nil {
		slog.Error("Error adding favorite", "addFavorite", err)
		return c.String(500, "Error adding favorite")
	}

	return c.NoContent(204)
}

func RemoveFavorite(c echo.Context) error {
	slog.Info("Removing favorite")
	userId := c.Param("userId")
	animeId := c.Param("animeId")
	if userId == "" {
		return c.String(400, "userId is required")
	}
	if animeId == "" {
		return c.String(400, "animeId is required")
	}

	err := removeFavorite(userId, animeId)
	if err != nil {
		slog.Error("Error removing favorite", "removeFavorite", err)
		return c.String(500, "Error removing favorite")
	}

	return c.NoContent(204)
}
