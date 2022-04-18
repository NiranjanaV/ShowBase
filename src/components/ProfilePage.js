import React from "react";
import "./styles.css";



import { Link } from 'react-router-dom'
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
            <li><Link to='/UserProfile/profile/displayFriends'>My Friends</Link></li>
            <li><Link to='/UserProfile/profile/searchUser'>Add new Friends</Link></li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="services">
    <div className="services container">
      <div className="service-top">
      <h1 className="section-title"><span>P</span>rofile</h1>
        <p>Welcome to your Profile Page!</p>
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