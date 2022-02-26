import MovieSection from "./MovieSection";
import movieData from './movie.js'
import movie22 from './movie22.js'
import movie21 from './movie21.js'
import movie20 from './movie20.js'




const moviemap = movieData.map(
    (movieData)=>(
    movieData.Poster_path!=="" ?
    <MovieSection img={`https://image.tmdb.org/t/p/original/${movieData.Poster_path}`}
    title={movieData.Original_title}
    rating={movieData.Vote_average}
     />
     :console.log("what to do now")
    ))

    const moviemap22 =  movie22.map(
      (movie22)=> (
      movie22.Poster_path!=="" ?
      <MovieSection img={`https://image.tmdb.org/t/p/original/${movie22.Poster_path}`}
      title={movie22.Original_title}
      rating={movie22.Vote_average}
       />
       :console.log("what to do now")
      ))
    
  
      const moviemap21 = movie21.map(
        (movie21)=>(
          movie21.Poster_path!=="" ?
        <MovieSection img={`https://image.tmdb.org/t/p/original/${movie21.Poster_path}`}
        title={movie21.Original_title}
        rating={movie21.Vote_average}
         />
         :console.log("what to do now")
        ))
  
        const moviemap20 = movie20.map(
          (movie20)=>(
            movie20.Poster_path!=="" ?
          <MovieSection img={`https://image.tmdb.org/t/p/original/${movie20.Poster_path}`}
          title={movie20.Original_title}
          rating={movie20.Vote_average}
           />
           :console.log("what to do now")
          ))
  
          
  
  

function MovieProfile(){
    return(

        <div>
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
   
        </div>
    )
}

export default MovieProfile;
