import React from 'react';


import MovieProfile from './components/MovieProfile';
import Home from './components/Home';
import {Route,Routes} from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Routes>
      <Route path="/" element={<Home/>} />
      <Route path="/MovieProfile" element={<MovieProfile/>} />
      </Routes>
    </div>
      
  );
}

export default App;
