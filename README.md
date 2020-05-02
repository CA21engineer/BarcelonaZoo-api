# 【おうちでハッカソン】BarcelonaZoo-api

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

