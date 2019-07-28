#BLOG
The following dependencies are used in this project:

- echo - HTTP 'framework'
- go-pg -  PostgreSQL ORM
- jwt-go - JWT Authentication
- bcrypt -  Password hashing

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

    - GET:    /u/name/:username    - Gets User by username
    - GET:    /u/:id               - Gets User by ID
    - GET:    /u/                  - Gets all Users
    - POST:   /u/                  - Creates a new User
    - DELETE: /u/log/:id           - deletes a user(after sign in)
    - PATCH:  /u/log/:id           - updates a user(after sign in)
    - DELETE: /u/log/admin/:id     - deletes any user(after sign in as ADMIN)
    - PATCH:  /u/log/admin/bl/:id  - blocks the user to add posts
    - PATCH:  /u/log/admin/ubl/:id - unblocks the user to add posts  
    
    - GET:    /a/                  - Gets all articles
    - GET:    /a/:id               - Gets article by ID
    - GET:    /a/name/:username    - Gets article by username(author)
    - POST:   /a/log/              - creates a new article(after sign in)
    - DELETE: /a/log/:id           - deletes an article(after sign in)
    - PATCH:  /a/log/:id           - updates an article(after sign in)
    - DELETE: /a/log/admin/:id     - deletes any article(after sign in as ADMIN)
    
    - POST: /signin/               - accepts username/passwords and returns jwt token 
    - GET:  /welcome               - gets some information about the logged user 
    - POST: /refresh               - refresh token