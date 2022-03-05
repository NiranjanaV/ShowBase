import React from "react";
import Carousel from "react-elastic-carousel";
import Item from "./Item";
import "./styles.css";
import witcher from "./witcher.jpg"
import sherlock from "./sherlock.jpg"
import lucifer from "./lucifer.jpg"
import manifest from "./manifest.jpg"
import peaky from "./peaky.jpg"
import riverdale from "./riverdale.jpg"
import MrRobot from "./MrRobot.jpg"
import GoT from "./GoT.jpg"
import BreakingBad from "./BreakingBad.jpg"
import money from "./money.jpg"
import dts from "./dts.jpg"

import { Link } from 'react-router-dom'


const breakPoints = [
  { width: 550, itemsToShow: 1 },
  { width: 550, itemsToShow: 2 },
  { width: 750, itemsToShow: 3 },
  { width: 750, itemsToShow: 4 },
  { width: 750, itemsToShow: 5 }

];

function UserProfile() {
  return (
    <>
	<section id="header">
    <div className="header container">
      <div className="nav-bar">
        <div className="brand">
          <a href="#hero">
            <h1><span>S</span>how <span>B</span>ase</h1>
          </a>
        </div>
        <div className="nav-list">
          <div className="hamburger">
            <div className="bar"></div>
          </div>
          <ul>
            <li><a href="#Completed" data-after="Completed">Completed</a></li>
            <li><a href="#Stopped" data-after="Stopped">Stopped</a></li>
            <li><Link to='/UserProfile/add'>Add</Link></li>
          </ul>
        </div>
      </div>
    </div>
  </section>


  
  <section id="Completed">
    <div className="Completed container">
      
      <div className="col-right">
        <h1 className="section-title"  data-testid="Completed">Completed</h1>

      <div className="App">
        <Carousel breakPoints={breakPoints}>
          <Item><img src={witcher} alt="img"/></Item>
          <Item><img src={sherlock} alt="img"/></Item>
          <Item><img src={lucifer} alt="img"/></Item>
          <Item><img src={manifest} alt="img"/></Item>
          <Item><img src={peaky} alt="img"/></Item>
          <Item><img src={riverdale} alt="img"/></Item>
        </Carousel>
      </div>
      </div>
    </div>
  </section>

  <section id="Stopped">
    <div className="Stopped container">
      
      <div className="col-right">
        <h1 className="section-title">Stopped</h1>
      <div className="App">
        <Carousel breakPoints={breakPoints}>
          <Item><img src={MrRobot} alt="img"/></Item>
          <Item><img src={GoT} alt="img"/></Item>
          <Item><img src={BreakingBad} alt="img"/></Item>
          <Item><img src={money} alt="img"/></Item>
          <Item><img src={dts} alt="img"/></Item>
        </Carousel>
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