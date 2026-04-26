# 012-github-actions-cicd

GitHub ActionsでGo開発のCI/CDを実務レベルで構築するステップです。

## ゴール
- CIでlint/test/buildを自動化できる
- CDの基本戦略（自動/手動承認）を説明できる
- セキュアなシークレット運用を実装できる

## 学習トピック
- workflow / job / step / matrix
- キャッシュ戦略（Go module cache）
- PRチェック（lint、unit test、race test）
- リリースフロー（tag、artifact、deploy）
- OIDC連携の概念（クラウド認証）

## 実装課題（自分で実装）
1. PR時に`go test`と`go test -race`を実行
2. mainマージ時にビルドとイメージ作成
3. 手動承認付きデプロイワークフローを作る

## 自分への確認質問
- なぜ`-race`をCIで回す価値がある？
- matrix戦略はいつ有効？
- 秘密情報をworkflowに直書きしてはいけない理由は？

## 完了条件
- CI/CD全体を図解なしで口頭説明できる
