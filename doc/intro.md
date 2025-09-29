# 项目目录结构

本文档采用 Markdown 展示项目的目录结构，并给出推荐的标准化结构，便于团队协作与后续维护。

## 当前目录结构

```
.
├── .gitignore
├── Jenkinsfile
├── Makefile
├── README.md
├── build/
│   └── build.sh
├── cfg/
│   ├── application.yaml
│   └── local.application.yaml
├── cmd/
│   ├── cron-cli/
│   └── xuanyuan/
│       └── main.go
├── deployments/
├── doc/
│   ├── context.md
│   ├── dbSession.md
│   ├── gin.md
│   ├── intro.md
│   ├── jwt.md
│   ├── mongo.md
│   ├── mysql.md
│   └── process.md
├── env.sh
├── go.mod
├── go.sum
├── internal/
│   ├── api/
│   │   ├── common/
│   │   └── controller/
│   ├── bootstrap/
│   │   ├── config.go
│   │   ├── cron.go
│   │   ├── db.go
│   │   ├── log.go
│   │   ├── redis.go
│   │   ├── router.go
│   │   └── validator.go
│   ├── common/
│   │   ├── context.go
│   │   ├── error.go
│   │   └── response.go
│   ├── config/
│   │   ├── app.go
│   │   ├── config.go
│   │   ├── database.go
│   │   ├── jwt.go
│   │   ├── log.go
│   │   ├── redis.go
│   │   └── storage.go
│   ├── dto/
│   │   ├── common.go
│   │   └── process.go
│   ├── middleware/
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   └── recovery.go
│   ├── model/
│   │   ├── common.go
│   │   ├── order.go
│   │   ├── package.go
│   │   ├── process.go
│   │   ├── recipes.go
│   │   └── user.go
│   ├── repository/
│   ├── routes/
│   │   ├── admin.go
│   │   ├── api.go
│   │   ├── cron.go
│   │   ├── recipes.go
│   │   └── user.go
│   ├── schedule/
│   │   ├── cron_schedule.go
│   │   ├── job.go
│   │   ├── jobs/
│   │   └── recipes.go
│   ├── service/
│   │   ├── account/
│   │   └── recipes/
│   └── util/
│       ├── bcrypt.go
│       ├── directory.go
│       ├── http.go
│       ├── md5.go
│       ├── random.go
│       ├── time.go
│       └── validator.go
├── migrations/
├── pkg/
│   ├── global/
│   │   ├── app.go
│   │   ├── error.go
│   │   └── lock.go
│   ├── health/
│   │   └── package.go
│   ├── main.go
│   ├── mapper/
│   │   ├── package.go
│   │   ├── process.go
│   │   ├── recipes.go
│   │   ├── repo.go
│   │   ├── user.go
│   │   └── user_repo.go
│   └── process/
│       └── interface.go
├── scripts/
│   └── build.sh
├── sql/
│   ├── createdb.sql
│   └── createtable.sql
└── test/
    └── unit-test.sh
```

## 推荐的标准结构

```
.
├── cmd/
│   ├── xuanyuan/              # Web 服务入口
│   │   └── main.go
│   └── cron-cli/              # CLI 工具入口
│       └── main.go
├── internal/                  # 业务实现（不对外暴露）
│   ├── bootstrap/             # 初始化：config/db/log/router/cron
│   ├── server/                # Gin Server 封装（可选）
│   ├── routes/                # 路由定义
│   ├── middleware/            # 中间件
│   ├── api/                   # 控制器/handler
│   ├── service/               # 领域服务
│   ├── repository/            # 数据访问层
│   ├── model/                 # 领域模型
│   ├── dto/                   # 请求/响应 DTO
│   ├── schedule/              # 定时任务
│   ├── common/                # 公共类型与错误
│   └── util/                  # 工具函数
├── pkg/                       # 对外可复用的库（少量）
│   └── health/
├── configs/                   # 配置文件
│   ├── application.yaml
│   └── application.local.yaml
├── migrations/                # DB 迁移/建表脚本
│   ├── 0001_init.sql
│   └── ...
├── scripts/                   # 构建/发布/测试脚本
│   └── build.sh
├── deployments/               # 部署文件（Docker/K8s 等）
├── docs/                      # 项目文档
├── Makefile
├── go.mod
├── go.sum
├── README.md
└── test/
    └── unit-test.sh
```

## 约定与说明

- 入口组织：采用 `cmd/<binary>/main.go` 作为各二进制入口。
- 模块边界：`internal` 存放业务实现，`pkg` 仅面向外部复用的库。
- 依赖方向：`api → service → repository → model`，避免环依赖。
- 命名统一：目录与包名小写、无下划线；`cron-cli` 使用连字符。
- 资源分类：配置 `configs/`，迁移 `migrations/`，脚本 `scripts/`，文档 `docs/`。