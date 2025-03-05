# タスク管理アプリ

LayerX 様 インターンの課題です

## 設計
こちらを参照ください

https://github.com/yudai2929/task-app/blob/main/DESIGN.md

## 環境

| **項目**           | **バージョン / 詳細**   |
| ------------------ | ----------------------- |
| **Go**             | `1.22`                  |
| **Docker Compose** | `3.8`                   |
| **PostgreSQL**     | `15`                    |
| **Swagger UI**     | `swaggerapi/swagger-ui` |

## 起動方法

### Docker コンテナの起動

```
make up
```

- API サーバー (8080)
- PostgreSQL (5432)
- Swagger UI (8081)

### API サーバーの起動 (オプション)

docker を使わずに go のサーバーを立ち上げる場合

```
make run_server
```

### PostgreSQL に接続

```
make psql
```

### UT 実行

```
make test
```

## 環境セットアップ

```
make set_up
```

- 依存パッケージの取得 (go mod tidy)

```
make install
```

- `ogen` (OpenAPI 生成ツール)
- `xo` (DB スキーマ生成ツール)

```
make gen
```

- `go generate` によるコード自動生成

## Swagger UI (API ドキュメント)

Swagger UI は以下の URL からアクセス可能：

```
http://localhost:8081
```

## 所要時間
計12時間程度
- 設計 1.5 時間
- ドキュメント 1 時間
- 環境整備 1.5 時間
- 実装 8 時間
