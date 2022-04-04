# Sprint 3
### In Sprint 3 the following user stories were successfully implemented:
### 2. In this User story, backend Integration for search functionality is done.
Users can search for a movie/show in the ‘Add’ page and data comes from the backend database and is displayed on the page. 
### 3. In this User story, added 2 new functionalities: Watched and Watchlist. 
Users can search and add a show to a new page in their profile ‘Watchlist’ if they haven't watched the show.
Users can search and add a show to a new page in their profile ‘Watched’ if they have already watched the show.
Users can move the show from ‘Watchlist’ to ‘Watched’ page after finishing the show.
Users can remove shows from ‘Watchlist’ or ‘Watched’ pages if they don’t want it displayed anymore.
### 4. In this User story, added a navigation bar to all pages so that you can redirect to the home page from any page on the website.
### 5. In this User story, in the contact section, if users click on any mail id, it will redirect them to the Outlook app and autofill the 'To' and 'Subject' so they can contact us anytime. 
 
### 6. CSS Fixes
Fixed CSS issues: Search bar alignment, About section alignment.
New UI color changes.
Removed Completed and Stopped and replaced it with Watched, Watchlist and Add.

### 7. In this user story, we haved added functionalities to Movies&Shows page
The data displayed in this page is received from the backend database. Movies are displayed according to the year released. Users can click on a movie
and see movie poster, movie description and movie genre. Once a click on a particular genre, you'll be redirected to another page where you can see 
other movies in the same genre

### 8. In this user story, the backend integration is done for the sign in page. 
Users can now go to the sign in page, sign up and then sign in to go to their profile. Their username and password(encrypted) is stored in the backend 
database.

### Video Links
Frontend: https://drive.google.com/file/d/1Ro08UHd2lYhPnxCM3SlkoNE0um4CVjr0/view?usp=sharing

## FrontEnd

- Added 2 new functionalities: Watched and Watchlist. User can mark the show as Watched or Add it to the Watchlist if they haven't watched the show.
- ![Profile_Frontend](https://user-images.githubusercontent.com/71694219/161363186-007ecc70-a188-46fa-868a-6dd785ec4c22.gif)
- Added navigation bar to all pages so that you can redirect to home page from any page on the website.
-![Navigation](https://user-images.githubusercontent.com/71694219/161363191-682758a0-1c28-49fa-9371-027223683cb8.gif)
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
