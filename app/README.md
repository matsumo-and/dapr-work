# Dapr gRPC-Connect Microservices

Googleスタイルのモノレポ構成で、複数のgRPC-Connectマイクロサービスを管理します。

## 📁 ディレクトリ構造

```
app/
├── cmd/                    # サービスごとのエントリーポイント
│   ├── auth-service/
│   └── user-service/
├── proto/                  # 共通protoディレクトリ
│   ├── auth/v1/
│   └── user/v1/
├── internal/service/       # サービス実装
│   ├── auth/
│   └── user/
├── go.mod
├── buf.yaml
└── Makefile
```

## 🚀 セットアップ

### 1. Buf のインストール

Protoファイルからコード生成するために必要です。

```bash
# macOS
brew install bufbuild/buf/buf

# その他のOS
# https://buf.build/docs/installation を参照
```

### 2. 依存関係のインストール

```bash
go mod tidy
```

### 3. Protoからコード生成

```bash
make proto
```

これで以下のファイルが生成されます:
- `proto/auth/v1/*.pb.go`
- `proto/auth/v1/*connect.go`
- `proto/user/v1/*.pb.go`
- `proto/user/v1/*connect.go`

## 📝 使い方

```bash
# 各サービスを個別に起動
make run-auth    # auth-service (port 8080)
make run-user    # user-service (port 8081)

# ビルド
make build-auth
make build-user
make build-all

# Dockerイメージをビルド
make docker-build-auth
make docker-build-user
make docker-build-all

# サービス分離のテスト
make test-isolation  # サービス間の依存関係がないことを確認
```

## 🔧 サービス

### auth-service (port 8080)
- `Login` - ユーザーログイン
- `ValidateToken` - トークン検証
- `Logout` - ログアウト

### user-service (port 8081)
- `GetUser` - ユーザー情報取得
- `CreateUser` - ユーザー作成
- `UpdateUser` - ユーザー更新
- `DeleteUser` - ユーザー削除

## 🏗️ 開発

### 新しいサービスを追加する場合

1. `proto/SERVICE_NAME/v1/` にprotoファイルを作成
2. `internal/service/SERVICE_NAME/handler.go` を実装
3. `cmd/SERVICE_NAME/main.go` を作成
4. `make proto` でコード生成
5. Dockerfileを追加（必要に応じて）
