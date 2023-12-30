# RSS Aggregator: Go, PostgresSQL

Key Features:
- REST api 
- Authentication using api keys
- RSS feed web scraper
- PostgresSQL database


### Cloning the repository

```shell
git clone https://github.com/EthanHosier/rssagg.git
```

### Setup .env file


```js
PORT=
DB_URL=postgresql://{username}:{password}@{host}:{port}/{database_name}

```

### Start the app

```shell
go build && ./rssagg
```
