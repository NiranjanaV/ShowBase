import React from 'react';
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import {ip} from './global.js'
import './index.css';
import { Link } from 'react-router-dom';
import useAuth from "../hooks/useAuth";



export const Watchlist = () => {
  const [ToWatchMovieData, setToWatchMovieData] = useState([]);
  
  const { auth } = useAuth();
const detailURL ="http://"+ip+":8080/getUserProfile/"+auth.user;
useEffect(() => {
  fetchData();
}, []);

const fetchData = () => {
  axios
    .get(detailURL)
    .then((res) => {
      console.log(res);

      let temp ;
     console.log(res['data']['ToWatch'].length);
      
      for(var key in res['data']['ToWatch']){
       // console.log(key);
        console.log( res['data']['ToWatch'][key]);
        let movieData = res['data']['ToWatch'][key];
        let temp = <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
        title={movieData.Original_title}
        rating={movieData.Vote_average}
        identifier = {movieData.Id}
        />
        const struct = <div>{temp}</div>

         setToWatchMovieData(ToWatchMovieData=>[...ToWatchMovieData,
             struct
         ])

        
    
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
     
      {ToWatchMovieData}
      
      </div>
      </div>
      {/* <section id="footer">
    <div className="footer container">
      <div className="brand">
        <h1>ShowBase</h1>
      </div>
      <h2>Software Engineering Project</h2>
      <div className="social-icon">
        <div className="social-item">
          <a href="https://github.com/NiranjanaV/ShowBase"><img src="https://img.icons8.com/bubbles/100/000000/github.png" /></a>
        </div>
      </div>
    </div>
  </section> */}
    </div>
  );
};


export default Watchlist;