import React from 'react';


import MovieProfile from './components/MovieProfile';
import Home from './components/Home';
import Login from './components/Login'
import {Route,Routes} from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Routes>
      <Route path="/" element={<Home/>} />
      <Route path="/MovieProfile" element={<MovieProfile/>} />
      <Route path="/login" element={<Login/>} />
      </Routes>
    </div>
      
  );
}

export default App;
