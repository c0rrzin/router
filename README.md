Go Router
=========

# Installing

    $ go get github.com/c0rrzin/router

# Usage

## Importing

    import "github.com/c0rrzin/router"

## Defining routes

    router.DefRoute("POST", "/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello POST")
    })
    router.DefRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello GET")
    })

## Registering routes
Before you instantiate the server:

    func main() {
      ...

      router.RouteAll()

      ...
    }
