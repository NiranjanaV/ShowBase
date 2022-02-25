import React from "react";
import Carousel from "react-elastic-carousel";
import Item from "./Item";
import "./styles.css";
import Discover from "./discover.png"
import Share from "./share.png"
import Expand from "./expand.png"
import Track from "./track.png"
import img1 from "./img-1.jpeg"
import img2 from "./img-2.jpeg"
import img3 from "./img-3.jpeg"
import img4 from "./img-4.jpeg"
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

function Home() {
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
            <li><a href="#hero" data-after="Home">Home</a></li>
            <li><a href="#services" data-after="Service">About</a></li>

            <li><a href="#projects" data-after="Projects">Features</a></li>
            <li><Link to='/MovieProfile'>Movies & Shows</Link></li>
            <li><a href="#Completed" data-after="Completed">Completed</a></li>
            <li><a href="#Stopped" data-after="Stopped">Stopped</a></li>
            {/* <li><a href="#about" data-after="About">Profile</a></li> */}

            <li><a href="#contact" data-after="Contact">Contact</a></li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="hero">
    <div className="hero container">
      <div>
        <h1>An amazing Tracker,   <span></span></h1>
        <h1>for all your TV Shows & Movies <span></span></h1>
        <h1>at one place<span></span></h1>
        <a href="#projects" type="button" className="cta">Features</a>
      </div>
    </div>
  </section>

  <section id="services">
    <div className="services container">
      <div className="service-top">
        <h1 className="section-title">Ab<span>o</span>ut</h1>
        <p>ShowBase provides these and so much more</p>
      </div>
      <div className="service-bottom">
        <div className="service-item">
          <div className="icon"><img src={Track} /></div>
          <h2>TRACK</h2>
          <p> shows and movies you love</p>
        </div>
        <div className="service-item">
          <div className="icon"><img src={Discover} /></div>
          <h2>DISCOVER</h2>
          <p> what to watch next</p>
        </div>
        <div className="service-item">
          <div className="icon"><img src={Share} /></div>
          <h2>SHARE</h2>
          <p> it with your friends</p>
        </div>
        <div className="service-item">
          <div className="icon"><img src={Expand} /></div>
          <h2>EXPAND</h2>
          <p>your horizon of interests</p>
        </div>
      </div>
    </div>
  </section>

  <section id="projects">
    <div className="projects container">
      <div className="projects-header">
        <h1 className="section-title">Fea<span>tures</span></h1>
      </div>
      <div className="all-projects">
        <div className="project-item">
          <div className="project-info">
            <h1>Personal Account</h1>
            <p>Have your own personal account that keeps track of all the TV shows you're currently watching, that you want to watch and those you have already completed! You can also add favorites so you can later rewatch those!! </p>
          </div>
          <div className="project-img">
            <img src={img1} alt="img"/>
          </div>
        </div>
        <div className="project-item">
          <div className="project-info">
            <h1>Be up-to-date</h1>
            <p>You can send friend requests and accept friends rquests. Now you'll be able to see all your friends latest interests in TV shows and be up-to-date with the latest trends!</p>
          </div>
          <div className="project-img">
            <img src={img2} alt="img"/>
          </div>
        </div>
        <div className="project-item">
          <div className="project-info">
            <h1>TV Show Match</h1>
            <p>When a group of your friends are confused on what movie to watch on friday night, 
              don’t fret! Sync up your devices and start the movie match feature. 
              Choose the genre and right swipe/left swipe movies based on your preference. 
              The movie with the most right swipes wins that night!!</p>
          </div>
          <div className="project-img">
            <img src={img3} alt="img"/>
          </div>
        </div>
        <div className="project-item">
          <div className="project-info">
            <h1>Review directly</h1>
            <p>Like a tv show so much? that you're bursting to share it with someone or was the tv show not up to your standards?Either way, you can review the shows directly from the website which is linked to google reviews!
            Now others can also make an informed decision.</p>
          </div>
          <div className="project-img">
            <img src={img4} alt="img"/>
          </div>
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
            <h2>k.iyer@ufl.edu</h2>
            
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

export default Home;