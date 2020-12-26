# gqlgen

- [gqlgen](https://gqlgen.com/getting-started/)

```bash
gqlgen init

# Install go modules by selecting `Go:Install/Update Tools`

# implement the resolvers

go run server.go

# query
mutation createArticle {
  createArticle(input:{url:"https://www.helloworld.com", title: "hello world", sourceId:"1"}) {
    id,
    url,
    title,
    source {
      id,
      name,
    }
  }
}

query findArticle {
  articles {
    id,
    url,
    title,
    source {
      id,
      name,
    }
  }
}
```
