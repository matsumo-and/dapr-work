# dapr-work
A demo repository for Dapr with multiple Go gRPC-Connect services running on a Kind cluster.

## Get Started

### 1. kubectl のインストール

公式ページを参照してください: https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/

### 2. Helm のインストール

公式ページを参照してください: https://helm.sh/ja/docs/intro/install/

### 3. kind のインストール

公式ページを参照してください: https://kind.sigs.k8s.io/docs/user/quick-start/#installation

### 4. kind でクラスターを作成

```bash
kind create cluster --config cluster/cluster.yaml --name kindcluster
```