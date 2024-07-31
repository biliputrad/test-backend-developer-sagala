# test-backend-developer-sagala

## How to run the project

1. Clone the project
2. Run go mod tidy
3. Setup the database
4. Crate a app.env file and copy the content of app.env.example to it
5. Fill the app.env file with your database credentials and other configurations
6. Run go run .
7. For testing you can use the postman collection in the postman folder
8. Make sure to set the environment variable in postman to the correct url
9. You must create a user and login to get the token for the articles endpoints