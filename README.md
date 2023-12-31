# Distributes

- api: user facing part of the application
- core: runs in the background in two modes (provider and resolver)
- public: user facing part of the application
- admin: inhouse management part of the application
- docs: documentation with markdoc

# Documentation for Gymbuddy

this is the official documentation for the gymbuddy web application.

Jump To:

- [Environment Setup](#Environment-Setup)
- [API Endpoints](#API-Endpoints)

## Environment-Setup

In the root directory of the application, create an 'app.env' file with the following;

    POSTGRES_HOST=localhost
    POSTGRES_DB=postgres
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_PORT=5432
    PORT=:3000
    DB_DSN=postgres://postgres:postgres@localhost/postgres?sslmode=disable
    DB_MAX_OPEN_CONNS=10
    DB_MAX_IDLE_CONNS=5
    DB_MAX_IDLE_TIME=1m

## API-Endpoints

### User Management API Endpoints

- This documentation outlines the endpoints available for managing user-related operations within the API. The API is versioned under `/v1/users`.

### Create a User
    POST `/`

Request body is going to include:
- `fullName` (string): The full name of the user.
- `email` (string): The email address of the user.
- `phoneNumber` (string): The phone number of the user.

Example Request:

    {
    "fullName": "John Doe",
    "email": "john.doe@example.com",
    "phoneNumber": "+1-555-123-4567"
    }

Example Response:

    {
    "message": "User created successfully.",
    "details": {
    "ID": "user_id",
    "AccountType": "regular",
    "FullName": "John Doe",
    "Email": "john.doe@example.com",
    "PhoneNumber": "+1-555-123-4567",
    "Activated": true
      }
    }

### Set User Information
    PATCH `/`
Description: This endpoint allows updating user information such as full name and phone number.

### Create User Bio
    POST `/bio`

### Set User Bio
    PATCH `/bio`

### Create User Preferences
    POST `/pref`

### Set User Bio
    PATCH `/bio`


### User Sessions

- This handles uers sessions that checks if a user has a valid/active session.

- POST `/`

### Description: creates a models.SessionModel for the user valid for the current day if the user has a models.SessionModel that is valid, it will ignore the request with an error.

- POST `/:sessionId`

### Description:

- GET `/requests`

### Description:

- GET `/connections`

### Description:

- GET `/sessions`

### Description:

- POST `/accept/:requestId`

### Description:

- POST `/decline/:requestId`

### Description:

## Your Gym Buddy

- Account Types
  - Client
- Sign In Options
  - OAuth
    - Google
    - Apple
  - Email

### Client

- Email
- First
