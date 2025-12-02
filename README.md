### INSPIRED BY 
- Tiago's video [[https://www.youtube.com/watch?v=s3XItrqfccw]]
- Twelve factor [[https://12factor.net/]]

# Ecommercy API written in Golang
I am trying to be the most minimalist as possible, relying most on the standard library. The only library used is Ubers's Fx to better organize the business logic. 

# Technologies
- PostgreSQL
- Docker Swarm
- Fx [[https://github.com/uber-go/fx]] 

## Internal packages summary

### Internals package
Files not intended to be changed.

### Modules package
Hold the logic of database, auth and logger modules ( each of which could also be a package by it's own ).

### Handlers package
Hold the logic of what would be the controllers in other frameworks

### Handlers package
Table definitions

### cmd package
Application's entrypoint
