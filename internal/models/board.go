package models

import "log"

func (b BoardBlock) SelectBoardList() (list []BoardBlock, err error) {
	rows, err := SQL.Queryx("SELECT board_idx, title, hit_cnt FROM board ORDER BY board_idx DESC")
	if err != nil {
		log.Fatalln(err)
		return
	}

	list, err = entityList(rows)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
