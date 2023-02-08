# Sprint 1 - Fortune Cookie

## User stories
- As a site member, I want to be able to send and receive ONLY ONE new fortune each day.
- As a site member, I want to write and receive a fortune by clicking on the "eat a cookie" button on the user menu.
- As a site user I want ot receive a random fortune from someone else
- As a site member, I want a way to recover my password in case I forget or lose it.
- As a user, I want to view my past fortunes to keep track of how good or bad my luck is.
- As a site member, I want my passwords to be stored securely in the database so that my account is kept private.
- As a visitor, I want to learn what the web application is about before signing up.
- As a user, I want to easily log into or create my account using a pre-existing third-party authentication system, so that I can quickly begin submitting and generating fortunes.
- As a site member, once logged in, I want to be able to navigate around the site through an interactive menu.
- As a site visitor, I want to be able to view, sign-up, or login from the landing page.
- As a user, I want to make sure my username and password are viable through input verification by the login.

## What issues your team planned to address

The front-end team planned to address creating a landing page, an about the web app section, as well as creating sign in and login pages. They also planned to update the internal routing (our temporary, "mocked backend") and pre-templated UI-kit pages as necessary.

The backend team planned to implement a log-in system using third-party authentication through Google so that users could easily log-in and use the website.

## Which ones were successfully completed

Front-end successfully completed their planned tasks for Sprint 1.

Back-end successfully implemented a log-in strategy using Google third-party authentication. 

## Which ones didn't and why?

Though the back-end hoped to establish a working server from which to serve front-end files before the end of Sprint 1, we ran into trouble with archived and/or old go package dependencies which made it difficult to accomplish both serving files and establishing an authentication system. However, since we were able to retrieve user information from Google's authentication server, we are in a good place to begin building a database from which to serve user requests. 
