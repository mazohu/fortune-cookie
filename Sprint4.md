# The Fortune Cookie Wars Episode 4: The Phantom Cookie (Sprint 4)

## Work Accomplished
### Front-end
The front-end team continued to implement the remaining functionalities left from sprint 3. Gabi focused on eat cookie, past fortunes, and logout. She updated the html, ts, and css files respectively to get the form and function of each page working. Ben used his time to create a separate user profile page and also helped clean up some of the pages' UI design. 

On an individual basis, Gabi spent some time catching up to what Ben had learned last sprint regarding the Tour of Heroes tutorial and how that related to the way we planned to implement our sending and receiving fortunes from the database.
Separately, Ben took the time to develop more end-to-end Cypress tests to confirm the security of the website navigation, ensuring that user pages and fortune sending/receiving was only possible if signed in. Finally, a change was made to separate out the user's profile into its own Angular component, as opposed to being displayed in the signed-in homepage.



### Back-end
There was major reorganization done in the back-end by Maria. The database information was moved to another go package called 'storage.go', where functions related to accessing a database row, updating database information, checking the day, and more was done. This left all the HTTP POST and GET requests in 'main.go', with functions from storage being called within those requests. There were also appropriate test functions created in a file called 'storage_test.go', all developed by Maria.

With the front-end already partially linked to the back-end through the userpage, an essential goal was to stitch the rest of the program together. This responsibility fell on Alexia, as she was the one who linked the two parts of the program initially. Three essential locations to do this was in 'eat cookie', 'past history', and 'user profile'. The back-end worked to make appropriate GET and POST requests for these pages, completing the link by adding to in the Typescript of those pages. Some POST and GET requests included finding/creating the user in the database, retrieving variables that the front-end needs about the user, and more.



## Testing
### Front-end

Front-end tests primarily targeted the end-to-end functionality of the web application, using Cypress's HTTP functionalities to query and visit webpages, navigate the Document Object Model (DOM) and interacting through Javascript with the various components/elements on the page.

Our tests unit + end-to-end testing consists of:
- Visiting the page on its corresponding network port, checking to see if the Fortune Cookie application is what has been returned.
- Attempting to visit the "User homepage" while not being signed in, checking returned DOM objects to see if the box containing the Google Sign-in option is present.
- Attempting to access content on the "User profile" page, which should not be possible when not signed in
- Attempting to "eat" a cookie (which gives you a fortune) as an unauthenticated user, which should fail and instead redirect to the Google Sign-in.
- Accessing past fortunes and checking if the table gets populated if not signed in.

In the future, further testing would implement Google authentication tokens within the Cypress testing framework to more broadly test the functionality of the application, with activities that signed-in users can carry out.


### Back-end
Backend unit testing involves various functions that ensure that database integrity is maintained and behaves appropriately with new user creation, fortune creation, submission flags, etc. The full list of tests, along with their documentation can be found at:

[https://github.com/mazohu/fortune-cookie/blob/main/back-end/storage/storage.go](https://github.com/mazohu/fortune-cookie/blob/main/back-end/storage/storage.go)
- `func dataBaseTesting(username string, em string, uID string)`
- `func fortuneTimerTesting(userPointer Users)`


## API Documentation
### Back-end
`app.Post("/api/user/populate", func(c *fiber.Ctx) error`
: Receives the username, email, and user ID from the front-end when logged in, and prompts for either creating a new entry in the users database or retreiving an existing user. 

`app.Post("/api/user/submitFortune", func(c *fiber.Ctx) error`
: Receives a new fortune from the front-end, storing it in the fortunes database if it has been more than a day. 

`app.Get("/api/user/frontend/getFortune", func(c *fiber.Ctx) error`
: In response to requesting a new fortune, this Get request sends a new, random fortune to the front-end.

`app.Get("/api/user/frontend/fid", func(c *fiber.Ctx) error`
: Sends a list of the current user's stored received fortunes to the front-end.

`app.Get("/api/user/frontend/lastdate", func(c *fiber.Ctx) error`
: Sends the lastdate (which is the current user's LastTime value but in a readable, string format) to the front-end.

`app.Get("/api/user/frontend/submitted", func(c *fiber.Ctx) error`
: Sends the current user's Submitted value to the front-end.

`app.Get("/api/user/frontend/todayfortune", func(c *fiber.Ctx) error`
: Sends the current user's LastFortune value to the front-end.

`app.Get("/api/user/frontend/lastTime", func(c *fiber.Ctx) error`
: Sends the current user's LastTime value to the front-end.

`func clearDatabase(dbfile string)`
: Deletes the entire database, mostly used for testing purposes.

---
`func InitStorage(file string)()`
: Creates or opens the sqlite database described by file and connects to this database. Because this function also populates the database with default fortunes and seeds the random number generator used in ReceiveFortune, this function should be called every time the server starts.

`func GetUser(user User)()`
: Creates or gets the User described by user and updates the global CurrentUser to point to this User. Since this function modifies or queries the User table in the database, it is necessary to pass a User with an initialized user ID.

`func GetSubmit() bool`
: Returns the CurrentUser's Submitted field. It is false if the user has not submitted and true otherwise.

`func GetLastTime() time.Time`
: Returns the LastTime that the CurrentUser submitted a fortune.

`func GetLastFortune() string`
: Returns the content of the last fortune received by the CurrentUser as a string.

`func CheckToday()`
: Checks whether the CurrentUser can submit today and, if true, updates the CurrentUser's Submitted and LastFortune fields and saves them to the database.

`func SubmitFortune(content string)(error)`
: Updates the fortune database with a fortune only if the user has not yet submitted a fortune. It creates a new Fortune with ID equal to the hash of content and content equal to the argument. If the operation fails, it returns an error.

`func ReceiveFortune()(Fortune, error)`
: Returns a random fortune not yet received by the user or an error if the database query fails. 

`func GetReceivedFortunes()([]Fortune, error)`
: Returns the list of fortunes that the user has received or an error if the database query fails.

`func GetFortunesNotReceived()([]Fortune, error)`
: Returns the list of fortunes not received by the user or an error if the database query fails.

`func FormatDate(ourTime time.Time)(string)`
: Returns the date ourTime as a string with the format Month Day, Year.

`func hashFortune(fortune string) (uint32)`
: Private function which returns the hash of the string fortune.

`func canSubmit()(bool)`
: Private function which returns true if today's date is different than the one stored in LastTime and false otherwise.
