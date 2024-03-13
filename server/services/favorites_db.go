package services

import (
	"task/system"
)

type Favorite struct {
	Created string `json:"created"`
	UserId  string `json:"userId"`
	AnimeId int    `json:"animeId"`
}

type FavoriteWithAnimeDetails struct {
	Favorite
	Title string `json:"title"`
	Image string `json:"image"`
}

func dest(fav *FavoriteWithAnimeDetails) []interface{} {
	return []interface{}{
		&fav.Favorite.Created,
		&fav.Favorite.UserId,
		&fav.Favorite.AnimeId,
		&fav.Title,
		&fav.Image,
	}
}

func selectAllFavorites(userID string) ([]FavoriteWithAnimeDetails, error) {
	rows, err := system.Db.Query(`
		SELECT f.created, f.user_id, f.anime_id, a.title, a.image 
		FROM favorites AS f 
		INNER JOIN anime AS a ON f.anime_id = a.id 
		WHERE f.user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favs []FavoriteWithAnimeDetails = make([]FavoriteWithAnimeDetails, 0)
	for rows.Next() {
		fav := FavoriteWithAnimeDetails{}
		err = rows.Scan(dest(&fav)...)
		if err != nil {
			return nil, err
		}
		favs = append(favs, fav)
	}

	return favs, nil
}

func addFavorite(userID string, animeID int) error {
	stmt, err := system.Db.Prepare("INSERT INTO favorites(user_id, anime_id) VALUES(?, ?) ON CONFLICT DO NOTHING")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, animeID)
	if err != nil {
		return err
	}

	return nil
}

func removeFavorite(userID, animeID string) error {
	stmt, err := system.Db.Prepare("DELETE FROM favorites WHERE user_id = ? AND anime_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, animeID)
	if err != nil {
		return err
	}

	return nil
}

type Anime struct {
	AnimeId int    `json:"animeId"`
	Title   string `json:"title"`
	Image   string `json:"image"`
}

func upsertAnime(anime Anime) error {
	stmt, err := system.Db.Prepare("INSERT INTO anime (id, title, image) VALUES (?, ?, ?) ON CONFLICT(id) DO UPDATE SET title = EXCLUDED.title, image = EXCLUDED.image, updated = CURRENT_TIMESTAMP")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(anime.AnimeId, anime.Title, anime.Image)
	if err != nil {
		return err
	}

	return nil
}
