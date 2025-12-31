-- 初期データベーススキーマ

-- Users テーブル
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Auth Tokens テーブル
CREATE TABLE IF NOT EXISTS auth_tokens (
    token VARCHAR(255) PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- インデックス
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_auth_tokens_user_id ON auth_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_auth_tokens_expires_at ON auth_tokens(expires_at);

-- サンプルデータ（開発用）
INSERT INTO users (id, username, email, password_hash)
VALUES
    ('00000000-0000-0000-0000-000000000001', 'admin', 'admin@example.com', '$2a$10$dummy_hash_for_development'),
    ('00000000-0000-0000-0000-000000000002', 'user1', 'user1@example.com', '$2a$10$dummy_hash_for_development')
ON CONFLICT (username) DO NOTHING;
