import movieData from './movie.js'
 import "./styles.css";
import _ from "lodash"
import { Link } from 'react-router-dom'
import React, { useState, useRef } from "react";



export const Add = () => {
  const [rows, setRows] = useState(movieData);
  const [searchDataset, setSearchDataset] = useState("");

  const debounceValue = useRef(
    _.debounce((searchText) => {
      requestSearch(searchText);
    }, 300)
  );

  const handleInputChange = (val) => {
    setSearchDataset(val);
    debounceValue.current(val);
  };

  const requestSearch = (searchedVal) => {
    const filteredRows = movieData.filter((row) => {
      return (
        row.Original_title.toLowerCase().includes(searchedVal.toLowerCase())
      );
    });
    setRows(filteredRows);
  };

  return (

    <div className="Add">
      <section id="header">
       <div className="header container">
    <div className="nav-bar">
      <div className="brand">
        <a href="#hero">
        <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
        </a>
      </div>
    </div>
  </div>
  </section>

      <input
      type="text" 
      placeholder="Search for a movie"
        value={searchDataset}
        onChange={(event) => handleInputChange(event.target.value)}
      />
      {rows.length > 0 && rows.map((r) => {
        return(
          <>
          <p>{r.Original_title}</p>
          <div className='image'>
            <img src={`https://image.tmdb.org/t/p/original${r.Poster_path}`} alt='exp1' className='exp1'>
              </img></div>
          </>
        )
      })
      }
    </div>
  );
}