
import React from 'react';
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import {ip} from './global.js'
import './index.css';
import "./watch.css";
import { Link } from 'react-router-dom';


export const Watched = () => {
  const [WatchedMovieData, setWatchedMovieData] = useState([]);
  

const detailURL ="http://"+ip+":8080/getUserProfile/Srikant";
useEffect(() => {
  fetchData();
}, []);

const fetchData = () => {
  axios
    .get(detailURL)
    .then((res) => {
      //console.log(res);

      let temp ;
     console.log(res['data']['Watched'].length);
      
      for(var key in res['data']['Watched']){
        console.log(key);
        console.log( res['data']['Watched'][key]);
        let movieData = res['data']['Watched'][key];
        let temp = <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
        title={movieData.Original_title}
        rating={movieData.Vote_average}
        identifier = {movieData.Id}
        />
        const struct = <div>{temp}</div>

         setWatchedMovieData(WatchedMovieData=>[...WatchedMovieData,
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
      <h2>The movies you have watched and rated..</h2>
      <br/>
      <br/>
      <div className='list'>
     
      {WatchedMovieData}
      
      </div>
      </div>
    </div>
  );
};



// import React, { useContext } from "react";

// import "./watch.css";
// import { GlobalContext } from "../context/GlobalState";
// import { MovieCard } from "./MovieCard";
// import { Link } from 'react-router-dom';



// export const Watched = () => {
//   const { watched } = useContext(GlobalContext);

//   return (
  //   <div className="moviespage">
  //     <section id="header">
  //      <div className="header container">
  //   <div className="nav-bar">
  //     <div className="brand">
  //       <a href="#hero">
  //       <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
  //       </a>
  //     </div>
  //     <div className="nav-list">
  //         <div className="hamburger">
  //           <div className="bar"></div>
  //         </div>
  //         <ul>
  //           <li><Link to='/UserProfile/watchlist'>watchlist</Link></li>
  //           <li><Link to='/UserProfile/add'>Add</Link></li>
  //         </ul>
  //       </div>
  //   </div>
  // </div>
  // </section>

//     <div className="services container2">
//       <div className="service-top">
//       <h1 className="section-title"><span>W</span>atched</h1>
//         <p>Add your favorite movies to Watched and Watchlist Section!</p>
//       </div>
//       <div className="service-bottom">
//       </div>
//     </div>
  
//       <div className="container">
//         <div className="header">
//           <h1 className="heading">Watched Movies</h1>

//           <span className="count-pill">
//             {watched.length} {watched.length === 1 ? "Movie" : "Movies"}
//           </span>
//         </div>

//         {watched.length > 0 ? (
//           <div className="movie-grid">
//             { watched.map((movie) => (
//               <MovieCard movie={movie} key={movie.Id} type="watched" />
//             ))}
//           </div>
//         ) : (
//           <h2 className="no-movies">Add watched movies here!</h2>
//         )}
//       </div>
//     </div>
//   );
// };