help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dbgen: ## sqlboilerによるコード自動生成
	# sqlboilerのインストール
	GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
	GO111MODULE=off go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql

	# DDL定義を元にコードを自動生成
	sqlboiler mysql

fmt: ## fmtの実行
	# goimportsのインストール
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

	# tidy, fmt, goimportsの実行
	go mod tidy -v
	gofmt -s -w pkg/
	goimports -w pkg/

run: ## APIをビルドせずに立ち上げるコマンド
	go run ./cmd

build: ## APIをビルドして立ち上げるコマンド
	go build -o binary/main ./cmd
	./binary/main