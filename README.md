# Greenlight Website

This is a Website I made in GO, it has the next features:

-   Preflight CORS and simple CORS.
-   Permission Based Authorization (READ - WRITE ACCESS control).
-   User Authentication.
-   User Activation.
-   Sending emails.
-   User Registration.
-   Graceful shut down of the server tu respond to last infly requests.
-   Rate Limiting to prevent excessive strain on the server.

<img width="1600" height="896" alt="greenlight logo for database" src="https://github.com/user-attachments/assets/b36d393d-f543-4b90-8379-8a341254637d" />

## CORS

This is a request made with Preflight CORS by another server into my API.
From a different origin, 

The web browser first sends a preflight CORS request and I sent valid request methods and valid request headers for real request.

Then the real request is made for user authentication.

![ale](cors.png)

## EMAIL

This is an Email sent to user:

![ale](emails.png)
