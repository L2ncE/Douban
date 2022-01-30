package dao

import "douban/model"

// SelectMovieById 通过id来搜索电影
func SelectMovieById(movieId int) (model.Movie, error) {
	var movie model.Movie

	row := dB.QueryRow("SELECT id, name, year, director, screenwriter, starring, type, country, language, length, imdb, starnum, score, star, havewatched, wanttowatch, synopsis, URL, peopleURL, NameInfo, CoverInfo FROM movie WHERE id = ? ", movieId)
	if row.Err() != nil {
		return movie, row.Err()
	}

	err := row.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Type, &movie.Country, &movie.Language, &movie.Length, &movie.IMDb, &movie.StarNum, &movie.Score, &movie.Star, &movie.HaveWatched, &movie.WantToWatch, &movie.Synopsis, &movie.URL, &movie.PeopleURL, &movie.NameInfo, &movie.CoverInfo)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

// SelectMovie1 查找主页1
func SelectMovie1() ([]model.MovieBrief, error) {
	var movies []model.MovieBrief
	rows, err := dB.Query("SELECT id, name, URL FROM movie WHERE Id BETWEEN 1 AND 35")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.MovieBrief

		err = rows.Scan(&movie.Id, &movie.Name, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovie2 查找主页2
func SelectMovie2() ([]model.MovieBrief, error) {
	var movies []model.MovieBrief
	rows, err := dB.Query("SELECT id, name, URL FROM movie WHERE Id BETWEEN 36 AND 85")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.MovieBrief

		err = rows.Scan(&movie.Id, &movie.Name, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovie3 查找主页3
func SelectMovie3() ([]model.MovieBrief, error) {
	var movies []model.MovieBrief
	rows, err := dB.Query("SELECT id, name, URL FROM movie WHERE Id BETWEEN 86 AND 135")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.MovieBrief

		err = rows.Scan(&movie.Id, &movie.Name, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
