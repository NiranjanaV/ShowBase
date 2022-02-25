import React from 'react';
import './index.css';

import star from './star.png'


  function MovieSection(props){
  return(
    
    <section class='reusable'>
    <div className='image'><img src={props.img} alt='exp1' className='exp1'></img></div>
    <div className='rating'><img src={star} alt='star' className='star'></img><span>{props.rating}</span></div>
    <p>{props.title}</p>
    </section>
  )
}


export default MovieSection;