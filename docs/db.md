# db

## PostgreSQL
- [春の入門祭り🌸 #7 作って学ぶGraphQL。gqlgenを用いて鉄道データ検索API開発入門](https://future-architect.github.io/articles/20200609/)

```bash
# start db container
docker-compose up -d db
docker-compose exec db bash -c "psql -U user -d db"
```

```postgres
# postgresql client

db=# \dt
         List of relations
 Schema |   Name   | Type  | Owner
--------+----------+-------+-------
 public | articles | table | user
(1 row)

db=# \d articles
                                        Table "public.articles"
   Column   |            Type             | Collation | Nullable |               Default
------------+-----------------------------+-----------+----------+--------------------------------------
 id         | integer                     |           | not null | nextval('articles_id_seq'::regclass)
 title      | character varying           |           | not null |
 url        | character varying           |           | not null |
 created_at | timestamp without time zone |           |          |
 updated_at | timestamp without time zone |           |          |
Indexes:
    "articles_pkey" PRIMARY KEY, btree (id)

db=# select * from articles;
 id |   title    |            url             |     created_at      |     updated_at
----+------------+----------------------------+---------------------+---------------------
  1 | helloworld | https://www.helloworld.com | 2014-11-17 15:02:36 | 2014-11-17 15:02:36
  2 | google     | https://www.google.com     | 2014-11-17 15:02:36 | 2014-11-17 15:02:36
  3 | yahoo      | https://www.yahoo.co.jp    | 2014-11-17 15:02:36 | 2014-11-17 15:02:36
(3 rows)

db=# \q
```
