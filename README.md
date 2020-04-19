# jetbook

## Database schema Settings

開発段階のため、main.go にて Schema 設定済み。
各 Table の日本語入力対応のため、以下 SQL を MySQL にて流し込むこと。

`ALTER TABLE users CONVERT TO CHARACTER SET utf8mb4`
`ALTER TABLE books CONVERT TO CHARACTER SET utf8mb4`
`ALTER TABLE rents CONVERT TO CHARACTER SET utf8mb4`
