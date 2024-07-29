# Wire Example

这是一个使用 Wire 进行依赖注入的 Go 示例项目。

## 项目结构
```text
wire-example/
├── main.go
├── wire.go
├── wire_gen.go
├── database/
│ └── database.go
├── service/
│ └── service.go
└── app/
└── app.go
```

## 步骤一：初始化项目

1. 创建项目目录：

    ```sh
    mkdir wire-example
    cd wire-example
    ```

2. 初始化 Go 模块：

    ```sh
    go mod init wire-example
    ```

3. 安装 Wire：

    ```sh
    go get github.com/google/wire/cmd/wire
    ```
4. 安装 Wire 命令行工具：

    ```sh
    go install github.com/google/wire/cmd/wire@latest
    ```

## 步骤二：定义各个组件

`database/database.go`

```go
package database

import "fmt"

type Database struct {
    DSN string
}

func NewDatabase(dsn string) (*Database, error) {
    // 模拟数据库连接
    fmt.Println("Connecting to database with DSN:", dsn)
    return &Database{DSN: dsn}, nil
}
```

`service/service.go`

```go
package service

import (
    "fmt"
    "wire-example/database"
)

type Service struct {
    DB *database.Database
}

func NewService(db *database.Database) *Service {
    return &Service{DB: db}
}

func (s *Service) Serve() {
    fmt.Println("Service is serving with database:", s.DB.DSN)
}
```

`app/app.go`
```go
package app

import (
    "wire-example/service"
)

type App struct {
    Service *service.Service
}

func NewApp(service *service.Service) *App {
    return &App{Service: service}
}

func (a *App) Run() {
    a.Service.Serve()
}
```

## 步骤三：设置 Wire 注入
`wire.go`
```go
// +build wireinject

package main

import (
    "github.com/google/wire"
    "wire-example/app"
    "wire-example/database"
    "wire-example/service"
)

func InitializeApp(dsn string) (*app.App, error) {
    wire.Build(
        database.NewDatabase,
        service.NewService,
        app.NewApp,
    )
    return &app.App{}, nil
}
```

## 步骤四：生成 Wire 注入代码
在项目根目录运行以下命令：
```shell
wire
```
这将生成 wire_gen.go 文件。

## 步骤五：主程序
`main.go`
```go
package main

import (
    "fmt"
)

func main() {
    app, err := InitializeApp("user:password@/dbname")
    if err != nil {
        fmt.Println("Failed to initialize app:", err)
        return
    }
    app.Run()
}
```
在项目根目录运行以下命令：
```shell
go run .
```
你应该会看到类似下面的输出：
```shell
Connecting to database with DSN: user:password@/dbname
Service is serving with database: user:password@/dbname
```

这个例子展示了如何使用 Wire 进行依赖注入，生成 wire_gen.go 文件，并在主程序中使用生成的初始化函数来创建和运行应用程序。

