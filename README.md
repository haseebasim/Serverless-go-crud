# Golang Serverless CRUD app

- Place your DB connection string in middleware.go file to connect to DB. 
- For building go files to binary executeables run `make` command in the root directory
- Once the bin folder is created, the lambda function can be deployed to aws. 

Note:

For testing the APIs make sure you are sending the correct parameters in correct format,

> GET : /{id} (Get a single user. Here the `id` is passed as path parameter)

> GET : /     (Get all the users)

> POST : /    (Insert a user. Accpets `firstname` and `lastname` as query params)

> PATCH : /    (Update a user. Accpets `id`, `firstname` and `lastname` as query params)

> DELETE : /    (Insert a user. Accpets `id` as query params)
