import React from "react";
// import Item from "./Item";
import "./styles.css";
import AddmovieIcon from "./addmovie.png"
import WatchedIcon from "./watching.jpg"
import WatchlistIcon from "./Watchlist.png"

import { Link } from 'react-router-dom'
// import { Watched } from "./Watched";
// import { Add } from "./Add";
// import { Watchlist } from "./Watchlist";

function UserProfile() {
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
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="services">
    <div className="services container">
      <div className="service-top">
      <h1 className="section-title"><span>P</span>rofile</h1>
        <p>ShowBase provides these and so much more</p>
      </div>
      <div className="service-bottom">
        <div className="service-item">
          <div className="icon"><img src={WatchedIcon} /></div>
          <h2><Link to='/UserProfile/watched'>watched</Link></h2>
          <p> shows and movies you love</p>
        </div>
        <div className="service-item">
          <div className="icon"><img src={WatchlistIcon} /></div>
          <h2>Watchlist</h2>
          <p> add movies or shows that you want to watch</p>
        </div>
        <div className="service-item">
          <div className="icon"><img src={AddmovieIcon} /></div>
          <h2>Add</h2>
          <p> it with your friends and family</p>
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