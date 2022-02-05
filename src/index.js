import React from "react";
import ReactDOM from "react-dom";
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

const breakPoints = [
  { width: 550, itemsToShow: 1 },
  { width: 550, itemsToShow: 2 },
  { width: 750, itemsToShow: 3 },
  { width: 750, itemsToShow: 4 },
  { width: 750, itemsToShow: 5 }

];

function App() {
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

      </div>
    </div>
  </section>

 
  <section id="Completed">
    <div className="Completed container">
      
      <div className="col-right">
        <h1 className="section-title">Completed</h1>
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
  
  <section id="contact">
    <div class="contact container">
      <div>
        <h1 class="section-title">Contact</h1>
      </div>
      <div class="contact-items">
        
        <div class="contact-item">
          <div class="icon"><img src="https://img.icons8.com/bubbles/100/000000/new-post.png" /></div>
          <div class="contact-info">
            <h1>Email</h1>
            <h2>srinivasansanjay@ufl.edu</h2>
            <h2>ni.vasudevan@ufl.edu</h2>
            <h2>sw.velavan@ufl.edu</h2>
            
          </div>
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

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);