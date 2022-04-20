import React from 'react';
import axios from "axios"
import { useLocation } from 'react-router-dom'
import { useEffect, useState } from "react"
import  './movieDetail.css';
import {ip} from './global.js'

import { Link } from 'react-router-dom'
import Rating from './Rating';
import UserWatched from './user-watched';
import ToWatch from './user-to-watch';
import useAuth from "../hooks/useAuth";

function MovieDetail(props){


    const [moviedetails, setMovieData] = useState([]);
    const [watchedcolor, setWatchedColor] = useState([]);
    const [towatch, setToWatchColor] = useState([]);
    const [prev, setRate] = useState([]);

    var similarMovieName='';
   

     const location = useLocation()
     const { from } = location.state
     console.log(from);

     const {auth} = useAuth();
     const user = auth.user;

     let val;
   const detailURL ="http://"+ip+":8080/getMovieOfUser/"+from+"/"+user;
   console.log(detailURL);



        useEffect(() => {
            async function fetchData() {
              // You can await here
              const resp = await axios.get(detailURL) 
              console.log("hi");
              console.log(resp);
              return resp;
              /*
              resp.data.movie.url = "https://image.tmdb.org/t/p/original" + resp.data.movie.Poster_path;
              var genremap = await axios.get(detailURL)
                genremap = resp.data.movie.Genres.map(genre=>(<h2>{genre.Name}</h2>))
               resp.data.movie.genremap=genremap
         setMovieData(resp.data.movie)
       
*/
              
       
            
    
              // ...
            }
            fetchData().then( (response) => setResponse(response)
          
            
            );
          }, [detailURL]);

       function setResponse(resp){
         
        resp.data.movie.url = "https://image.tmdb.org/t/p/original" + resp.data.movie.Poster_path;
        
         var genremap = resp.data.movie.Genres.map(genre=>(<Link to='/TopMovieGenre' state={{ from: genre.Name }}><button>{genre.Name}</button></Link>))
         resp.data.movie.genremap=genremap;

         for( var genre in resp.data.movie.Genres) console.log(genre.Name)
        similarMovieName = resp.data.movie.Genres.map(genre=><div><h2>{genre.Name}</h2><h2>{resp.data[genre.Name].Results}</h2></div>);
         console.log(similarMovieName);

         //for(var i=0;i<similarMovieName.length;i++) {console.log(similarMovieName[i]);console.log(similarMovieName[i].Results)};

         
        //var similarMovies = resp.data.Genres.map(genre=>resp.data.genre.Name);
        //console.log(similarMovies);
         setMovieData(resp.data.movie)
   setRate(resp.data.user.Like);
   console.log(resp.data);
   console.log(resp.data.movie);
   console.log(resp.data.user.Watched);
   console.log(resp.data.user.Like)
   if(resp.data.user.Watched==1){
     setWatchedColor( "#ceead6");
     console.log("inside if");
   }else{
     setWatchedColor("white");
   }
   if(resp.data.user.ToWatch == 1){
    setToWatchColor( "#ceead6");
  }else{
    setToWatchColor("white");
  }
                

       };
          

    return(

    <div className="bg">
    <div className="detail">
    <img src={moviedetails.url} alt='image_poster' className='expdetail'></img>
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

<div className="col">
    <div className="row">
  <h1>{moviedetails.Title}</h1>

 
  </div>
  <h4>{moviedetails.Overview}</h4>
<div class='genre'>{moviedetails.genremap}</div>
<div className="caption">ShowBase Rating : <span className='bold'>{moviedetails.Vote_average}</span><span className='grey'>/10</span></div>
<Rating movie={from} prev={prev}/>
<UserWatched movie={from} Wcolor={watchedcolor}/>
<ToWatch movie={from} Tcolor={towatch} />
  {similarMovieName}
 
  </div>



   
    </div>
  
</div>
    )
}

export default MovieDetail;