
import React from 'react';
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import {ip} from './global.js'




const MovieProfile = () => {
  const [HomeMovieData, setHomeMovieData] = useState([]);
  

const detailURL ="http://"+ip+":8080/top/";
useEffect(() => {
  fetchData();
}, []);

const fetchData = () => {
  axios
    .get(detailURL)
    .then((res) => {
      console.log(res);

      let temp ;
      
      for(var key in res['data']){
        console.log(key);

        console.log(res['data'][key]);

      


        temp = res.data[key].Results.map(
        (movieData)=>(
        movieData.Poster_path!=="" ?
        <div>
        
        <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
        title={movieData.Original_title}
        rating={movieData.Vote_average}
        identifier = {movieData.Id}
        />
      
        </div>
        :console.log("")
        ));

        let year = key + " Movies";


        console.log(year);

        const struct = <div><h2>{year}</h2><div className='list'>{temp}</div></div>

        setHomeMovieData(HomeMovieData=>[...HomeMovieData,
            struct
        ])
        


        console.log(temp);
    
    }
    
    
});
}
return (
    <div>
      
      <div>
     
      {HomeMovieData}
      
      </div>
    </div>
  );
};
export default MovieProfile;
/*
import MovieSection from "./MovieSection";
import movieData from './movie.js'
import movie22 from './movie22.js'
import movie21 from './movie21.js'
import movie20 from './movie20.js'
import { Link } from 'react-router-dom'
import "./styles.css";






const moviemap = movieData.map(
    (movieData)=>(
    movieData.Poster_path!=="" ?
    <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
    title={movieData.Original_title}
    rating={movieData.Vote_average}
    identifier = {movieData.Id}
     />
     :console.log("")
    ))

    const moviemap22 =  movie22.map(
      (movie22)=> (
      movie22.Poster_path!=="" ?
      <MovieSection img={`https://image.tmdb.org/t/p/original/${movie22.Poster_path}`}
      title={movie22.Original_title}
      rating={movie22.Vote_average}
      identifier = {movie22.Id}
       />
       :console.log("")
      ))
    
  
      const moviemap21 = movie21.map(
        (movie21)=>(
          movie21.Poster_path!=="" ?
        <MovieSection img={`https://image.tmdb.org/t/p/original/${movie21.Poster_path}`}
        title={movie21.Original_title}
        rating={movie21.Vote_average}
        identifier = {movie21.Id}
         />
         :console.log("")
        ))
  
        const moviemap20 = movie20.map(
          (movie20)=>(
            movie20.Poster_path!=="" ?
          <MovieSection img={`https://image.tmdb.org/t/p/original/${movie20.Poster_path}`}
          title={movie20.Original_title}
          rating={movie20.Vote_average}
          identifier = {movie20.Id}
           />
           :console.log("")
          ))
  
          //console.log(moviemap);
  
  

function MovieProfile(){
    return(
      
        <div>

          <section id="header">
       <div className="header container">
    <div className="nav-bar">
      <div className="brand">
        <a href="#hero">
        <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
        </a>
      </div>
    </div>
  </div>
  </section>


            <div> <h2 className='head'>2022 Movies</h2> </div>
 <div class='list'>
 {moviemap22}
    </div>
    <div> <h2 className='head'>2021 Movies</h2> </div>
    <div class='list'>
    {moviemap21}
    </div>

    <div> <h2 className='head'>2020 Movies</h2> </div>
    <div class='list'>
    {moviemap20}
    </div>

   <div> <h2 className='head'>2017 Movies</h2> </div>
    <div class='list'>
    {moviemap}
    </div>
   
        </div>
    )
}

export default MovieProfile;

*/