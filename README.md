# Cuisine api project

Dummy project in Go that implements a very simple API to store cuisine receipts

## Testing the API with curl

### register a new user

To register a new user, you need to send a POST request to the /api/register endpoint with a JSON payload containing the username and password.

    curl -X POST -H "Content-Type: application/json" -d '{"username":"testuser","password":"testpassword"}' http://localhost:8080/api/register

### Login with a Registered User

To log in with a registered user, you need to send a POST request to the /api/login endpoint with a JSON payload containing the username and password. This will return a cookie with a JWT token if the login is successful.

    curl -X POST -H "Content-Type: application/json" -d '{"username":"testuser","password":"testpassword"}' http://localhost:8080/api/login -i

The -i option in the curl command includes the HTTP response headers in the output, which will show you the Set-Cookie header containing the JWT token.

### Testing an Authenticated Endpoint

After successfully logging in and receiving the token cookie, you can test an authenticated endpoint by including the cookie in your curl request.

    curl -X GET http://localhost:8080/api/recipes -b "token=<JWT_TOKEN>"

### create recipe
    curl -X POST -H "Content-Type: application/json" -d '{"name":"Spaghetti Carbonara","ingredients":"Spaghetti, Eggs, Parmesan Cheese, Pancetta","instructions":"Boil spaghetti. Fry pancetta. Mix eggs and cheese. Combine all."}' http://localhost:8080/api/recipes -b "token=<JWT_TOKEN>"

### get all recipes
    curl http://localhost:8080/api/recipes -b "token=<JWT_TOKEN>"

### get recipe by id
    curl http://localhost:8080/api/recipes/1 -b "token=<JWT_TOKEN>"

### update a recipe
    curl -X PUT -H "Content-Type: application/json" -d '{"name":"Spaghetti Carbonara","ingredients":"Spaghetti, Eggs, Parmesan Cheese, Pancetta, Pepper","instructions":"Boil spaghetti. Fry pancetta. Mix eggs and cheese. Combine all. Add pepper."}' http://localhost:8080/api/recipes/1 -b "token=<JWT_TOKEN>"

### delete a recipe
    curl -X DELETE http://localhost:8080/api/recipes/1 -b "token=<JWT_TOKEN>" 
