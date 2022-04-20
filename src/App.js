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
import DisplayFriends from './components/displayFriends';
import {FriendsWatched} from './components/FriendsWatched';
import {FriendsWatchlist} from './components/FriendsWatchlist';
import { AuthProvider } from './context/AuthProvider';
import RequireAuth from './components/RequireAuth';
import UserHome from './components/UserHome'


function App() {
  return (
    <AuthProvider>
    <GlobalProvider> 

    
    <div className="App">
      
      <Routes>

      <Route path="/" element={<Home/>} />
      <Route element={<RequireAuth />}>
      <Route path="/MovieProfile" element={<MovieProfile/>} />
      <Route path="/watchedlist" element={<UserWatchedList/>} />
      <Route path="/MovieDetail" element={<MovieDetail/>} />
      <Route path="/TopMovieGenre" element={<TopMovieGenre/>} />
      <Route path="/UserProfile/add" element={<Add/>} />
      <Route path="/UserProfile/add" element={<Add/>} />
      <Route path="/UserProfile/watched" element={<Watched/>} />
      <Route path="/UserProfile/watchlist" element={<Watchlist/>} />
      <Route path="/UserProfile/profile" element={<ProfilePage/>} />
      <Route path="/UserProfile/profile/searchUser" element={<SearchProfile/>} />
      <Route path="/UserProfile/profile/displayFriends" element={<DisplayFriends/>} />
      <Route path="/UserProfile/profile/FriendsWatched" element={<FriendsWatched/>} />
      <Route path="/UserProfile/profile/FriendsWatchlist" element={<FriendsWatchlist/>} />
      <Route path="/UserProfile" element={<UserProfile/>} />
      <Route path="/UserHome" element={<UserHome/>} />
      </Route>
      
     
          
      
       
     
      
     
      <Route path="/login" render={(routeProps) => <Login {...routeProps}/>} element={<Login/>} />
      <Route path="/signup" element={<SignUp/>} />
     
      
     
      </Routes>
    </div>
    </GlobalProvider>
    </AuthProvider>
      
  );
}

export default App;