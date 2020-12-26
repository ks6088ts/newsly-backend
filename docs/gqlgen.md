# gqlgen

- [gqlgen](https://gqlgen.com/getting-started/)
- [開発ライブ実況 #2 GoLand編 (メルペイSolutionsチーム デフォルト好きエンジニア)](https://www.youtube.com/watch?v=8MdxqDb07eQ)

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
