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
              
        setMovieData(resp.data.movie)
        console.log(resp.data);
              // ...
            }
            fetchData();
          }, [detailURL]);

         

    return(

    <div className="bg">
    <div className="detail">
 <center>


  <h2>{moviedetails.Title}</h2>
  <h4>{moviedetails.Overview}</h4>

  
    </center>

    
   
    </div>
</div>
    )
}

export default MovieDetail;