import React from "react";
import movieData from './movie.js'
import { MovieControls } from "./MovieControls";

export const MovieCard = ({ movie, type }) => {
  return (
    <div className="movie-card">
      <div className="overlay"></div>

      <img
        src={`https://image.tmdb.org/t/p/w200${movie.Poster_path}`}
        alt={`${movie.Original_title} Poster`}
      />

      <MovieControls type={type} movie={movie} />
    </div>
  );
};
