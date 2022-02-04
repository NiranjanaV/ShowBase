import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import logo from './logo.png'


import './katie-zaferes.png'

import star from './star.png'


import movieData from './movie.js'



function Nav(){
return(
  <div>
    <nav>
      {/* <img src={logo} id='logo' alt='logo'></img> */}
      <span>ShowBiz</span>
    </nav>
    
  </div>
)
}

/*
function ExperienceSection(props){
  return(
    <section class='reusable'>
    <div className='image'><h4 className='imageText'>Sold out</h4><img src={props.img} alt='exp1' className='exp1'></img></div>
    <div className='rating'><img src={star} alt='star' className='star'></img><span>{props.rating}<span className='grey'> ({props.reviewCount}).{props.country}</span></span></div>
    <p>{props.title}</p>
   <p>From ${props.price}/per person</p>
    </section>
  )
}
*/
const moviemap = movieData.map(
  movieData=>
  <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
  title={movieData.Original_title}
  rating={movieData.Vote_average}
   />
  )

  console.log(moviemap);

function MovieSection(props){
  return(
    <section class='reusable'>
    <div className='image'><img src={props.img} alt='exp1' className='exp1'></img></div>
    <div className='rating'><img src={star} alt='star' className='star'></img><span>{props.rating}</span></div>
    <p>{props.title}</p>
    </section>
  )
}



ReactDOM.render(
  <React.StrictMode>
    <Nav />
  
    <div class='list'>
    {moviemap}
    </div>
    <div class='list'>
    {moviemap}
    </div>
    <div class='list'>
    {moviemap}
    </div>

  </React.StrictMode>,
  document.getElementById('root')
);
