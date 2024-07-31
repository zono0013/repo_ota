# Docker Hubから公式のGoイメージを使用
FROM golang:1.22-alpine

# 必要な依存関係をインストール
RUN apk add --no-cache git

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# go.modとgo.sumファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコンテナにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/main.go

# 実行権限を付与
RUN chmod +x main

# コンテナ実行時に実行されるコマンド
CMD ["./main"]
