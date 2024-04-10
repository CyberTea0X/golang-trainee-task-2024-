package models

import "github.com/lib/pq"

type Banner struct {
	Id        int
	TagIds    []int  `json:"tag_ids"`
	FeatureId int    `json:"feature_id"`
	Content   string `json:"content"`
	IsActive  bool   `json:"is_active"`
}

func (b *Banner) InsertToDB(db Database) (int64, error) {
	const query = "INSERT INTO banners (id, \"content\", feature_id, tag_ids, is_active)" +
		"VALUES (DEFAULT,$1,$2,$3,$4) RETURNING id"
	var lastInsertId int64
	row := db.QueryRow(query, b.Content, b.FeatureId, pq.Array(b.TagIds), b.IsActive)
	err := row.Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func DeleteBannerFromDB(db Database, id int64) error {
	const query = "DELETE FROM banners WHERE id=$1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
