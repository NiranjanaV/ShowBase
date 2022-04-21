import React from 'react';
import './index.css';

import star from './star.png'
import { Link } from 'react-router-dom'
import Moment from "react-moment";

  function MovieAdd(movie){
    return(
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
            >
                <Link to='/MovieDetail' state={{ from: movie.Id }}>
              Movie Detail
              </Link>
            </button>
  
          </div>
        </div>
      </div>
      )
  
}


export default MovieAdd;