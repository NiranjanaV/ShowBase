import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import movieData from './movie.js'
import movie22 from './movie22.js'
import movie21 from './movie21.js'
import movie20 from './movie20.js'
import star from './star.png'


// function Nav(){
// return(
//   <div>
//     <nav>
//       {/* <img src={logo} id='logo' alt='logo'></img> */}
//       <span>ShowBiz</span>
//     </nav>
    
//   </div>
// )
// }
const moviemap = movieData.map(
  movieData=>
  <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
  title={movieData.Original_title}
  rating={movieData.Vote_average}
   />
  )
  const moviemap22 = movie22.map(
    movie22=>
    <MovieSection img={`https://image.tmdb.org/t/p/original/${movie22.Poster_path}`}
    title={movie22.Original_title}
    rating={movie22.Vote_average}
     />
    )

    const moviemap21 = movie21.map(
      movie21=>
      <MovieSection img={`https://image.tmdb.org/t/p/original/${movie21.Poster_path}`}
      title={movie21.Original_title}
      rating={movie21.Vote_average}
       />
      )

      const moviemap20 = movie20.map(
        movie20=>
        <MovieSection img={`https://image.tmdb.org/t/p/original/${movie20.Poster_path}`}
        title={movie20.Original_title}
        rating={movie20.Vote_average}
         />
        )

  // console.log(movie22.poster_path);
  // console.log(movie22);


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


ReactDOM.render(
  <React.StrictMode>
    {/* <Nav />
   */}
 <div> <h2 className='head'>2022 Movies</h2> </div>
 <div class='list'>
  {moviemap22}
    </div>
    <div> <h2 className='head'>2021 Movies</h2> </div>
    <div class='list'>
    {moviemap21}
    </div>

    <div> <h2 className='head'>2020 Movies</h2> </div>
    <div class='list'>
    {moviemap20}
    </div>

   <div> <h2 className='head'>2017 Movies</h2> </div>
    <div class='list'>
    {moviemap}
    </div>
   
  

  </React.StrictMode>,
  document.getElementById('try')
);
