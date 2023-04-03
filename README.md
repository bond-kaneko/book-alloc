# book-alloc

## Overview

読書ポートフォリオマネジメントツールです。

どういうカテゴリの本を読んだのかではなく、自分にとってどんな意味を持つ本を読んだのか、を記録するサービスです。
<img width="1430" alt="スクリーンショット 2023-04-03 22 40 23" src="https://user-images.githubusercontent.com/20700893/229527123-9dbade20-2e97-4bc3-9502-59a40df113b0.png">

## Init

```
$ docker compose up -d
$ docker compose exec app sql-migrate up -dryrun -config=config/sql-migrate/dbconfig.yml -env=local
$ docker compose exec app sql-migrate up -config=config/sql-migrate/dbconfig.yml -env=local
```

## Dev

API側はGo、Web側はVueを使用しています。Auth0を利用して認証を行います。

### API
- Go 1.18
  - gin(Webフレームワーク)
  - air(ホットリロード)
  - gorm(ORM)
  - sql-migrate(DBマイグレーション)
- PostgreSQL 15.2

https://github.com/bond-kaneko/book-alloc

### Web 
- Vue 3.2.47
  - pinia 2.0.33
  - vue-router 4
  
https://github.com/bond-kaneko/book-alloc-web
