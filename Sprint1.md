FRONT END – DOCUMENTATION

HOME PAGE - ShowBase

Created seven different buttons – HOME, ABOUT, FEATURES, MOVIES and SHOWS, COMPLETED, STOPPED and CONTACT, which when clicked on will go to that particular location on the webpage. This was done using hyperlinks by creating the element target and a link to that target in java script.

HOME
This displays the home section of the website. It explains to users about the website in a short and interactive manner.

ABOUT
This section shows the overview of what our ShowBase website will help the user with. It will help user Track, discover, share, and learn about new TV shows and movies. We have created the boxes in CSS and displayed the data.

FEATURES
This section shows all the features provided by this website in detail. It explains about having your own personal account, how to stay UpToDate, matching Tv shows with your friends and about reviewing directly from the website. The section was created with HTML and CSS with the zoom-in option done in JavaScript.

COMPLETED SHOWS
In this section, all the completed shows of the user will be displayed. We have set up a carousel using ReactJS. Upon completion, shows and movies will be displayed here. As of now, the images of the shows shown in the completed section are displayed manually. But, in the next sprint the completed shows will be tracked and displayed from the backend.

STOPPED SHOWS
Similarly, the same was done for shows that the user has stopped watching and don’t want to watch anymore. These are also displayed as a carousel with ReactJS.

CONTACT
In case of any issues, users can contact us using this section. We have create a box effect to display this.

FOOTER
This contains a GitHub logo, which acts as a hyperlink as well and takes you to our project GitHub page.

PROFILE PAGE

This page is created by parsing JSON data and then using map function to map the data into a react reusable component. React Router was used to route the web page, however this didn't perform as intended and has been marked as an issue.
This page will display as the users home page after the Login functionality has been implemented.


BACKEND FUNCTIONALITY
In GoLang we have managed to call an API functionality which returns a JSON file with nested Json elements and arrays, which have been parsed using unmarshalling.
The required data has been repacked (marshalled) and sent to the front end using HTTP request.
On running the file, we can observe the local host set port receives the JSON.