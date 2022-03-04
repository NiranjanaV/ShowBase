import React from 'react';
import './index.css';

import star from './star.png'
import { Link } from 'react-router-dom'



  function MovieSection(props){
  return(
    
    <section class='reusable'>
     
    <Link to='/MovieDetail' state={{ from: props.identifier }}>
    <div id={props.identifier} className='image'><img src={props.img}  alt='exp1' className='exp1'></img></div>
    </Link>
    <div className='rating'><img src={star} alt='star' className='star'></img><span>{props.rating}</span></div>
    <p>{props.title}</p>
    </section>
  )
}


export default MovieSection;