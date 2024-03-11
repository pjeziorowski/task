package services

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func GetFavorites(c echo.Context) error {
    slog.Info("Getting favorites")
    favs, err := selectAllFavorites()
    if err != nil {
        slog.Error("Error getting favorites", "selectAllFavorites", err)
        return c.String(500, "Error getting favorites")
    }
    return c.JSON(200, favs)
}
