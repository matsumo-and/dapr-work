# dapr-work
A demo repository for Dapr with multiple Go gRPC-Connect services running on a Kind cluster.

## Get Started

### 1. kubectl のインストール

公式ページを参照してください: https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/

### 2. Helm のインストール

公式ページを参照してください: https://helm.sh/ja/docs/intro/install/

### 2.5. Helmfile のインストール

公式ページを参照してください: https://helmfile.readthedocs.io/en/latest/#installation

### 3. kind のインストール

公式ページを参照してください: https://kind.sigs.k8s.io/docs/user/quick-start/#installation

### 4. kind でクラスターを作成

```bash
kind create cluster --config cluster/cluster.yaml --name kindcluster
```

### 5. 基本アドオンのセットアップ

Helmfileを使って一括でインストールします。

```bash
cd manifests
helmfile apply
```

これにより以下がインストールされます:
- **NGINX Ingress Controller** - 外部アクセス用（ALB相当）
- **Metrics Server** - リソース監視用
- **Dapr** - Daprランタイム（Zipkinによる分散トレーシング機能を含む）
- **PostgreSQL** - データベース（開発用）

#### アクセス情報

- PostgreSQL: localhost:30432
  - Database: `appdb`
  - Username: `app`
  - Password: `app123`
  - Admin user: `postgres` / `postgres`

> **Note:** 分散トレーシングにはDaprが自動セットアップするZipkinを使用します。
> 長期的なメトリクス監視が必要な場合は、Prometheus/Grafanaを追加してください。

## ローカル開発環境

Dapr CLIとDocker Composeを使ったローカル開発については [README.local.md](README.local.md) を参照してください。

```bash
# アプリケーションディレクトリに移動
cd app

# アプリケーションインフラ起動（PostgreSQL）
make docker-up

# Dapr初期化（初回のみ - Redis, Zipkin, Placement自動セットアップ）
make dapr-init

# サービス起動
make dapr-run-auth
make dapr-run-user
```

## アプリケーション開発

アプリケーションのセットアップと開発方法については [app/README.md](app/README.md) を参照してください。