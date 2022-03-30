import React, { useContext } from "react";
import movieData from './movie.js'
import { GlobalContext } from "../context/GlobalState";
import Moment from "react-moment";

export const ResultCard = ({ movie }) => {
  const {
    addMovieToWatchlist,
    addMovieToWatched,
    watchlist,
    watched,
  } = useContext(GlobalContext);

  let storedMovie = watchlist.find((o) => o.Id === movie.Id);
  let storedMovieWatched = watched.find((o) => o.Id === movie.Id);

  const watchlistDisabled = storedMovie
    ? true
    : storedMovieWatched
    ? true
    : false;

  const watchedDisabled = storedMovieWatched ? true : false;

  return (
    <div className="result-card">
      <div className="poster-wrapper">
        {movie.Poster_path ? (
          <img src={`https://image.tmdb.org/t/p/original${movie.Poster_path}`} alt='exp1' className='exp1'>
          </img>
        ) : (
          <div className="filler-poster" />
        )}
      </div>

      <div className="info">
        <div className="header">
          <h3 className="title">{movie.Original_title}</h3>
          <h4 className="release-date">
            <Moment format="YYYY">{movie.Release_date}</Moment>
          </h4>
        </div>

        <div className="add">
          <button
            className="click"
            disabled={watchlistDisabled}
            onClick={() => addMovieToWatchlist(movie)}
          >
            Add to Watchlist
          </button>

          <button
            className="click"
            disabled={watchedDisabled}
            onClick={() => addMovieToWatched(movie)}
          >
            Add to Watched
          </button>
        </div>
      </div>
    </div>
  );
};
