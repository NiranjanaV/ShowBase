# ShowBase
Software Engineering Project - 2022
# Team Members
- Niranjana Vasudevan: Frontend
- Sanjay Srinivasan: Frontend
- Swaathi Reena Velavan: Frontend
- Kasiviswanathan Srikant Iyer: Backend
# Technology Stack
FrontEnd: React, Backend: GO

## Run these commands before building the React App
- npm install --save moment
- npm install font-awesome --save

# Project Description
ShowBase website, where we can create account, track the shows we are watching, add likes and dislikes, follow friends and see what shows they are watching, 
what they like and dislike watching.
# Sprint 1
## Frontend
- HomePage with content about the website and which contains 7 buttons - Home, About, Features, Movies & Shows, Completed, Stopped, Contact and Footer.
## Backend 
-In GoLang we have managed to call an API functionality which returns a JSON file with nested Json elements and arrays, which have been parsed using unmarshalling.The required data has been repacked (marshalled) and sent to the front end using HTTP request. On running the file, we can observe the local host set port receives the JSON.
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


