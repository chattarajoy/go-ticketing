# Folder Structure

```bash
.
├── cmd
│   └── server
├── docs
│   └── api
├── internal
│   ├── cache
│   ├── json
│   ├── router
│   ├── server
│   ├── testhelpers
│   └── workgroup
├── pkgs
│   ├── api
│   ├── models
│   └── service
│       ├── booking
│       ├── cinema
│       └── movie
└── test-reports
```


* `pkgs` directory contains
    * models
    * services (light version of controller)
        * business logic of the app
        * contracts for each application / API
    * handlers (handles request related operations of a controller)
        * serialize / deserialize requests
        * Handle errors and status codes
  
* `internal` directory has
    * helpers
    * http server
    * http router 
    * interface for cache

* `cmd` directory is related to command line specific content which includes
    * Initializing the server
    * running DB Migrations [server.go](../cmd/server/server.go)
    * setup Routes [routes.go](../cmd/server/routes.go)
    * handlers or middlewares that can be added to each request [handler.go](../cmd/server/handler.go)