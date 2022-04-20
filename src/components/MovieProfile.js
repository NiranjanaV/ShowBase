
import React from 'react';
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import {ip} from './global.js'
import { useContext } from "react";
import AuthContext from "../context/AuthProvider";
import useAuth from "../hooks/useAuth";
import { Link } from 'react-router-dom';
import './styles.css';




const MovieProfile = () => {
  const [HomeMovieData, setHomeMovieData] = useState([]);

  //const { setAuth } = useContext(AuthContext);
  const { auth } = useAuth();
 console.log("auth details");
  console.log(auth.user);
  

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
            <li><Link to='/UserProfile/add'>Add</Link></li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  {/* <div class="input-container">
  <button type="submit"><li><Link to='/UserProfile/add'>Search for Movies & Shows</Link></li> </button>
  </div> */}
        <div className="moviespage">
 
      <div>
     
      {HomeMovieData}
      
      </div>
    </div>
    </div>
  );
};

export default MovieProfile;