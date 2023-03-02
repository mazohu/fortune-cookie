# Sprint 2: Fortune Cookie ("on the verge of a breakthrough")

## Work Accomplished
### Front-end
The front end team added more pages by following the original wireframe plan. The user-home menu was added along with the eat a cookie and see past fortune pages. 

As a group, we all attempted to find the best login solution to integrate our front and backends respectively. Ben worked []() on integrating the Auth0 authentication platform, focusing heavily on the front-end work that goes into integrating it into the end-to-end product.

Gabi tried to implement this [youtube](https://www.youtube.com/watch?v=1hMvJsSDnvU) tutorial that would allow for users to sign in with google and communicate with the backend. However, after implementing the tutorial and trying to resolve errors the implementation was still unsuccessful. Best guess is that the dependencies of the package and the angular file were not compatible with versions of one another. This is something to address for next sprint.
### Back-end
Back-end team collaborated with front-end to execute a golang To-Do app example in which the Golang server handled requests to create, delete, complete, and display todos to the Angular client. Part of this example was integrated with work from Sprint 1 to make the development environment portable using [viper](https://github.com/spf13/viper).

Additionally, the back-end team began development on a database for users using [GORM](https://gorm.io/docs/index.html) and generated unit tests of updates to the user table and of the system clock feature for limiting user submissions to once per day.
## Unit Tests
### Front-end
The unit test incorporated in the front-end tests for the ability to mount the home component through the Cypress testing framework.

This test is relatively simple, and is included in this sprint simply to show that the functionality for future unit testing is there, and that the Cypress framework can successfully detect and interact with our Angular components.

Having passed this component-mounting test, we now are able to expand onto larger and more complex unit tests, such as testing how the page changes as a result of certain user interaction.

Additionally, we can now take steps to start setting up end-to-end testing in the future.

### Back-end
[storage.go](https://github.com/mazohu/fortune-cookie/blob/alexia-5/storage.go)
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

`type Users struct`
: This is the struct which contains the columns of the user database. The contents are *Username  string, Email string, UserID string, Fid string, Submitted bool, and LastTime time.Time*

`func getUserPointer(username string, em string, uID string) (Users, *gorm.DB)`
: Accesses a pointer to the database row of the currently logged in User. This userPointer and the database is returned for use in other functions. 

`func submittedCheck(userPointer Users) (Users)`
: CAlls both checkTime and updateTime to see if the user's saved LastTime is more than a day old. If it is, it will update the time in the database. 

`func checkTime(userPointer Users)(bool)`
: Checks the current time and compares it to the saved time stored in the User database. This pertains to the last time the logged in user made a fortune. 

`func updateTime(userPointer Users) time.Time`
: Updates the time in the database