import React, { useContext } from "react";
import { GlobalContext } from "../context/GlobalState";
import Moment from "react-moment";
import "./watch.css";

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
    <div className="resultscard">
      <div className="moviecard-poster-wrapper">
        {movie.Poster_path ? (
          <img src={`https://image.tmdb.org/t/p/original${movie.Poster_path}`} alt='exp1' className='exp1'>
          </img>
        ) : (
          <div className="poster" />
        )}
      </div>

      <div className="movieinfo">
        <div className="header">
          <h3 className="movietitle">{movie.Original_title}</h3>
          <h4 className="movierelease">
            <Moment format="YYYY">{movie.Release_date}</Moment>
          </h4>
        </div>

        <div className="add">
          <button
            className="button"
            disabled={watchlistDisabled}
            onClick={() => addMovieToWatchlist(movie)}
          >
            Add to Watchlist
          </button>

          <button
            className="button"
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