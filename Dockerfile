# ベースとなるDockerイメージ指定
FROM golang:1.17-alpine
# コンテナ内に作業ディレクトリを作成
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src