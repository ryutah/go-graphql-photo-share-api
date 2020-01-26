# Photo Share API

## 実装メモ

### 1. 初期化

1. スキーマ定義を作成
1. `gqlgen.yml` を定義
1. コードの生成

    ```console
    go run github.com/99designs/gqlgen generate
    ```

1. resolverを実装する
1. mainパッケージを定義
