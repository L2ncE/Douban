package dao

import "douban/model"

// SelectMovieById 通过id来搜索电影
func SelectMovieById(movieId int) (model.Movie, error) {
	var movie model.Movie

	row := dB.QueryRow("SELECT id, name, year, director, screenwriter, starring, type, country, language, length, imdb, starnum, score, onestar, twostar, threestar, fourstar, fivestar, havewatched, wanttowatch, synopsis, URL, peopleURL FROM movie WHERE id = ? ", movieId)
	if row.Err() != nil {
		return movie, row.Err()
	}

	err := row.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Type, &movie.Country, &movie.Language, &movie.Length, &movie.IMDb, &movie.StarNum, &movie.Score, &movie.OneStar, &movie.TwoStar, &movie.ThreeStar, &movie.FourStar, &movie.FiveStar, &movie.HaveWatched, &movie.WantToWatch, &movie.Synopsis, &movie.URL, &movie.PeopleURL)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

// SelectMovie 查找电影
func SelectMovie() ([]model.Movie, error) {
	var movies []model.Movie
	rows, err := dB.Query("SELECT id, name, year, director, screenwriter, starring, type, country, language, length, imdb, starnum, score, onestar, twostar, threestar, fourstar, fivestar, havewatched, wanttowatch, synopsis, URL, peopleURL FROM movie")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.Movie

		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Type, &movie.Country, &movie.Language, &movie.Length, &movie.IMDb, &movie.StarNum, &movie.Score, &movie.OneStar, &movie.TwoStar, &movie.ThreeStar, &movie.FourStar, &movie.FiveStar, &movie.HaveWatched, &movie.WantToWatch, &movie.Synopsis, &movie.URL, &movie.PeopleURL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
