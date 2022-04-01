import React, { useContext } from "react";
import "./watch.css";
// import movieData from './movie.js'
import { GlobalContext } from "../context/GlobalState";
import { MovieCard } from "./MovieCard";
import { Link } from 'react-router-dom';

export const Watchlist = () => {
  const { watchlist } = useContext(GlobalContext);

  return (
    <div className="movie-page">
      <section id="header">
       <div className="header container">
    <div className="nav-bar">
      <div className="brand">
        <a href="#hero">
        <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
        </a>
      </div>
      <div className="nav-list">
          <div className="hamburger">
            <div className="bar"></div>
          </div>
          <ul>
            <li><Link to='/UserProfile/watched'>watched</Link></li>
            <li><Link to='/UserProfile/add'>Add</Link></li>
          </ul>
        </div>
    </div>
  </div>
  </section>

    <div className="services container2">
      <div className="service-top">
      <h1 className="section-title"><span>W</span>atchlist</h1>
        <p>Add your favorite movies to Watched and Watchlist Section!</p>
      </div>
      <div className="service-bottom">
      </div>
    </div>
  

      <div className="container">
        <div className="header">
          <h1 className="heading">My Watchlist</h1>

          <span className="count-pill">
            {watchlist.length} {watchlist.length === 1 ? "Movie" : "Movies"}
          </span>
        </div>

        {watchlist.length > 0 ? (
          <div className="movie-grid">
            {watchlist.map((movie) => (
              <MovieCard movie={movie} key={movie.Id} type="watchlist" />
            ))}
          </div>
        ) : (
          <h2 className="no-movies">No movies in your list! Add some!</h2>
        )}
      </div>
    </div>
  );
};
