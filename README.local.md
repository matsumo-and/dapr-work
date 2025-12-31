# ローカル開発環境

Dapr CLIとDocker Composeを使ったローカル開発環境のセットアップガイドです。

## 📦 前提条件

以下のツールをインストールしてください:

- [Docker](https://docs.docker.com/get-docker/)
- [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
- [Go 1.23+](https://go.dev/dl/)
- [Buf](https://buf.build/docs/installation)

## 🚀 クイックスタート

```bash
# 1. アプリケーションディレクトリに移動
cd app

# 2. アプリケーション用インフラを起動（PostgreSQL）
make docker-up

# 3. Daprを初期化（初回のみ - Redis, Zipkin, Placementが自動セットアップ）
make dapr-init

# 4. auth-serviceをDaprで起動（別ターミナル）
make dapr-run-auth

# 5. user-serviceをDaprで起動（さらに別ターミナル）
make dapr-run-user
```

これで以下が起動します:
- ✅ PostgreSQL (port 5432) - アプリケーション用DB
- ✅ Redis (port 6379) - Daprが自動作成 (state store / pub-sub)
- ✅ Zipkin (port 9411) - Daprが自動作成 (分散トレーシング)
- ✅ Placement (port 50005) - Daprが自動作成 (actor用)
- ✅ auth-service (port 8080) + Daprサイドカー (3500)
- ✅ user-service (port 8081) + Daprサイドカー (3501)

**メリット:**
- 🔥 **ホットリロード** - コード変更が即座に反映
- 🐛 **デバッグ可能** - IDEでブレークポイント設定可能
- 🚀 **高速起動** - Docker不要、Goの高速コンパイル
- 📝 **ログが見やすい** - Dapr CLIが整形して表示

## 📦 サービス構成

### PostgreSQL
- Host: `localhost`
- Port: `5432`
- Database: `appdb`
- Username: `app`
- Password: `app123`

接続方法:
```bash
# アプリケーションディレクトリから
cd app
make db

# または直接
psql -h localhost -p 5432 -U app -d appdb
```

### Redis (Daprが自動作成)
- Host: `localhost`
- Port: `6379`
- 用途: Dapr state store / pub-sub

接続方法:
```bash
# アプリケーションディレクトリから
cd app
make redis

# または直接
docker exec -it dapr_redis redis-cli
```

### auth-service
- HTTP: `http://localhost:8080`
- Dapr HTTP: `http://localhost:3500`
- Dapr gRPC: `50001`

### user-service
- HTTP: `http://localhost:8081`
- Dapr HTTP: `http://localhost:3501`
- Dapr gRPC: `50002`

## 🛠️ よく使うコマンド

**注意:** すべてのコマンドは `app/` ディレクトリから実行してください

```bash
# インフラ操作
make docker-up           # アプリケーションインフラ起動 (PostgreSQL)
make docker-down         # インフラ停止
make docker-logs         # インフラのログ
make docker-logs-db      # PostgreSQLのログ
make docker-clean        # データを含めて完全削除

# Dapr
make dapr-init           # Dapr初期化 (Redis, Zipkin, Placement自動セットアップ)
make dapr-run-auth       # auth-serviceをDaprで起動
make dapr-run-user       # user-serviceをDaprで起動

# データベース接続
make db                  # PostgreSQL接続
make redis               # Dapr Redis接続

# 開発
make proto               # protoファイルからコード生成
make build-auth          # auth-service バイナリビルド
make build-user          # user-service バイナリビルド
make test-isolation      # サービス分離テスト

# その他
make help                # ヘルプ表示
```

## 🧪 動作確認

### ヘルスチェック

```bash
# auth-service
curl http://localhost:8080/health

# user-service
curl http://localhost:8081/health
```

### gRPC-Connect経由でAPI呼び出し

```bash
# auth-service: Login
curl -X POST http://localhost:8080/auth.v1.AuthService/Login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'

# user-service: GetUser
curl http://localhost:8081/user.v1.UserService/GetUser \
  -H "Content-Type: application/json" \
  -d '{"id":"00000000-0000-0000-0000-000000000001"}'
```

### Dapr経由でサービス呼び出し

```bash
# Dapr HTTPエンドポイント経由
curl http://localhost:3500/v1.0/invoke/user-service/method/user.v1.UserService/GetUser \
  -H "Content-Type: application/json" \
  -d '{"id":"00000000-0000-0000-0000-000000000001"}'
```

## 📁 ファイル構成

```
app/
├── docker-compose.yml           # アプリケーションインフラ (PostgreSQL)
├── Makefile                     # 統合されたすべてのコマンド
├── .dapr/
│   └── components/             # Dapr コンポーネント定義
│       ├── statestore.yaml     # Redis state store (dapr initで作成されたRedis使用)
│       └── pubsub.yaml         # Redis pub/sub (dapr initで作成されたRedis使用)
├── scripts/
│   ├── init-db/                # DB初期化スクリプト
│   │   └── 01-init.sql         # テーブル作成・サンプルデータ
│   └── check-isolation.sh      # サービス分離テストスクリプト
├── proto/                      # Protocol Buffers定義
├── cmd/                        # サービスエントリポイント
└── internal/                   # サービス実装

※ Dapr関連インフラ(Redis, Zipkin, Placement)は `dapr init` で自動作成されます
```

## 🔧 トラブルシューティング

### ポートが既に使用されている

```bash
# 使用中のポートを確認
lsof -i :5432
lsof -i :8080

# プロセスを停止してから再起動
cd app
make docker-down
make docker-up
```

### データベースをリセットしたい

```bash
cd app

# ボリュームを含めて削除
make docker-clean

# 再起動
make docker-up
```

### Daprサイドカーが起動しない

```bash
# Dapr placementが起動しているか確認
docker ps | grep dapr

# Dapr再初期化
dapr uninstall
make dapr-init
```

## 🎯 次のステップ

- [Kubernetes環境へのデプロイ](../README.md#5-基本アドオンのセットアップ)
- [アプリケーション開発ガイド](README.md)
- [Daprコンポーネント設定](.dapr/components/)
