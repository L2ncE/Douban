package dao

import "douban/model"

// SelectCelebrityById 通过id来搜索影人
func SelectCelebrityById(Id int) (model.Celebrity, error) {
	var celebrity model.Celebrity

	row := dB.QueryRow("SELECT id, Name, Info, synopsis, urlinfo1, urlinfo2, urlinfo3, award FROM celebrity WHERE id = ? ", Id)
	if row.Err() != nil {
		return celebrity, row.Err()
	}

	err := row.Scan(&celebrity.Id, &celebrity.Name, &celebrity.Info, &celebrity.Synopsis, &celebrity.URLInfo1, &celebrity.URLInfo2, &celebrity.URLInfo3, &celebrity.Award)
	if err != nil {
		return celebrity, err
	}

	return celebrity, nil
}
