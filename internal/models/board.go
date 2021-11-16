package models

import "log"

// SelectBoardList 전체 게시판 조회
func (b BoardBlock) SelectBoardList() (list []BoardBlock, err error) {
	rows, err := SQL.Queryx("SELECT board_idx, title, hit_cnt FROM t_board ORDER BY board_idx DESC")
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

// InsertBoard 게시판 글쓰기
func (b BoardBlock) InsertBoard(title, contents, now string) (err error) {
	_, err = SQL.Exec("INSERT INTO t_board(title, contents, created_datetime, creator_id) VALUES($1,$2,$3,$4)",
		title, contents, now, "admin")
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}

// SelectBoardDetail 게시판 상세 조회
func (b BoardBlock) SelectBoardDetail(id int) (*BoardBlock, error) {
	err := SQL.QueryRowx("SELECT board_idx, title, contents, hit_cnt, created_datetime, creator_id FROM t_board "+
		"WHERE board_idx = $1", id).StructScan(&b)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &b, nil
}

// UpdateHitCount 게시판 조회수 업데이트
/*func (b BoardBlock) UpdateHitCount(id int) (err error) {
	_, err = SQL.Exec("UPDATE t_board SET hit_cnt = hit_cnt + 1 WHERE board_idx = $1", id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}*/

// UpdateBoard 게시판 수정
func (b BoardBlock) UpdateBoard(title, contents, now string, id int) (err error) {
	_, err = SQL.Exec("UPDATE t_board SET title=$1, contents=$2, updated_datetime=$3, updater_id=$4 WHERE board_idx=$5",
		title, contents, now, "admin", id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}

// DeleteBoard 게시판 삭제
func (b BoardBlock) DeleteBoard(id int) (err error) {
	_, err = SQL.Exec("DELETE FROM t_board WHERE board_idx = $1", id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}
