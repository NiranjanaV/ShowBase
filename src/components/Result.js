import React, { useContext } from "react";
import { GlobalContext } from "../context/GlobalState";
import Moment from "react-moment";
import "./watch.css";
import MovieSection from "./MovieSection";
import MovieDetail from "./MovieDetail";
import star from './star.png'
import { Link } from 'react-router-dom'

function Result({props}){

  return (

    <Link to='/MovieDetail' state={{ from: props.identifier }}>
    <div id={props.identifier} className='resultscard'>
        <div className="moviecard-poster-wrapper">
            <img src={props.img}  alt='exp1' className='exp1'></img>
            </div>
            </div>
    </Link>

  

    
    

      )
}

export default Result;