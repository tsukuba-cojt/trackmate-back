# ベースイメージとして公式のGolangイメージを使用
FROM golang:1.24.3

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコンテナにコピー
COPY ./app .

# アプリケーションをビルド（バイナリを / に作成）
RUN go build -o /myapp

# 実行ファイルを起動
CMD ["/myapp"]
