# Sprint 4

## FrontEnd

- _Profile:_ New Profile Page is created where the user can Add New Friends, view All Friends and learn about their friend's Watched and Watchlist.
  - _My Friends:_ The 'My Friends' page which will display all friends. The user can also look into their friend's Watched and Watchlist to find out what they're upto.
  - _Add New Friends:_ The user can search for their friend's profile and add them to friends list.
  - _Add:_ In the new add feature, you can search for the movie that you want to see, and when you click on the 'Movie Detail' button, it will display all the details of the movie. You can also Rate the movie here, Add it to your Watched and Watchlist.
  - _Login and Logout:_ Login has been implemented using protected routes and react context. Incorrect credentials are checked, and if the credentials are checked appropriate rerouting is done using react navigate.
  - _Movie Detail:_ Movie Detail page has added functionality such as add to watch list, mark as watched. The page also marks this with a green color if it has already been added by the user. Rating has also been added. Users may rate the movie out of 10. This state is also maintained on reload.
  - _WatchList, To Watch and Rated List:_ These lists can be viewed by clicking the user's username in the nav bar. The stats are displayed at first, and on further clicking the entire lists are displayed.

### Cypress testing

## Backend

### Go testing

Tests created for every API call

![TestingSprit4](https://user-images.githubusercontent.com/96463545/164143172-7347dca4-da40-4ae6-878b-1f2ba979ea67.gif)

### APIs made in this sprint:

#### 9. Get users rating of movies API

    GET  /userLikes/:username

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 10. Save user preference of a movie API

    PUT  /updateUserLike

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 11. Display every users movie rating API

    GET  /displayLike

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 12. Display movies of a genre API

    GET  /getGenre/:GenreId

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 13. Display movies of a genre with page numberAPI

```
   GET /getGenraPage
```

```
{"Genre id": genre,
"PageNumber: page
}
```

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 14. Display a user's profile numberAPI

```
   GET /getUserProfile/:username
```

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 15. Adding friends to an user API

```
   PUT /putFriends
```

```
{"username": user1,
"friendname: user2
}
```

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 16. Displaying friends of a user API

```
   GET  /getFriends/:username
```

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 17. Showing all friends of every user API

```
   GET /displayFriends
```

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request

#### 18. Returns a list of all the users whos name start with the given input API

```
   GET /getUsers/:username
```

##### Header:

    Authorization: Access token

#### Status Codes:

- **200**: Status OK
- **400**: Bad Request
