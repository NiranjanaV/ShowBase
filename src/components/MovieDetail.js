import React from 'react';
import { Link } from 'react-router-dom';
import axios from "axios"
import { useLocation } from 'react-router-dom'
import { useEffect, useState } from "react"
import  './movieDetail.css';

function MovieDetail(props){


    const [moviedetails, setMovieData] = useState([]);
    
   

     const location = useLocation()
     const { from } = location.state
     console.log(from);

  const detailURL ="http://192.168.0.206:8080/getMovie/"+from;
   console.log(detailURL);



        useEffect(() => {
            async function fetchData() {
              // You can await here
              const resp = await axios.get(detailURL)
              resp.data.movie.url = "https://image.tmdb.org/t/p/original" + resp.data.movie.Poster_path;
               
        setMovieData(resp.data.movie)
        console.log(resp.data);
       
        
    
              // ...
            }
            fetchData();
          }, [detailURL]);

         

    return(

    <div className="bg">
    <div className="detail">
    <img src={moviedetails.url} className='exp1'></img>


<div className="col">
  <h1>{moviedetails.Title}</h1><br/>
  <h4>{moviedetails.Overview}</h4>
 
  </div>


    
   
    </div>
</div>
    )
}

export default MovieDetail;