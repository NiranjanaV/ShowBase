import React, { useContext } from "react";
import 'font-awesome/css/font-awesome.min.css';
import { GlobalContext } from "../context/GlobalState";

export const MovieControls = ({ type, movie }) => {
  const {
    removeMovieFromWatchlist,
    addMovieToWatched,
    moveToWatchlist,
    removeFromWatched,
  } = useContext(GlobalContext);

  return (
    <div className="cardcontrols">
      {type === "watchlist" && (
        <>
          <button className="click" onClick={() => addMovieToWatched(movie)}>
            <i className="fa-fw far fa-eye"></i>
          </button>

          <button
            className="click"
            onClick={() => removeMovieFromWatchlist(movie.Id)}
          >
            <i className="fa-fw fa fa-times"></i>
          </button>
        </>
      )}

      {type === "watched" && (
        <>
          <button className="click" onClick={() => moveToWatchlist(movie)}>
            <i className="fa-fw far fa-eye-slash"></i>
          </button>

          <button
            className="click"
            onClick={() => removeFromWatched(movie.Id)}
          >
            <i className="fa-fw fa fa-times"></i>
          </button>
        </>
      )}
    </div>
  );
};