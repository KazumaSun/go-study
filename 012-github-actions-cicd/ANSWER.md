# 012 正解例（課題別）

## 課題1: PR時にtest/race実行

```yaml
name: ci

on:
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.26.2"
      - run: go test ./...
      - run: go test -race ./...
```

## 課題2: mainマージ時にビルド

```yaml
on:
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: docker build -t go-study-app:${{ github.sha }} .
```

## 課題3: 手動承認付きデプロイ

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    environment:
      name: production
    steps:
      - run: echo "deploy to production"
```

## 補足
- `environment: production` を作成し、GitHub側で required reviewers を設定すると手動承認フローになる
