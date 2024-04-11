package models

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type Banner struct {
	Id        int64
	TagIds    []int64 `json:"tag_ids"`
	FeatureId int64   `json:"feature_id"`
	Content   string  `json:"content"`
	IsActive  bool    `json:"is_active"`
}

type BannerPatch struct {
	TagIds    pq.Int64Array
	FeatureId *int64
	Content   *string
	IsActive  *bool
}

type BannerFilter struct {
	FeatureId *int64
	TagId     *int64
}

func GetBanner(db Database, tagId, featureId int64) (*Banner, error) {
	const query = "SELECT id, \"content\", tag_ids, is_active FROM banners" +
		" WHERE $1 = ANY(tag_ids) AND feature_id=$2 LIMIT 1"
	row := db.QueryRow(query, tagId, featureId)
	b := new(Banner)
	b.FeatureId = featureId
	tags := pq.Int64Array(b.TagIds)
	err := row.Scan(&b.Id, &b.Content, &tags, &b.IsActive)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Banner) GetBanners(db Database, filter *BannerFilter, limit *int, offset *int) ([]Banner, error) {
	qb := new(strings.Builder)
	qb.WriteString("SELECT * FROM banners")
	filters := []string{}
	if filter.FeatureId != nil {
		filters = append(filters, fmt.Sprintf("feature_id = %d", *filter.FeatureId))
	}
	if filter.TagId != nil {
		filters = append(filters, fmt.Sprintf("%d = ANY(tag_ids)", *filter.TagId))
	}
	if len(filters) > 0 {
		qb.WriteString(" WHERE ")
		qb.WriteString(strings.Join(filters, " AND "))
	}
	if limit != nil {
		qb.WriteString(fmt.Sprintf(" LIMIT %d", *limit))
	}
	if offset != nil {
		qb.WriteString(fmt.Sprintf(" OFFSET %d", *offset))
	}
	return []Banner{}, nil
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
