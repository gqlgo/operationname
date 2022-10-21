# operationname

[![pkg.go.dev][gopkg-badge]][gopkg]

`operationname` print operation name in your GraphQL query files.

```graphql
# Query
query GetUser {
   user {
       name
   } 
}

mutation UpdateUser($name: String!) {
    updateUser(name: $name) {
        id
        name
    }
}
```

## Usage

```sh
$ go install github.com/gqlgo/operationname/cmd/operationname@latest
```

The `operationname` command has two flags, `schema` and `query` which will be parsed and analyzed by operationname's Analyzer.

```sh
$ operationname -schema="server/graphql/schema/**/*.graphql" -query="client/**/*.graphql"
GetUser
UpdateUser
```

The default value of `schema` is "schema/*/**.graphql" and `query` is `query/*/**.graphql`.

`schema` flag accepts URL for a endpoint of GraphQL server.
`operationname` will get schemas by an introspection query via the endpoint.

```sh
$ operationname -schema="https://example.com" -query="client/**/*.graphql"
```


## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/operationname
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/operationname?status.svg
