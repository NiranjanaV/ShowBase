import React, { useState } from "react";
import  './login.css';
import axios from "axios"
import "./watch.css";
import {ip} from './global.js'
import { Link } from 'react-router-dom';
import {FriendsWatched} from './FriendsWatched';
import {FriendsWatchlist} from './FriendsWatchlist';
import useAuth from "../hooks/useAuth";



function DisplayFriends() {
  const [results, setResults] = useState([]);
   const [OpenFriendsWatched, setOpenFriendsWatched] = useState(false);
   const [OpenFriendsWatchlist, setOpenFriendsWatchlist] = useState(false);
   const [Profilewl, setProfilewl] = useState(null);
   const [Profilew, setProfilew] = useState(null);

   const { auth } = useAuth();

    React.useEffect(() => {
      axios.get("http://"+ip+":8080/getFriends/"+ auth.user)
          .then((response) => {
            setResults(response.data.friends);
          });
  }, []);



  return (
    
    <div className="add-page">
          <div className="services container2">
      <div className="service-top">
      <h1 className="section-title"><span>M</span>y <span>F</span>riends</h1>
      </div>
      <div className="service-bottom">
      </div>
    </div>
      <section id="header">
       <div className="header container">
    <div className="nav-bar">
      <div className="brand">
        <a href="#hero">
        <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
        </a>
      </div>
      {/* <div className="nav-list">
          <div className="hamburger">
            <div className="bar"></div>
          </div>
          <ul>
            <li><Link to='/UserProfile/watched'>watched</Link></li>
            <li><Link to='/UserProfile/watchlist'>watchlist</Link></li>
          </ul>
        </div>
    </div> */}
    </div>
  </div>
  </section>

         

       <div className="container">
         <div className="add-content">
          <div className="input-wrapper">
             
           </div>
        
           {results.length>0 && (<ul className="results">
            {results.map((profile) => (
                 <li>
                   <div className="resultscard">

                  <div className="movieinfo">
                    <div className="header">
                      <h3 className="movietitle">{profile.Name}</h3>
                    </div>
                      
                    <div className="add">
                    <button
                      className="button"
                      // onClick={() => <FriendsProfile profile = {profile}/>} 
                      onClick = { () => {
                        setOpenFriendsWatchlist(true)
                        setProfilewl(profile)
                      }}

                    >

                      {/* state onlick=true default=false */}
                     {/* {profile.Name}'s Watchlist */}
                    <Link to="/UserProfile/profile/FriendsWatchlist" state={{ fromwl: profile }}> {profile.Name}'s Watchlist </Link>

                    </button>

                    <button
                      className="button"
                      // onClick={() => <FriendsProfile profile = {profile}/>} 
                      onClick = { () => {
                        setOpenFriendsWatched(true)
                        setProfilew(Profilew)
                      }}

                    >

                      {/* state onlick=true default=false */}
                      {/* {profile.Name}'s Watched */}
                      <Link to="/UserProfile/profile/FriendsWatched" state={{ fromw: profile }}> {profile.Name}'s Watched </Link>
                    {/* <Link to='/UserProfile/watched'> View Profile</Link> */}

                    </button>
                      
                   </div>

                  </div>
                </div>
                     
                </li>
              ))}
            </ul>
            )}
      {OpenFriendsWatched? <FriendsWatched profile = {Profilew}/>: null}
      {OpenFriendsWatchlist? <FriendsWatchlist profile = {Profilewl}/>: null}
         </div>
       </div>
</div>
    
 );
}
 

export default DisplayFriends;