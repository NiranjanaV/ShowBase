import React from 'react';
import MovieProfile from './components/MovieProfile';
import Home from './components/Home';
import Login from './components/Login';
import SignUp from './components/SignUp';
import UserProfile from './components/UserProfile';
import { Add } from './components/Add';
import { Watchlist } from './components/Watchlist';
import { Watched } from './components/Watched';
import {Route,Routes} from 'react-router-dom';
import MovieDetail from './components/MovieDetail';
import ProfilePage from './components/ProfilePage';
import SearchProfile from './components/SearchProfile';
import TopMovieGenre from './components/TopMovieGenre';
import {GlobalProvider} from './context/GlobalState';
import UserWatchedList from './components/user-watch-list';
import DisplayFriends from './components/displayFriends'


function App() {
  return (
    <GlobalProvider> 

    
    <div className="App">
      <Routes>

      <Route path="/" element={<Home/>} />
      <Route path="/MovieProfile" element={<MovieProfile/>} />
      <Route path="/MovieDetail" element={<MovieDetail/>} />
      <Route path="/TopMovieGenre" element={<TopMovieGenre/>} />
      <Route path="/UserProfile" element={<UserProfile/>} />
      <Route path="/UserProfile/add" element={<Add/>} />
      <Route path="/UserProfile/add" element={<Add/>} />
      <Route path="/UserProfile/watched" element={<Watched/>} />
      <Route path="/UserProfile/watchlist" element={<Watchlist/>} />
      <Route path="/UserProfile/profile" element={<ProfilePage/>} />
      <Route path="/UserProfile/profile/searchUser" element={<SearchProfile/>} />
      <Route path="/UserProfile/profile/displayFriends" element={<DisplayFriends/>} />
      <Route path="/login" element={<Login/>} />
      <Route path="/signup" element={<SignUp/>} />
      <Route path="/watchedlist" element={<UserWatchedList/>} />
     
      </Routes>
    </div>
    </GlobalProvider>
      
  );
}

export default App;