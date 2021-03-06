import React from 'react';
import axios from "axios"
import { useEffect, useState } from "react"
import MovieSection from './MovieSection';
import {ip} from './global.js'
import './index.css';
import { useLocation } from 'react-router-dom'


export const FriendsWatchlist = () => {
  const [FriendsToWatchMovieData, setFriendsToWatchMovieData] = useState([]);


  const location = useLocation()
  const name = location.state.fromwl.Name
  console.log(" aaa", name)


const detailURL ="http://"+ip+":8080/getUserProfile/" + name
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

         setFriendsToWatchMovieData(FriendsToWatchMovieData=>[...FriendsToWatchMovieData,
             struct
         ])
    }

});
}
return (
    <div>
      <h2> {name}'s Watchlist Movies & Shows </h2>
      <br/>
      <br/>
      <div className='list'>
     
      {FriendsToWatchMovieData}
      
      </div>
    </div>
  );
};




// import React, { useContext } from "react";
// import "./watch.css";
// import { GlobalContext } from "../context/GlobalState";
// import { MovieCard } from "./MovieCard";
// import { Link } from 'react-router-dom';

// export const Watchlist = () => {
//   const { watchlist } = useContext(GlobalContext);

//   return (
//     <div className="moviespage">
//       <section id="header">
//        <div className="header container">
//     <div className="nav-bar">
//       <div className="brand">
//         <a href="#hero">
//         <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
//         </a>
//       </div>
//       <div className="nav-list">
//           <div className="hamburger">
//             <div className="bar"></div>
//           </div>
//           <ul>
//             <li><Link to='/UserProfile/watched'>watched</Link></li>
//             <li><Link to='/UserProfile/add'>Add</Link></li>
//           </ul>
//         </div>
//     </div>
//   </div>
//   </section>

//     <div className="services container2">
//       <div className="service-top">
//       <h1 className="section-title"><span>W</span>atchlist</h1>
//         <p>Add your favorite movies to Watched and Watchlist Section!</p>
//       </div>
//       <div className="service-bottom">
//       </div>
//     </div>
  

//       <div className="container">
//         <div className="header">
//           <h1 className="heading">My Watchlist</h1>

//           <span className="count-pill">
//             {watchlist.length} {watchlist.length === 1 ? "Movie" : "Movies"}
//           </span>
//         </div>

//         {watchlist.length > 0 ? (
//           <div className="movie-grid">
//             {watchlist.map((movie) => (
//               <MovieCard movie={movie} key={movie.Id} type="watchlist" />
//             ))}
//           </div>
//         ) : (
//           <h2 className="no-movies">Add movies that you wanna watch!</h2>
//         )}
//       </div>
//     </div>
//   );
// };