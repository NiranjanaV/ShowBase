export default (state, action) => {
  switch (action.type) {
    case "add_watchlist":
      return {
        ...state,
        watchlist: [action.payload, ...state.watchlist],
      };
    case "remove_watchlist":
      return {
        ...state,
        watchlist: state.watchlist.filter(
          (movie) => movie.Id !== action.payload
        ),
      };
    case "add_watched":
      return {
        ...state,
        watchlist: state.watchlist.filter(
          (movie) => movie.Id !== action.payload.id
        ),
        watched: [action.payload, ...state.watched],
      };
    case "move_watchlist":
      return {
        ...state,
        watched: state.watched.filter(
          (movie) => movie.Id !== action.payload.id
        ),
        watchlist: [action.payload, ...state.watchlist],
      };
    case "remove_watched":
      return {
        ...state,
        watched: state.watched.filter((movie) => movie.Id !== action.payload),
      };
    default:
      return state;
  }
};
