# go-study

Go学習用リポジトリです。`go-template` をコピーして学習テーマごとにディレクトリを作成します。

## 使い方

1. テンプレートをコピー
   ```bash
   cp -R go-template 01-basics
   ```
2. 新しいディレクトリへ移動
   ```bash
   cd 01-basics
   ```
3. 起動
   ```bash
   docker compose up --build
   ```

`01-basics` の部分は学習テーマに合わせて変更してください。

## 推奨カリキュラム（順番）

1. `001-go-core-syntax`
2. `002-struct-interface-design`
3. `003-errors-and-testing`
4. `004-concurrency-basics`
5. `005-concurrency-patterns`
6. `006-context-cancellation-timeout`
7. `007-http-networking-internals`
8. `008-rest-api-advanced`
9. `009-sql-transactions-and-repository`
10. `010-observability-and-profiling`
11. `011-terraform-infra-basics`
12. `012-github-actions-cicd`
13. `013-rest-api-production`
14. `014-graphql-basics`
15. `015-grpc-basics`

各フォルダの`README.md`に、学習トピックと実装課題を記載しています。