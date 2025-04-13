# AuthKit

認証方法のテンプレートをまとめたリポジトリ

[要件Issue](https://github.com/mattuu0/authkit/issues/1)

# 環境構築
- コンテナ構成
  - Nginx
    - リバースプロキシ
    - Webサーバ
    - 各種ログとり
    - port: 8370
  - app
    - URL: /app/
    - メインの動作をするアプリ
    - 基本的に何でも良い
  - Auth
    - URL: /auth/
    - 認証コンテナ
    - アプリと連携する
  - Mysql
    - データベースコンテナ
    - 基本的には認証コンテナ
  - sslコンテナ
    - Nginx 用のSSL証明書を生成するコンテナ
    - 最初に起動して証明書を作成して終わる

## 各種リンク
- ベースURL: https://localhost:8370