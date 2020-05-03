# 【おうちでハッカソン】BarcelonaZoo-api

## セットアップ
### 【sqlboiler】
1. sqlboiler_example.tomlをコピーして、同じ階層にsqlboiler.tomlを作成する。
2. 自分の開発環境(localhost)に合わせて `user`と`pass`を編集する
例）user: `root` / pass: `password`の場合

```toml
pkgname="model"
output="pkg/model"
wipe=true
[mysql]
  dbname = "barcelona_zoo"
  host   = "localhost"
  port   = 3306
  user   = "root"
  pass   = "password"
  sslmode= "false"
```

### 【環境変数】
1. .env_exampleをコピーして、同じ階層に.envを作成する
2. 自分の開発環境(localhost)に合わせて`MYSQL_USER`と`MYSQL_PASSWORD`を修正

```.env
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_HOST=localhost
```

## 使用できるMakeコマンド一覧
```shell script
help  使い方
dbgen sqlboilerによるコード自動生成
fmt   fmtの実行
run   APIをビルドせずに立ち上げるコマンド
build APIをビルドして立ち上げるコマンド
```

## ブランチルール
feat/[issue番号]

## DB設計
### ChallengeTheme - チャレンジネタ
* id - int
* title - string
* description - string
* user_id - int
* recordable - boolean
* is_int - boolean
* unit - string(単位)
* ranking_type - int
    * 例）
        * 1: 昇順
        * 2: 降順
        * 3: いいね数
* created_at - datetime

### ChallengeRecord - チャレンジ記録
* id - int
* content - string(url)
* comment - string
* challenge_theme_id - int
* user_id - int
* record - float
* created_at - datetime

### User - ユーザ情報（サロゲートとしての役割のみ）
* id - int
* uid - string

### Favorite - お気に入り登録しているチャレンジ
* id - int
* user_id - int
* challenge_theme_id - int

### Tag - タグ一覧
* id - int
* text - string

### ThemeTagRelation - チャレンジネタについているタグ
* id - int
* challenge_theme_id - int
* tag_id - int

