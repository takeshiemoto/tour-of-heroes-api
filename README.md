# tour-of-heroes-api

## 事前準備

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [Herokuアカウント](https://signup.heroku.com/)
- [Heroku CLI](https://devcenter.heroku.com/ja/articles/heroku-cli)

## Heroku CLI

ログイン

```
heroku login
```

[Herokuへのログイン](https://devcenter.heroku.com/ja/articles/heroku-cli#getting-started)

## Herokuのセットアップ

Git Clone

```
git clone git@github.com:takeshiemoto/tour-of-heroes-api.git
cd tour-of-heroes-api
```

Herokuにアプリケーションを作成する

```
heroku create
```

Herokuにデータベースを作成

```
heroku addons:create heroku-postgresql:hobby-dev
```

データベースへの接続情報を確認

```
heroku config
```

接続情報参考にデータベースにログイン後テーブルを作成する

```postgresql
create table heroes (
    id serial primary key ,
    name varchar(255),
    createdAt timestamp not null default current_timestamp,
    updateAt timestamp not null default current_timestamp
)
```

アプリケーションをデプロイ

```
git push heroku main
```

アプリケーションを開く
```
heroku open
```

ログの表示

```
heroku logs --tail 
```

## ローカル開発環境のセットアップ

Postgres用のDockerを立ち上げる
```
docker-compose up
```

Docker環境にログイン

```
docker exec -i -t tour-of-heroes-api_postgres_1 bash
```

postgresにログイン

```
su - postgres
psql
```

ロールとユーザーを作成

```
CREATE ROLE toh LOGIN CREATEDB PASSWORD 'toh';
CREATE DATABASE toh OWNER toh;
```

作成したユーザーでログイン

```
psql --username=toh --password --dbname=toh
```

テーブルを作成

```postgresql
create table heroes (
    id serial primary key ,
    name varchar(255),
    createdAt timestamp not null default current_timestamp,
    updateAt timestamp not null default current_timestamp
)
```

## ローカル環境でAPIを起動

アプリケーションをビルド

```
go build -o bin/tour-of-heroes-api -v .
```

ローカル開発環境を起動

```
heroku local web
```