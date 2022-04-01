import React, { useContext } from "react";
import "./watch.css";
import movieData from './movie.js'
import { GlobalContext } from "../context/GlobalState";
import { MovieCard } from "./MovieCard";
import { Link } from 'react-router-dom';


export const Watched = () => {
  const { watched } = useContext(GlobalContext);

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
            <li><Link to='/UserProfile/watchlist'>watchlist</Link></li>
            <li><Link to='/UserProfile/add'>Add</Link></li>
          </ul>
        </div>
    </div>
  </div>
  </section>

    <div className="services container2">
      <div className="service-top">
      <h1 className="section-title"><span>W</span>atched</h1>
        <p>Add your favorite movies to Watched and Watchlist Section!</p>
      </div>
      <div className="service-bottom">
      </div>
    </div>
  
      <div className="container">
        <div className="header">
          <h1 className="heading">Watched Movies</h1>

          <span className="count-pill">
            {watched.length} {watched.length === 1 ? "Movie" : "Movies"}
          </span>
        </div>

        {watched.length > 0 ? (
          <div className="movie-grid">
            { watched.map((movie) => (
              <MovieCard movie={movie} key={movie.Id} type="watched" />
            ))}
          </div>
        ) : (
          <h2 className="no-movies">No movies in your list! Add some!</h2>
        )}
      </div>
    </div>
  );
};
