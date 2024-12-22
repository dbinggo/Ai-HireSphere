# AIHireSphere
> AI驱动的助力面试平台

## 项目简介
使用[Go-Zero](https://github.com/zeromicro/go-zero)框架开发的AI驱动的助力面试平台

## 项目技术栈
1. Go-Zero
2. Gorm
3. Mysql
4. Redis
5. zap
6. ...

## 项目目录结构
本项目遵
```text
.
├── Dockerfile          # Dockerfile
├── LICENSE             # License
├── README.md           # README   
├── app                 # Application Layer
├── common              # Common Layer
│   ├── interceptors    # Interceptor
│   ├── model           # Model
│   ├── xcode           
│   ├── xgorm
│   ├── xzap
│   └── zlog
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   └── repository
├── logs
│   └── foo
├── scripts
│   └── goctl.sh
├── temp.txt
├── template
│   ├── api
│   ├── docker
│   ├── gateway
│   ├── kube
│   ├── model
│   ├── mongo
│   ├── newapi
│   └── rpc
├── test
│   ├── gorm_test.go
│   └── zlog_test.go
└── utils
    └── path.go
```
## 开发之前

### 1. **使能git钩子**
```shell
git config core.hooksPath .githooks
chmod -R -x .githooks
```


### 2. **阅读以下开发规范**

分支命名规范
我们要求：

1. 分支命名应包含负责人的名字。

2. 分支命名必须清晰地表达分支正在处理的问题。

因此，分支命名必须标准化。
```text
<type>-<name>-<description>
```
例如：
- 如果是开发新功能的分支，命名规范如下
```text
feature-<name>-<feature description>
例如：feature-jett-dev_log_system
```

- 如果是修复bug：
```text
bugfix-<name>-<bug name>
e.g.：bugfix-jett-login_error
```
以及其他类型：
`hotfix`、`release`...


### 提交信息格式
提交信息应尽可能清晰，每次提交只做一件事。

```text
<type>(<scope>): <subject>

e.g.：feat: add new api
or：feat(common): add new api
```

### 类型

```text
# 主要类型
feat:      添加新功能
fix:       修复bug

# 特殊类型
docs:      仅更改文档相关内容
style:     不影响代码含义的更改，如删除空格、改变缩进、添加或删除分号
build:     更改构建工具或外部依赖，如webpack、npm
refactor:  重构代码时使用
revert:    执行git revert时打印的消息

# 暂不使用的类型
test:      添加或修改现有测试
perf:      改善性能的更改
ci:        与CI（持续集成服务）相关的更改
chore:     不修改src或测试的其他修改，如更改构建过程或辅助工具
```

### 主题

末尾不加句号或标点

例如：
```text
feat: add new feature
fix: fix a bug
```



## **你必须知道**
1. 项目遵循[uber的开发规范](https://github.com/xxjwxc/uber_go_guide_cn)。
1. **不要**在任何代码中提交任何敏感信息，例如`api_key`、`address`或`password`。
2. 你可以使用配置文件`config.yaml`来存储一些敏感信息，但不要尝试提交它。每次修改`config.yaml`的结构时，你也必须更新`config.yaml.template`。
3. 除非你知道自己在做什么，否则不要使用`git push --force`。

## Todo List
- [ ] Project Initialization
    - [ ] Project Initialization
    - [ ] Project Directory Structure
    - [ ] Project Configuration File
    - [ ] Project Global Variables and Constants
    - [ ] Project Log Configuration
    - [ ] Project Initialization File
- [ ] Gin Module Setup
    - [ ] Gin Framework Setup
    - [ ] Gin Routing Setup
    - [ ] Gin Middleware Setup
    - [ ] Gin Parameter Binding
    - [ ] Gin Response Data Encapsulation
- [ ] Login and Registration Module
    - [ ] Login and Logout functionality
    - [ ] Auto-login within 30 days
    - [ ] Force logout
    - [ ] Display commonly used devices
- [ ] Team Homepage Module
    - [ ] Personal Center
    - [ ] Points Overview
    - [ ] Message Module
    - [ ] Feishu Multi-dimensional Table
- [ ] Team Information Module
    - [ ] Like personal information
    - [ ] User list
    - [ ] Add new user
    - [ ] Admin team structure management
    - [ ] View and edit user information details