# ShowBase
Software Engineering Project - 2022
# Team Members
- Niranjana Vasudevan: Frontend
- Sanjay Srinivasan: Frontend
- Swaathi Reena Velavan: Frontend
- Kasiviswanathan Srikant Iyer: Backend
# Technology Stack
FrontEnd: React, Backend: GO
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
- Added navigation bar to all pages so that you can redirect to home page from any page on the website.
- In the contact section, if you click on any mail id, it will redirect you to the Outlook app and autofill the 'To' and 'Subject'
- Fixed CSS issues: Search bar alignment, About section alignment.
- 
  
## Backend


