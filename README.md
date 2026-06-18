# Go-Todo-API

Go + Echo + PostgreSQL を使用した、認証機能付きTodo管理APIサーバーです。
レイヤードアーキテクチャを意識し、保守しやすい構成を目指して開発しています。

## 🚀 特徴・機能
- **Todo管理:** TodoのCRUD（作成・取得・更新・削除）操作が可能。
- **認証機能:** JWT（JSON Web Token）を用いたセキュアなユーザー認証を実装。
- **レイヤー分離（Handler / Repository / Model）:** 責務の分離を意識したディレクトリ構成。

## 🛠 使用技術 (Tech Stack)
| カテゴリ | 技術 |
| :--- | :--- |
| **Language** | Go |
| **Framework** | Echo |
| **Database** | PostgreSQL |
| **Auth** | JWT (JSON Web Token) |

## 📂 ディレクトリ構成
```text
.
├── cmd/main.go      # エントリーポイント
├── internal/        # ビジネスロジック・レイヤー
│   ├── handler/     # HTTPリクエストのハンドリング
│   ├── model/       # データ構造の定義
│   └── repository/  # DB操作（PostgreSQL）
├── go.mod
└── ... 
```

## 今後の予定

- [ ☓ ] リフレッシュトークン
- [ ] ユニットテスト
- [ ] Docker対応
- [ ] バリデーション

## 🚀 サーバー起動
```bash
go run cmd/main.go
```

## 📌 API一覧

### 認証

- POST /register
- POST /login

### Todo

- GET /todos
- POST /todos
- PUT /todos/:id
- DELETE /todos/:id
