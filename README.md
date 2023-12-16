# Password-Manager CRUD API using Gofr

With the help of this CRUD API user can store edit delete their passswords for multiple platforms.

## Table of Contents
- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)

- ## Features

- - **User Authentication with JWT:**
  - Secure user authentication using JSON Web Tokens (JWT).
  - Users can register, log in, and receive a JWT token for accessing protected endpoints.

- **Password Encryption and Decryption:**
  - Utilizes AES encryption for secure storage and retrieval of user passwords.
  - Passwords are encrypted before being stored in the database and decrypted when retrieved.

- **Password Management:**
  - Users can store, retrieve, update, and delete their passwords.
  - Implements CRUD operations for managing user passwords.
 
  - - **One-to-Many Mapping:**
  - Establishes a one-to-many relationship between the `User` table and the `Password` table.
  - Each user can have multiple passwords, providing a comprehensive password management solution.

- **Secure API Endpoints:**
  - Implements secure API endpoints for user authentication and password management.
  - Protects sensitive operations with proper authorization checks.

- **Gofr Package Integration:**
  - Developed using the Gofr package in GoLang, providing a lightweight and modular framework for building APIs.

- **Well-Structured API Endpoints:**
  - Provides a clear and well-structured set of API endpoints for easy integration and usage.
  - Follows RESTful principles for consistency.

- ## Getting Started

- ## Installation:
  - Install Go:
- Firstly, ensure you have Go installed on your system. If not, follow the instructions [here](https://go.dev/doc/install).

  - Clone the Repository:
- Clone the main branch of the repository to your local system.

  - Open Code:
- Open the code using any text editor, such as Visual Studio Code.

  - Run the Application:
- Execute the following command to run the application:
- `go run main.go`

- ## API Endpoints
- **User Endpoints**
- `POST /users/signup`: Create a new user.
- `POST /users/login` : Login with email and password.
- `GET /users/getUser/{userId}`: Retrieve details of a specific user.
- `DELETE /users/deleteUser/{userId}`: Delete a user.

- **Password Endpoints**
- `POST /users/createPassword`: Create a new password.
- `GET /users/getPassword/{passwordId}`: Retrieve details of specific password.
- `GET /users/getPasswords/{userId}`: Retrieve list of passwords bases on specific users.
- `PUT /users/updatePassword/{passwordId}`: Update details of specific password.
- `DELETE /users/deletePassword/{passwordId}` Delete specific password.

- ## Usage
- Access this [PostmanCollection](https://www.postman.com/payload-specialist-84635644/workspace/password-manager/collection/13859246-e7b198e7-74b3-43e9-a501-10c9ae0f159b?action=share&creator=13859246) for testing the API.
- **Steps for testing the API**-:

  - **Create a New User**:

- Use the signup request from the provided Postman collection. The request includes mock data for testing.
  
  - **Login After Signup**:

- After successfully signing up, login with email and password using the login request.
- The login response will contain a JWT token.

  - **Use JWT Token for Authentication**:

- Use the obtained JWT token for authenticating subsequent requests.
- Add the JWT token in the Authorization header as Bearer "your JWT Token" for accessing protected endpoints.

  - **Create a Password**:

- To create a password, use the createPassword request.
  
- In the request header, add Bearer "your JWT Token".
  
- Also, add the userId obtained from the response body of the login request as a path parameter.

- User other requst given in the [PostmanCollection](https://www.postman.com/payload-specialist-84635644/workspace/password-manager/collection/13859246-e7b198e7-74b3-43e9-a501-10c9ae0f159b?action=share&creator=13859246).
