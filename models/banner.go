package models

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type Banner struct {
	Id        int64
	TagIds    []int  `json:"tag_ids"`
	FeatureId int    `json:"feature_id"`
	Content   string `json:"content"`
	IsActive  bool   `json:"is_active"`
}

type BannerPatch struct {
	TagIds    pq.Int64Array
	FeatureId *int
	Content   *string
	IsActive  *bool
}

func (b *Banner) Insert(db Database) (int64, error) {
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

func PatchBanner(db Database, id int64, patch *BannerPatch) error {
	params := []any{}
	columns := []string{}
	if patch.Content != nil {
		params = append(params, patch.Content)
		columns = append(columns, "content")
	}
	if patch.FeatureId != nil {
		params = append(params, patch.FeatureId)
		columns = append(columns, "feature_id")
	}
	if patch.TagIds != nil {
		params = append(params, patch.TagIds)
		columns = append(columns, "tag_ids")
	}
	if patch.IsActive != nil {
		params = append(params, patch.IsActive)
		columns = append(columns, "is_active")
	}
	if len(params) == 0 {
		return nil
	}
	var qb = new(strings.Builder)
	qb.WriteString(fmt.Sprintf("UPDATE banners SET %s = $1", columns[0]))
	for i := 1; i < len(params); i++ {
		qb.WriteString(fmt.Sprintf(", %s = $%d", columns[i], i+1))
	}
	qb.WriteString(fmt.Sprintf(" WHERE id = %d", id))
	fmt.Println(qb.String())
	_, err := db.Exec(qb.String(), params...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBanner(db Database, id int64) error {
	const query = "DELETE FROM banners WHERE id=$1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
