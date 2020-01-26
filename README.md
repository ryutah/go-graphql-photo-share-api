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

### 2. スキーマ定義更新

1. スキーマ定義を更新
1. モデルのコードを定義
    - スキーマの更新とモデルの更新が別になるのが少し気持ち悪いところ
1. コード生成
    - モデルに必要なフィールドなどがない状態でコードを生成すると、対応するresolverのインターフェースが生成されてしまうので注意

    ```console
    go run github.com/99designs/gqlgen generate
    ```

1. resolverを更新する
    - 未定義の関数をresolverに追加する感じ

### 3. enum, inputの追加

1. スキーマ定義を更新
1. コード生成
    - inputとenumを生成する

    ```console
    go run github.com/99designs/gqlgen generate
    ```

1. モデルの変更
    - 追加されたフィールドを追加

1. コード生成
    - 不要なresolver定義が生成されてしまうので、改めて更新

1. resolverを更新する

### 4. 1対多のリレーションを定義

1. スキーマ定義を更新
1. モデルの変更
1. コード生成

    ```console
    go run github.com/99designs/gqlgen generate
    ```

1. resolverを更新する
