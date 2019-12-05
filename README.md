## graphql-go

A simple books and authors store using graphql for interaction.

#### To compile & run
* Checkout source code
```sh
$ git clone https://github.com/sauravgsh16/graphql-go.git
```

* Change in project directory
```sh
$ cd graphql-go
```

* Execute docker-compose
```sh
$ docker-compose up --build
```

#### GraphQL API requests eg.
```json

query {
    Book (id: 1) {
        id
        name
    }
}

```
```json
mutation {
    createAuthor (name: "foo", age: 42) {
        id
        name
    }
}

```



