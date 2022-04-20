import React from "react";
import "./styles.css";
import AddmovieIcon from "./addmovie.png"
import WatchedIcon from "./watching.jpg"
import WatchlistIcon from "./Watchlist.png"
import useAuth from "../hooks/useAuth";

import { Link } from 'react-router-dom'
function UserProfile() {
  const { auth } = useAuth();
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
            <li><Link to='/UserProfile/watched'>watched</Link></li>
            <li><Link to='/UserProfile/watchlist'>watchlist</Link></li>
            <li><Link to='/UserProfile/add'>Add</Link></li>
            <li><Link to='/UserProfile/profile'>Profile</Link></li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="services">
    <div className="services container">
      <div className="service-top">
      <h1 className="section-title"><span>{auth.user.charAt(0)}</span>{auth.user.substring(1,auth.user.length)}'s <span>P</span>rofile</h1>
        <p>Add your favorite movies to Watched and Watchlist Section!</p>
      </div>
      <div className="service-bottom">
        <div className="service-item">
          <div className="icon"><Link to='/UserProfile/watched'><img src={WatchedIcon} /></Link></div>
          <h2><Link to='/UserProfile/watched'>watched</Link></h2>
          <p> Track the movies that you have completed</p>
        </div>
        <div className="service-item">
          <div className="icon"><Link to='/UserProfile/Watchlist'><img src={WatchlistIcon} /></Link></div>
          <h2><Link to='/UserProfile/Watchlist'>Watchlist</Link></h2>
          <p> The movies that you want to watch</p>
        </div>
        <div className="service-item">
          <div className="icon"><Link to='/UserProfile/Add'><img src={AddmovieIcon} /></Link></div>
          <h2><Link to='/UserProfile/Add'>Add</Link></h2>
          <p> Add movies to your Watchlist and Watched</p>
        </div>
      </div>
    </div>
  </section>

  <section id="footer">
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
  </section>
    </>
  );
}

export default UserProfile;