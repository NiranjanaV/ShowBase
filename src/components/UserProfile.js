import React from "react";
import "./styles.css";
import AddmovieIcon from "./addmovie.png"
import WatchedIcon from "./watching.jpg"
import WatchlistIcon from "./Watchlist.png"
import useAuth from "../hooks/useAuth";
import axios from "axios"
import {ip} from './global.js'
import { useEffect, useState, useContext } from "react"
import AuthContext from "../context/AuthProvider";



import { Link } from 'react-router-dom'


function UserProfile() {
const {auth} =useAuth();
//const[log,setlog] = useContext(AuthContext);

  const [ToWatchMovieNumber, setToWatchMovieNumber] = useState([]);
  const [WatchedMovieNumber, setWatchedMovieNumber] = useState([]);
  const [MovieRatedNumber, setMovieRatedNumber] = useState([]);
 

  const detailURL ="http://"+ip+":8080/getUserProfile/"+auth.user;
  useEffect(() => {
    fetchData();
  }, []);
  
  const fetchData = () => {
    axios
      .get(detailURL)
      .then((res) => {
        console.log(res);
  
       setToWatchMovieNumber(res['data']['ToWatch'].length);
       setWatchedMovieNumber(res['data']['Watched'].length);
       setMovieRatedNumber(res['data']['Like'].length);
       
        
        });
      
  }
  

const handleLogout= async(e) => {
  
  console.log("logged out");
  //setlog({});


  
  
}





  return (
    <>

  <section id="header">
    <div className="header container">
      <div className="nav-bar">
        <div className="brand">
          <a href="#hero">
          <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
          </a>
        </div>
        <div className="nav-list">
          <div className="hamburger">
            <div className="bar"></div>
          </div>
          <ul>
            <li><Link to='/UserProfile/profile'>Profile</Link></li>
            <li><Link to='/UserProfile/add'>Add</Link></li>
            <li><button onClick={handleLogout}>Logout</button></li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="services">
    <div className="services container">
      <div className="service-top">
      <h1 className="section-title">{auth.user}'s Profile</h1>
        {/* <p>Add your favorite movies to Watched and Watchlist Section!</p> */}
      </div>
      <div className="service-bottom">

      <Link to='/UserProfile/watched'>
     
        <div className="service-item">
          {/* <div className="icon"><img src={WatchedIcon} /></div> */}
          <div className='statLabel'>{WatchedMovieNumber}</div>
          <h2>Watched</h2>
        </div>
        </Link>

        <Link to='/UserProfile/Watchlist'>
     
        <div className="service-item">
          {/* <div className="icon"><img src={WatchlistIcon} /></div> */}
          <div className='statLabel'>{ToWatchMovieNumber}</div>
          <h2>Watchlist</h2>
        </div>
        </Link>

        <Link to='/UserProfile/Ratedlist'>
     
     <div className="service-item">
       {/* <div className="icon"><FaStar/></div> */}
       <div className='statLabel'>{MovieRatedNumber}</div>
       <h2>Rated</h2>
     </div>
     </Link>


        {/* <div className="service-item">
          <div className="icon"><Link to='/UserProfile/Watchlist'><img src={WatchlistIcon} /></Link></div>
          <div className='statLabel'>{ToWatchMovieNumber}</div>
          <h2><Link to='/UserProfile/Watchlist'>Watchlist</Link></h2>
          <p> The movies that you want to watch</p>
        </div> */}
        {/* <div className="service-item">
          <div className="icon"><Link to='/UserProfile/Add'><img src={AddmovieIcon} /></Link></div>
          <h2><Link to='/UserProfile/Add'>Add</Link></h2>
          <p> Add movies to your Watchlist and Watched</p>
        </div>
      </div>
    </div> */}
    </div>
    </div>
  </section>

   {/* <section id="footer">
    <div className="footer container">
      <div className="brand">
        <h1>ShowBase</h1>
      </div>
      <h2>Software Engineering Project</h2>
      <div className="social-icon">
        <div className="social-item">
          <a href="https://github.com/NiranjanaV/ShowBase"><img src="https://img.icons8.com/bubbles/100/000000/github.png" /></a>
        </div>
      </div>
    </div>
    
  </section>  */}
    </>
  );
}

export default UserProfile;