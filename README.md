## User Management API Endpoints

- This documentation outlines the endpoints available for managing user-related operations within the API. The API is versioned under `/v1/users`.

### Create a User

### POST `/v1/user/`

## Description: This endpoint allows the creation of a new user in the system.

## Request:

- `fullName` (string): The full name of the user.
- `email` (string): The email address of the user.
- `phoneNumber` (string): The phone number of the user.

# Example Request:
{
    "fullName": "John Doe",
    "email": "john.doe@example.com",
    "phoneNumber": "+1-555-123-4567"
}

# Example Response:
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

### Update User Information

### PATCH `/v1/user/`

## Description:  This endpoint allows updating user information such as full name and phone number.




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
