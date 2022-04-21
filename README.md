# ShowBase

Software Engineering Project - 2022

# Team Members

## Frontend

- Niranjana Vasudevan
- Sanjay Srinivasan
- Swaathi Reena Velavan

## Backend

- Kasiviswanathan Srikant Iyer

# Technology Stack

FrontEnd: React, Backend: GO

## To start the server

- cd /backend
- go run main.go

## Run these commands before building the React App

- npm install --save moment
- npm install font-awesome --save

# Project Description

ShowBase website, where we can create account, track the shows we are watching, add likes and dislikes, follow friends and see what shows they are watching,
what they like and dislike watching.

- Homepage: Description about our website's features and contact and link to our github profile. 
- Movies and shows Page: Displays all the movies in the database and when clicked on a movie's poster each user can see the movie rating and rate the movie, add movie to their watched and watchlist section. Users can also see the genre and click on a particular genre to see other movies in the same genre. It also has link to Add page which has search functionality to search for particular movie and again we'll be able to click on the poster and do many things. 
- Sign in Page: This page has the sign section and sign up option. Users can first sign up then sign in into our website and their own session will be created
- Profile Page: After signing in Users will be able to see their own Profile Page. This profile page shows the number of movies in users watched and watchlist and also user can click on the icons and go to watched and watchlist page to see the movies that they added there. There's add page which users can use to search and add friends. Then they can also see their friends watched and watchlist shows from there. 

## Links

- _Link to API Documentation:_ https://github.com/NiranjanaV/ShowBase/blob/main/APIs.md
- _Link to Project Board:_ https://github.com/NiranjanaV/ShowBase/projects/1
- _Link to Sprints Documentation:_
  - Sprint 1: https://github.com/NiranjanaV/ShowBase/tree/main/Sprints/Sprint%201
  - Sprint 2: https://github.com/NiranjanaV/ShowBase/tree/main/Sprints/Sprint%202
  - Sprint 3: https://github.com/NiranjanaV/ShowBase/tree/main/Sprints/Sprint%203
  - Sprint 4: https://github.com/NiranjanaV/ShowBase/tree/main/Sprints/Sprint%204

##  Videos

- Cypress Testing: https://drive.google.com/file/d/1VNPQb-ZeYBD3EldMbDG-SntC7FCAzArM/view?usp=sharing

# Sprint 1

## Frontend

- HomePage with content about the website and which contains 7 buttons - Home, About, Features, Movies & Shows, Completed, Stopped, Contact and Footer.

## Backend

- In GoLang we have managed to call an API functionality which returns a JSON file with nested Json elements and arrays, which have been parsed using unmarshalling.The required data has been repacked (marshalled) and sent to the front end using HTTP request. On running the file, we can observe the local host set port receives the JSON.
- Sprint 1 Demo Link: https://youtu.be/fNlQOtGbmYA
- Sprint 1 UI GIF:
  ![video](https://user-images.githubusercontent.com/71694219/158490462-1cd3c540-0e1a-4e69-ac69-60f4c073feb2.gif)

# Sprint 2

## FrontEnd

- Movies & Shows Section has been created. On clicking the Movies & Shows button from the top bar, it redirects to the list of movies in the database.
- Profile Page: On clicking the Profile button from the top bar, it redirects to the page where we can track the Completed and Stopped movies & shows.
- Profile>Add. We have added a new feature on the top bar called "ADD"
- Add > Search bar. On searching the show name, the shows come up in the search and the data is is displayed everytime a new character is entered.
- Sign In: Sign in page has been created for new users. You can either Login or SignUp for a new account. Backend Functionality still under progress.
- Testing: Frontend Testing has been performed using Cypress.

## Backend

- Backend code using go language has been written for User Registration,Enabling Password Visibility, Authenticate User, Get Movie, User Likes, Updating User Like, Displaying Like, Genre of the Movie, Get User ID, Search any movies, Search page.
- GIFS:

## Backend Testing:

![AutomatedBackendTesting ](https://user-images.githubusercontent.com/71694219/159502473-44d0a354-c82e-4a69-b1bd-d9edeb4b0e08.gif)

## Front end Testing:

![cypresstesting](https://user-images.githubusercontent.com/71694219/159502671-dc956a9a-c4ab-47af-b538-ca4374c586d1.gif)

# Sprint 3

## FrontEnd

- Added 2 new functionalities: Watched and Watchlist. User can mark the show as Watched or Add it to the Watchlist if they haven't watched the show.
- ![Profile_Frontend](https://user-images.githubusercontent.com/71694219/161363186-007ecc70-a188-46fa-868a-6dd785ec4c22.gif)
- Added navigation bar to all pages so that you can redirect to home page from any page on the website.
  ![Navigation](https://user-images.githubusercontent.com/71694219/161363191-682758a0-1c28-49fa-9371-027223683cb8.gif)
- In the contact section, if you click on any mail id, it will redirect you to the Outlook app and autofill the 'To' and 'Subject'.
- ![mail](https://user-images.githubusercontent.com/71694219/161364702-a1237799-5938-4fcc-a8df-630d63502f73.gif)
- Fixed CSS issues: Search bar alignment, About section alignment.
- New UI color changes
- Removed Completed and Stopped and replaced it with Watched, Watchlist and Add.

### Cypress testing

- To run the Functional tests, we run the following command -

  npx cypress open

- We have done it for Home page and Profile page.
- In Home page it checks if all buttons and links are working properly
- In profile page it checks if it searches for a movie correctly and checks the working of buttons and links in watched and watchlist page.
  ![cypressTest-sprint3](https://user-images.githubusercontent.com/30584808/161365272-eda7d2c2-4ac9-45bf-b4e3-726eb416aac5.gif)

## Backend

![Requests](https://user-images.githubusercontent.com/96463545/161365623-15853c7a-fd6c-4af3-a11d-4d3fafd2740e.gif)

- When we search for movies, the data comes from backend database.
- Returned all the movies of a genra.
- All the required tables created.
- Friend table concept introduced to socialise between people and to increase the functionalites to include group selection.
- Login functionality bugs hammered out.
- searchMovie.go. This API is used to search the movie database for a particular movie.
- Example response:
- <img width="1440" alt="Backend - search " src="https://user-images.githubusercontent.com/30584808/161364893-487cf506-03c9-44bd-9314-1e0236156bc9.png">

### Go testing

![GoTest-sprint3](https://user-images.githubusercontent.com/96463545/161365605-764ee7b9-7be0-47fc-82ea-aab9260a58b6.gif)

### APIs Created in 3rd sprint

#### 1. Search API

#### 2. Top movies API

#### 3. Search with page number API

#### 4. Register user API

#### 5. Display authentication table API

#### 6. authenticate user API

#### 7. Get a movie API

# Sprint 4

## FrontEnd

- _Profile:_ New Profile Page is created where the user can Add New Friends, view All Friends and learn about their friend's Watched and Watchlist.
  - _My Friends:_ The 'My Friends' page which will display all friends. The user can also look into their friend's Watched and Watchlist to find out what they're upto.
  - _Add New Friends:_ The user can search for their friend's profile and add them to friends list.
  - _Add:_ In the new add feature, you can search for the movie that you want to see, and when you click on the 'Movie Detail' button, it will display all the details of the movie. You can also Rate the movie here, Add it to your Watched and Watchlist. 
  - _Login and Logout:_ Login has been implemented using protected routes and react context. Incorrect credentials are checked, and if the credentials are checked appropriate rerouting is done using react navigate. 
  - _Movie Detail:_ Movie Detail page has added functionality such as add to watch list, mark as watched. The page also marks this with a green color if it has already been added by the user. Rating has also been added. Users may rate the movie out of 10. This state is also maintained on reload. 
  - _WatchList, To Watch and Rated List:_ These lists can be viewed by clicking the user's username in the nav bar. The stats are displayed at first, and on further clicking the entire lists are displayed.

![sprint4](https://user-images.githubusercontent.com/30584808/164369204-6d5b052b-1508-4754-adc0-c9b12d229c73.gif)


### Cypress testing
https://drive.google.com/file/d/1VNPQb-ZeYBD3EldMbDG-SntC7FCAzArM/view?usp=sharing

## Backend

### Go Testing sample

![TestingSprit4](https://user-images.githubusercontent.com/96463545/164143017-b6e13d12-be9e-474f-a65f-5cc2768b684e.gif)

## Other tests performed:

- Test Get Authenticate User API Response
- Test Get Movie API Response
- Test Get Movies Of Genre Likes API Response
- Test Get Genre With Page API Response
- Test Get Movie With User API Response
- Test Get User Likes API Response

### APIs Created in 4th sprint

#### 9. Get users rating of movies API

#### 10. Save user preference of a movie API

#### 11. Display every users movie rating API

#### 12. Display movies of a genre API

#### 13. Display movies of a genre with page numberAPI

#### 14. Display a user's profile numberAPI

#### 15. Adding friends to an user API

#### 16. Displaying friends of a user API

#### 17. Showing all friends of every user API

#### 18. Returns a list of all the users whos name start with the given input API
