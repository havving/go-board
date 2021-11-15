package models

import (
	"github.com/jmoiron/sqlx"
	"go-board/pkg/database"
)

// BoardBlock 게시판
type BoardBlock struct {
	BoardIdx        int    `db:"board_idx" json:"board_idx"`
	Title           string `db:"title" json:"title"`
	Contents        string `db:"contents" json:"contents"`
	HitCnt          int    `db:"hit_cnt" json:"hit_cnt"`
	CreatorId       string `db:"creator_id" json:"creator_id"`
	CreatedDateTime string `db:"created_datetime" json:"created_date_time"`
	UpdaterId       string `db:"updater_id" json:"updater_id"`
	UpdatedDateTime string `db:"updated_datetime" json:"updated_date_time"`
}

var SQL *database.SQL
var Board BoardBlock

// entityList SELECT Column
func entityList(rows *sqlx.Rows) (list []BoardBlock, err error) {
	for rows.Next() {
		entity := BoardBlock{}
		err = rows.StructScan(&entity)
		if err != nil {
			list = nil
			return
		}
		list = append(list, entity)
	}

	return
}

func Setup() {
	SQL = new(database.SQL)
	SQL.Connect()
}
