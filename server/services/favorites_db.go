package services

import "task/system"

type Favorite struct {
	ID      int    `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	// ...
}

func dest(fav *Favorite) []interface{} {
    return []interface{}{
        &fav.ID,
        &fav.Created,
        &fav.Updated,
        // ...
    }
}

func selectAllFavorites() ([]Favorite, error) {
    rows, err := system.Db.Query("select * from favorites")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var favs []Favorite = make([]Favorite, 0)
    for rows.Next() {
        fav := Favorite{}
        err = rows.Scan(dest(&fav)...)
        if err != nil {
            return nil, err
        }
        favs = append(favs, fav)
    }
    return favs, nil
}
