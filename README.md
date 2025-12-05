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
Files not intended to be shared; Organized by logical layers instead of domain layers. 

### Pkg package
Hold the logic of database connections, authentication, caching, routing, overall configuration, backend (server), and the logging system. The 'fx' postfix denotes that each module is handled by the uber-go/fx DI framework. 

### cmd package
Application's entrypoint
