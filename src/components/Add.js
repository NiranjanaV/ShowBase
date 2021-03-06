import React, { useState } from "react";
import { ResultCard } from "./ResultCard";
import axios from 'axios';
import "./styles.css";
import { Link } from 'react-router-dom';
import {ip} from './global.js';
import Result from './Result';
import MovieSection from "./MovieSection";
import MovieAdd from "./MovieAdd";



export const Add = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);

  const onChange = (e) => {
    e.preventDefault();

    setQuery(e.target.value);
   

     axios.get("http://"+ip+":8080/search/"+ e.target.value)
    .then((response) => {
      
      setResults(response.data.Movies.Results);
    });
  };

  return (
    
    <div className="add-page">
          <div className="services container2">
      <div className="service-top">
      <h1 className="section-title"><span>A</span>dd</h1>
        <p>Add your favorite movies from the Search bar!</p>
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
      <div className="nav-list">
          <div className="hamburger">
            <div className="bar"></div>
          </div>
          <ul>
            <li><Link to='/UserProfile/watched'>watched</Link></li>
            <li><Link to='/UserProfile/watchlist'>watchlist</Link></li>
          </ul>
        </div>
    </div>
  </div>
  </section>

      <div className="container">
        <div className="add-content">
          <div className="input-wrapper">
            <input
              type="text"
              placeholder="Search for a movie"
              value={query}
              onChange={onChange}
            />
          </div>
        
            <ul className="results">
            {results.map((movie) => (
                <li key={movie.id}>
                  <MovieAdd movie={movie}
        />
                </li>
              ))}
            </ul>
      
        </div>
      </div>
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
  </section> */}
    </div>
  );
};