# lightning-go 

## 环境依赖

### 版本依赖

- Go 1.14+
- MySQL 5.7+
- Redis 3.2+


## 模块介绍
- [x] 多云管理 
- [ ] 消息中心 
- [ ] 定时任务| 任务调度 


## 项目目录结构
```bash
├── README.md
├── build.sh
├── cmd              # 项目入口
│   ├── migrate      # 创建表
│   └── server       # 启动服务
├── config           # 配置文件
│   ├── migrate.toml # 创建表配置文件
│   └── server.toml  # 启动服务配置文件
├── go-ops.service   # 服务systemd文件
├── go.mod
├── go.sum
├── internal         # 内部依赖
│   ├── app          # 应用
│   ├── cron         # 定时任务
│   ├── db           # 数据库
│   └── http         # http
├── log              # 日志目录
├── pkg
├── scripts
└── test             # 测试用例
```

## 部署

- 克隆代码
```bash
$ git clone git@github.com:zhengyansheng/lightning-go.git 
```

- 同步数据库
```bash
# make migrate
```

- 启动服务
```bash
# make run
```

- 生成 api docs
```bash
# make swagger
```

