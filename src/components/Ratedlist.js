import React from 'react';
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import {ip} from './global.js'
import './index.css';
import { Link } from 'react-router-dom';
import useAuth from "../hooks/useAuth";



export const Ratedlist = () => {
  const [LikeMovieData, setLikeMovieData] = useState([]);
  
  const { auth } = useAuth();
const detailURL ="http://"+ip+":8080/getUserProfile/"+auth.user;
useEffect(() => {
  fetchData();
}, []);

const fetchData = () => {
  axios
    .get(detailURL)
    .then((res) => {
      

      let temp ;
   console.log(res['data']);
     console.log(res['data']['Like']);
      
      for(var key in res['data']['Like']){
       // console.log(key);
        let movieData = res['data']['Like'][key]['UserLikeMovie'];
       //console.log(movieData);
       //console.log(`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`);
        if(  movieData.Poster_path!=="" ){
  
        let temp = <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
        title={movieData.Original_title}
        rating={movieData.Vote_average}
        identifier = {movieData.Id}
        />
    

        const struct = <div>{temp}</div>

         setLikeMovieData(LikeMovieData=>[...LikeMovieData,
             struct
         ])
        }
        
    
    }
   
    
});
}
return (
    <div className="moviespage">
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
  <div class='section-content'>
      <h2>The movies that are on your watch list..</h2>
      <br/>
      <br/>
      <div className='list'>
     
      {LikeMovieData}
      
      </div>
      </div>
    </div>
  );
};


export default Ratedlist;