#!/bin/bash
set -e

echo "🔍 Checking service isolation..."
echo ""

# auth-service が user-service に依存していないかチェック
echo "Checking auth-service dependencies..."
AUTH_DEPS=$(go list -f '{{join .Deps "\n"}}' ./cmd/auth-service 2>/dev/null)
if echo "$AUTH_DEPS" | grep -q "github.com/matsumo_and/dapr-work/app/internal/service/user"; then
    echo "❌ FAIL: auth-service depends on user-service"
    exit 1
fi
echo "✅ PASS: auth-service has no dependency on user-service"
echo ""

# user-service が auth-service に依存していないかチェック
echo "Checking user-service dependencies..."
USER_DEPS=$(go list -f '{{join .Deps "\n"}}' ./cmd/user-service 2>/dev/null)
if echo "$USER_DEPS" | grep -q "github.com/matsumo_and/dapr-work/app/internal/service/auth"; then
    echo "❌ FAIL: user-service depends on auth-service"
    exit 1
fi
echo "✅ PASS: user-service has no dependency on auth-service"
echo ""

# proto依存関係のクロスチェック
echo "Checking proto dependencies..."
if echo "$AUTH_DEPS" | grep -q "github.com/matsumo_and/dapr-work/app/proto/user"; then
    echo "❌ FAIL: auth-service imports user proto"
    exit 1
fi
if echo "$USER_DEPS" | grep -q "github.com/matsumo_and/dapr-work/app/proto/auth"; then
    echo "❌ FAIL: user-service imports auth proto"
    exit 1
fi
echo "✅ PASS: No cross-proto dependencies"
echo ""

echo "✅ All service isolation checks passed!"
