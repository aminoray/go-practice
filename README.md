## go-handson

golang練習用のリポジトリになります。

## 準備

```shell
gh repo clone github.com/FJC-OMUSUBI/go-handson
```

## 実行

### 1. [Effective Go](http://go.shibu.jp/)

- cmd/tutorial

```shell

go run cmd/tutorial/main.go
```

### 2. [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md)

- cmd/standard

```shell
go run cmd/standard/main.go
```

### 3. DB

- [Pgx](https://github.com/jackc/pgx)
  - cmd/pgx
- [Bun](https://bun.uptrace.dev/)
  - cmd/bun

```shell
準備中
```

### 4. HTTPサーバー

- [net/http mux](https://pkg.go.dev/net/http)
  - cmd/muxpgx
- [Gin](https://gin-gonic.com/ja/docs/)
  - cmd/ginpgx

```shell
準備中
```

### 5. CI/CD (Continuous Integration/Continuous Delivery)

#### CI

- [GitHub Actions](https://docs.github.com/ja/actions)
- [GitHub Actions - composite action](https://docs.github.com/ja/actions/creating-actions/creating-a-composite-action)
- [Buf](https://docs.buf.build/ci-cd/github-actions)
- [golangci-lint](https://golangci-lint.run/)
- [govulncheck](https://github.com/golang/vuln)

#### CD

- [Semantic Versioning Release Action](https://github.com/marketplace/actions/create-new-semantic-version)
- [go-winres](https://github.com/tc-hib/go-winres)
- [UPX](https://upx.github.io/)
- [osslsigncode](https://github.com/mtrojnar/osslsigncode)
- [GoReleaser](https://goreleaser.com/)
- [GitHub Release](https://docs.github.com/ja/repositories/releasing-projects-on-github)
- [ko](https://ko.build/)
- [GitHub Packages](https://docs.github.com/ja/packages)
- [Google Artifact Registry](https://cloud.google.com/artifact-registry?hl=ja)
- [Google Cloud Run](https://cloud.google.com/run?hl=ja)
