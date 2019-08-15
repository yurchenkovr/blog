#BLOG
The following dependencies are used in this project:

- echo - HTTP 'framework'
- go-pg -  PostgreSQL ORM
- jwt-go - JWT Authentication
- bcrypt -  Password hashing
- Yaml - Unmarshalling YAML config file

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

    - POST:   /users/login             - accepts username/passwords and returns jwt token
    - GET:    /users/name/:username    - Gets User by username
    - GET:    /users/:id               - Gets User by ID
    - GET:    /users/                  - Gets all Users
    - POST:   /users/                  - Creates a new User
    - DELETE: /users/:id               - deletes a user(after sign in)
    - PATCH:  /users/:id               - updates a user(after sign in)
    - PATCH:  /users/bl/:id            - blocks the user to add posts
    - PATCH:  /users/unb/:id           - unblocks the user to add posts  
    
    - GET:    /articles/                  - Gets all articles
    - GET:    /articles/:id               - Gets article by ID
    - GET:    /articles/name/:username    - Gets article by username(author)
    - POST:   /articles/                  - creates articles new article(after sign in)
    - DELETE: /articles/:id               - deletes an article(after sign in)
    
