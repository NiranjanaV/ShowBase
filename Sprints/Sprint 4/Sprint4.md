# Sprint 4

## FrontEnd


### Cypress testing

  
## Backend


### Go testing


### APIs made in this sprint:

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
