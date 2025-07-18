# Trackmate-Back

## 概要
Trackmate-Backは、家計簿・借金管理などの機能を持つGo製バックエンドAPIサーバーです。ユーザーごとに支出・カテゴリ・借金・借金相手などを管理できます。

## 主な機能
- ユーザー認証（サインアップ・ログイン）
- 支出管理（登録・一覧取得）
- 支出カテゴリ管理（登録・集計・削除）
- 借金管理（登録・集計・削除）
- 借金相手管理

## セットアップ手順
1. **リポジトリをクローン**
   ```bash
   git clone <このリポジトリのURL>
   cd trackmate-back
   ```
2. **依存パッケージのインストール**
   ```bash
   go mod tidy
   ```
3. **DBのセットアップ**
   - `docker-compose up -d` でDBを起動
   - 必要に応じて `migrations/` 配下のマイグレーションを実行
   ```
   go run migration/migration.go
   ```
4. **環境変数の設定**
   - `.env` ファイルを作成し、DB接続情報やJWTシークレットなどを記載
5. **サーバーの起動**
   ```bash
   go run .
   ```

## ディレクトリ構成
```
├── controllers/   # ルーティング・APIエンドポイント
├── services/      # ビジネスロジック
├── repositories/  # DBアクセス
├── models/        # DBモデル
├── dto/           # DTO（データ転送オブジェクト）
├── migrations/    # マイグレーション
├── infra/         # DB初期化など
├── main.go        # エントリポイント
```

## 主なAPI例

### 認証
- POST `/auth/signup`  ユーザー登録
- POST `/auth/login`   ログイン

### 支出
- GET `/expenses`      支出一覧取得
- POST `/expenses`     支出登録

### 支出カテゴリ
- GET `/categories`    カテゴリ集計取得
- POST `/categories`   カテゴリ登録
- DELETE `/categories/:category_id` カテゴリ削除

### 借金
- GET `/loan`          借金集計取得
- POST `/loan`         借金登録

### 借金相手
- GET `/person`        借金相手一覧
- POST `/person`       借金相手登録

## 注意事項
- API利用時は認証トークンが必要です。
- DBスキーマやAPI仕様は今後変更される可能性があります。

---

ご不明点・バグ報告はIssueまたは開発者までご連絡ください。
