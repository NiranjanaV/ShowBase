import React, { createContext, useReducer, useEffect } from "react";
import AppReducer from "./AppReducer";

const initialState = {
  watchlist: localStorage.getItem("watchlist")
    ? JSON.parse(localStorage.getItem("watchlist"))
    : [],
  watched: localStorage.getItem("watched")
    ? JSON.parse(localStorage.getItem("watched"))
    : [],
};

export const GlobalContext = createContext(initialState);

export const GlobalProvider = (props) => {
  const [state, dispatch] = useReducer(AppReducer, initialState);

  useEffect(() => {
    localStorage.setItem("watchlist", JSON.stringify(state.watchlist));
    localStorage.setItem("watched", JSON.stringify(state.watched));
  }, [state]);

  const addMovieToWatchlist = (movie) => {
    dispatch({ type: "add_watchlist", payload: movie });
  };

  const removeMovieFromWatchlist = (id) => {
    dispatch({ type: "remove_watchlist", payload: id });
  };

  const addMovieToWatched = (movie) => {
    dispatch({ type: "add_watched", payload: movie });
  };

  const moveToWatchlist = (movie) => {
    dispatch({ type: "move_watchlist", payload: movie });
  };

  const removeFromWatched = (id) => {
    dispatch({ type: "remove_watched", payload: id });
  };

  return (
    <GlobalContext.Provider
      value={{
        watchlist: state.watchlist,
        watched: state.watched,
        addMovieToWatchlist,
        removeMovieFromWatchlist,
        addMovieToWatched,
        moveToWatchlist,
        removeFromWatched,
      }}
    >
      {props.children}
    </GlobalContext.Provider>
  );
};
