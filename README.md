# RSS Aggregator: Go, PostgreSQL

Key Features:
- REST API 
- Authentication using api keys and middleware
- RSS feed web scraper
- PostgreSQL database


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
