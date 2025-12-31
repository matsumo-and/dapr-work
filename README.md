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
- **Dapr** - Daprランタイム
- **Prometheus + Grafana** - メトリクス監視と可視化

#### アクセス情報

- Grafana: http://localhost:30030 (admin/admin)
- Prometheus: http://localhost:30090

## アプリケーション開発

アプリケーションのセットアップと開発方法については [app/README.md](app/README.md) を参照してください。