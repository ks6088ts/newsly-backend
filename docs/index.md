# setup logs

```bash
docker network create shared-network
cp .env.sample .env
docker-compose up -d
docker-compose exec go bash
go mod init

# cobra
make cobra-init
make cobra-add COBRA_CMD=hello
make build
./outputs/template-go --help
./outputs/template-go hello --help

# ファイルが windows から見えない場合の WA
# sudo chmod -R 777 ./

# Cross Compile in Golang
# https://medium.com/@georgeok/cross-compile-in-go-golang-9f0d1261ee26
make build \
    GOBUILD='GOOS=windows GOARCH=amd64 go build' \
    BIN_PATH='outputs/windows/cli.exe'

# マルチステージビルドで最小サイズイメージを作成
docker-compose -f docker-compose.release.yml build
docker-compose -f docker-compose.release.yml run --rm release ash -c "./cmd --help"

# tag を打ってリリース
git tag v0.0.0
git push origin v0.0.0

# GitHub Actions の仕様変更に追従
# https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-commands-for-github-actions#setting-an-environment-variable
# https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-commands-for-github-actions#adding-a-system-path

# version 情報をバイナリに埋め込む
# https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
# go build -ldflags="-X 'package_path.variable_name=new_value'"

# 自動的なバッファリングの確認:
# strace でシステムコールの発行回数を計測してみる
strace -e trace=write -c ./outputs/template-go

# リポジトリ構成(慣習に従う)
# https://medium.com/veltra-engineering/how-to-release-golang-tools-with-go-get-1c856739f5e3
# github.com/<owner>/<repository>
# github.com/<owner>/<repository>/cmd/<repository>

# template リポジトリ利用時、リポジトリ名をリネームする
make rename REPO_NAME=example-go
```
