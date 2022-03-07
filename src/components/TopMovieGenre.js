import React from 'react';
import { useLocation } from 'react-router-dom'
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import movieData from './movie.js'


const TopMovieGenre = () => {
  const [genreSpecificMovies, setgenreSpecificMovies] = useState([]);
  const location = useLocation()
  const { from } = location.state
  console.log(from);

const detailURL ="http://192.168.0.206:8080/getMovie/"+from;
useEffect(() => {
  fetchData();
}, []);
const fetchData = () => {
  axios
    .get(detailURL)
    .then((res) => {
      console.log(res);
      setgenreSpecificMovies(res.data[from].Results.map(
        (movieData)=>(
        movieData.Poster_path!=="" ?
        <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
        title={movieData.Original_title}
        rating={movieData.Vote_average}
        identifier = {movieData.Id}
         />
         :console.log("")
        )));
    })
    .catch((err) => {
      console.log(err);
    });
};
return (
    <div>
      
      <div> <h2 className='head'>{from}</h2> 
     <div className='list'>
      {genreSpecificMovies}
      </div>
      </div>
    </div>
  );
};
export default TopMovieGenre;