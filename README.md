Go Router
=========

# Installing

    $ go get github.com/c0rrzin/router

# Usage

## Defining routes

    DefRoute("POST", "/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello POST")
    })
    DefRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello GET")
    })

## Registering routes
Before you instantiate the server:

    func main() {
      ...

      RouteAll()

      ...
    }
