import React, { useState } from "react";
import { ResultCard } from "./ResultCard";
import axios from 'axios';


export const Add = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);

  const onChange = (e) => {
    e.preventDefault();

    setQuery(e.target.value);
    console.log(e.target.value)

     axios.get("http://70.171.43.6:8080/search/"+ e.target.value)
    .then((response) => {
      console.log(response.data.Movies.Results)
      setResults(response.data.Movies.Results);
    });
  };

  return (
    <div className="add-page">
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
                  <ResultCard movie={movie} />
                </li>
              ))}
            </ul>
      
        </div>
      </div>
    </div>
  );
};