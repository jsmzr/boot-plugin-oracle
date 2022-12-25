# boot-plugin-oracle

[![Build Status](https://github.com/jsmzr/boot-plugin-oracle/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-plugin-oracle/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-plugin-oracle/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-plugin-oracle)

基于 go-ora 的 boot 系列 oracle 插件

## 使用说明

1. 拉取依赖，`go get -u github.com/jsmzr/boot-plugin-oracle`
2. 同其他插件一样，需要在入口 `main.go` 中显式导入以完成插件注册、初始化
3. 通过 `github.com/jsmzr/boot-plugin-oracle/connection` 来使用连接

```go
// main.go
import _ "github.com/jsmzr/boot-plugin-oracle"

// db.go
import "github.com/jsmzr/boot-plugin-oracle/connection"

func PrePare(statement string) (*sql.Stmt, error) {
    return connection.DB().PrePare(statement)
}
```