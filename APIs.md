#Sprint 3:

#### 1. Search API

This API is used search for movies

```
GET  /search/:name
```

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

#### 2. Top movies API

This API allows user to see list of top movies.

```
GET  /top
```

##### Header:

    Authorization: Access token

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

####  3. Search with page number API

This API is used to Search for a movie with page number.

```
GET  /searchPage
```

```
{ "name" : movie name,
  "page" : page number }
```
##### Header:

    Authorization: Access token

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

####  4. Register user API

This API is used to register a user.

```
put  /userReg
```

```
{ "username" :  name,
  "password" : pass }
```
##### Header:

    Authorization: Access token

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

####  5. Display authentication table API

This API is used to show all users and password.

```
GET  /displayPass
```


##### Header:

    Authorization: Access token

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

####  6. authenticate user API

This API is used to login using username pass.

```
PUT /authenticateUser
```
#####  Example Request Body:

```
{
    "username": user,
    "password" : pass
}
```

##### Header:

    Authorization: Access token

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable


#### 7. Get a movie API
This API fetches a movie

    GET  /getMovie/:movie


##### Header:

    Authorization: Access token


#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request


#### 8. Get Movie detials with the unique user API
This API gets the profile and movie detials of a user's movie

    GET  /getMovieOfUser/:movie/:username
    
##### Header:

    Authorization: Access token


#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request

### Sprint 4

#### 9. Get users rating of movies API

    GET  /userLikes/:username

##### Header:

    Authorization: Access token


#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request



#### 10. Save user preference of a movie API

    PUT  /updateUserLike

##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request



#### 11. Display every users movie rating API

    GET  /displayLike
    
##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request


#### 12. Display movies of a genre API

    GET  /getGenre/:GenreId
    
##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request

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

-   **200**: Status OK
-   **400**: Bad Request

#### 14. Display a user's profile numberAPI
```
   GET /getUserProfile/:username
```
##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request

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

-   **200**: Status OK
-   **400**: Bad Request

#### 16. Displaying friends of a user API
```
   GET  /getFriends/:username
```

##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request


#### 17. Showing all friends of every user API
```
   GET /displayFriends
```
##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request

#### 18. Returns a list of all the users whos name start with the given input API
```
   GET /getUsers/:username
```
##### Header:

    Authorization: Access token

#### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
