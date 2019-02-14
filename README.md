# gqlgen-todos

Simple GraphQL API generated using [`gqlgen`](https://gqlgen.com/).

## Run
- Generate the GraphQL execution runtime
```sh
go generate ./...
```

- Start the GraphQL server
```sh
go run server/server.go
```

- Try some queries
```graphql
mutation createTodo {
  createTodo(input: { text: "todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}

query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}
```

## Project
- `gqlgen.yml` - `gqlgen` [config file](https://gqlgen.com/config/)
- `generated.go` - GraphQL execution runtime, entirely generated by `gqlgen`.
- `resolver.go` - The bulk of your application code. `generated.go` will call into this to get the data the user has requested.
- `models/generated.go` - Generated models required to build the graph. These are usually augmented with non-generated models that map to a database schema (see `models/todo.go`). `gqlgen` will generate types for models that are defined in a GraphQL schema and *not* in a database (and not defined elsewhere in the project) like input types, which output to this file.
- `server/server.go` - Minimal entry point that sets up an `http.Handler` to the generated GraphQL server.
- `scripts/gqlgen.go` - Script to execute the `gqlgen` binary. This is called via the instruction in `resolver.go` using the `go generate` command.