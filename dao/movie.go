package dao

import "douban/model"

// SelectMovieById 通过id来搜索电影
func SelectMovieById(movieId int) (model.Movie, error) {
	var movie model.Movie

	row := dB.QueryRow("SELECT id, name, year, director, screenwriter, starring, type, country, language, releasedate, length, aka, imdb, starnum, score, onestar, twostar, threestar, fourstar, fivestar, havewatched, wanttowatch, synopsis FROM movie WHERE id = ? ", movieId)
	if row.Err() != nil {
		return movie, row.Err()
	}

	err := row.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Type, &movie.Country, &movie.Language, &movie.ReleaseDate, &movie.Length, &movie.Aka, &movie.IMDb, &movie.StarNum, &movie.Score, &movie.OneStar, &movie.TwoStar, &movie.ThreeStar, &movie.FourStar, &movie.FiveStar, &movie.HaveWatched, &movie.WantToWatch, &movie.Synopsis)
	if err != nil {
		return movie, err
	}

	return movie, nil
}
