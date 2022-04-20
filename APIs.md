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


