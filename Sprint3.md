# Lord of the Fortune Cookies: The Return of the Fortune Cookie (Sprint 3)

## Work Accomplished

### Front-end
The frontend work for this sprint consisted of two main parts:
1. Redoing our frontend visual design and structure from the ground-up, without relying on pre-made Angular frontend UI kits, and instead leveraging the Bootstrap CSS framework in order to use a more up-to-date Angular version that was compatible with the backend tech stack as well as the Cypess testing framework
2. Built a new frontend foundation that uses a temporary mock HTTP server to asynchronously send and receive data involving fortunes and user login sessions, all while keeping track of this in a built-in logging system ("messaging" system) that can interact with the backend to give information on the latest user actions on the app, as well as information on errors or HTTP requests that fail/succeed.

For part 1, we wanted to guarantee that the new front-end would have the same utility, visual impact, and usability as the last UI kit and theme (notus-angular). By starting out with an up-to-date Angular project as well as a compatible back-end tech stack, we were able to leverage the Bootstrap CSS framework to make all of our pages from scratch. However, this time we were able to integrate the google authorization as we made the pages. So, authorization to certain pages (based on if the user was logged in) and the login functionality was able to be implemented. 

The motivation behind #2 is that we want to ensure that we head into the final sprint with a solid idea of how we are going to integrate everything together. With this, we now have a proper working example where the frontend can handle HTTP traffic (GET, POST, PUT, etc.), log and send data to both the backend (for devs to look at) and the frontend (for users to know that there was an error), and asynchronously query and receive all of this data -- that is to say, our app can keep running other functionality while at the same time waiting for a response or request that it does not expect to happen immediately. With this in place, it is now just a matter of interfacing this with the backend's database made with GORM, along with their authentication mechanism (Google Auth).


### Back-end
The back-end was focused on improving our databases and making sure the information was accessible to the front-end and vice versa. The creation and access of the database tables and content were related directly to different HTTP requests from the front-end. HTTP traffic is handled mostly by Go fiber and pusher. Post requests were used to add entries to the database. Get requests were used to send database information from the back-end to the front-end. To make this work, I needed to move most of my database function out of Storage.go. I removed many unused functions from that file, but there are still some functions left that deal with the LastSubmitted content of the Users database. 

Besides back-end and front-end integration, the back-end improved the structure of the Users database and made a Fortune database. We used GORM to handle the databases. There was an attempt to use GORM's many2many realtionship to link together the Users and Fortune databases by userID, but there were many unresolved bugs. Therefore, for this sprint, we settled on our initial plan: the Fortune ID's that make up a User's history would be stored as a list of integers, concatinated into a string. Then, to get a user's history, we would split the string, read each Fortune ID, find it in the fortune database, and return a list of all the contents. It would be more tedious than if we were to use a GORM function, therefore we want to change this in the future and make the many2many relationship (or something else provided by GORM) work with out schema. 


### Inital Issues From Sprint 2
#### Back-end and front-end integration made difficult
Our front-end template was dependent on a version of angular that wasn’t compatible with ways of integrating with our backend. 
#### Log-in was carried out in the backend
For our log-in to be carried out through our third-party authentication method (Google), we should have implemented it in the frontend. 

#### Goals
- Either incorporate our current front-end template into the most updated version of Angular (Angular 15), or make a new compatible front-end.
- Complete front-end and back-end integration, with a working database and log-in method
- Incorporate more testing
- Start the fortune database (and incorporate it with the user database if time allows)
#### Initial Integration as a Team
Each of us spend hours after the last sprint submission exploring ways to integrate the front-end and the back-end. Some of us continued to look for a way with the front-end template we currently had. Others used temporary, simple, and compatible front-ends to at least get integration with the back-end down before further developing the front-end. Eventually, after we had a successful integration method with the simple front-end. We started building our front-end again from that working version. 

## Testing
### Front-end
For this sprint, since we now have HTTP client/server functionality to retrieve and send data, a logical step forward was to start incorporating **end-to-end testing**, which we had not done up to this point for the frontend.

The approach to end-to-end testing for our Fortune cookie app revolves around using the HTTP and Javascript data-sending and receiving functionality of the Cypress framework, along with its ability to parse the traffic, and consequently the HTML and HTML DOM atributes that get returned/sent.

For this sprint, our end-to-end testing involved testing a sample to-do app unrelated to our Fortune Cookie app, just to get the understand how to make them, and then, an end-to-end test for our Fortune Cookie dashboard page that visits the local web server and verifies 1) that it is online and reachable, 2) that it has the right title.

The end-to-end testing that we wrote is in the `ben-angular-learning` branch, located at `angular-tour-of-heroes/cypress/e2e/app.e2e.spec.js`:
```javascript=
describe('Tour of Heroes App Testing', () => {
    it(`Should have the title 'Fortune Cookie'`, () => {
        cy.visit('localhost:4200')
        cy.contains('Fortune Cookie');
    });
});
```
See the submitted video by Ben for more details.

### Back-end
[storage.go](https://github.com/mazohu/fortune-cookie/blob/alexia-8/back-end/storage.go)
- `func dataBaseTesting(username string, em string, uID string)`
- `func fortuneTimerTesting(userPointer Users)`


## API Documentation
`func setAuth0Variables()`
: Configures `viper` and parses a hidden `env` file for the `AUTH0_API_IDENTIFIER`, which identifies and validates the Golang server in the Auth0 scheme, and `AUTH0_DOMAIN`, which is used by the Auth0 server to access the Golang API

`func authRequired() gin.HandlerFunc`
: Defines a Gin handler function (which serves as middleware for the router) which validates the JWT token sent by the client

`func terminateWithError(statusCode int, message string, c *gin.Context)`
: Parses the statusCode of the HTTP request and, if there is an error, logs the error with the client and aborts the request

`func CORSMiddleware() gin.HandlerFunc`
: A custom middleware handler function for enabling Cross Origin Resource Sharing. We intend to replace this function using a suitable CORS Golang package.

`type User struct`
: This is the struct which contains the columns of the user database. The contents are `Username  string, Email string, ID string, fid string, Submitted bool, and LastTime time.Time`

`type Fortune struct`
: Defines the model for the fortune database which stores all of the user-generated fortunes. Structure is `{ID uint32, Author string, Content string}`

`func getUserPointer(username string, em string, uID string) (Users, *gorm.DB)`
: Accesses a pointer to the database row of the currently logged in User. This userPointer and the database is returned for use in other functions. 

`func submittedCheck(userPointer Users) (Users)`
: Calls both checkTime and updateTime to see if the user's saved LastTime is more than a day old. If it is, it will update the time in the database. 

`func checkTime(userPointer Users)(bool)`
: Checks the current time and compares it to the saved time stored in the User database. This pertains to the last time the logged in user made a fortune. 

`func updateTime(userPointer Users) time.Time`
: Updates the time in the database

`func hashFortune(fortune string)(uint32)`
: Takes the content of a user-submitted fortune and generates its unique hash to identify the fortune in the database

` app.Post("/api/user/populate", func(c *fiber.Ctx) error`
: Adds the user to the User database if not already there. Otherwise, it retrieves the information and sets our pointer to it.

` app.Post("/api/user/submitFortune", func(c *fiber.Ctx) error `
: Accesses the hashFortune function and adds the new fortune to the database as long as the submitted fortune isn't an empty string. 