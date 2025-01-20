# AIHireSphere
> AI驱动的助力面试平台

## 项目简介
使用[Go-Zero](https://github.com/zeromicro/go-zero)框架开发的AI驱动的助力面试平台

## 快速开始
### · docker启动
#### 1. 到deploy目录下面
```shell
  cd deploy
```
#### 2. 执行docker build 命令
```shell
  docker build -f Dockerfile -t repository:tag .  
```
#### 3. 运行docker容器
```shell
  docker run -d -p 8080:8080 --name test repository:tag
```
如果您看到``Starting server at 0.0.0.0:8888...`` 字样视为启动成功

####  ！特别注意
本项目不含最基础配置文件，对于每个微服务，我们规定：配置文件模板放在
``application/application/xxx-center/interfaces/xxx/etc/xx.yaml.templeate`` 文件中。 

若想正常启动微服务，需要根据配置文件模板``xxx.yaml.templeate``修改配置文件，之后把配置文件命名为 ``xx.yaml``到原来目录下，再启动``docker build`` 和 ``docker run``。

或者是根据配置文件模板``xxx.yaml.templeate``修改配置文件，之后把配置文件命名为 ``xx.yaml``,再通过 ``docker run -v your/config_yaml_path:/app/etc`` 进行挂载，确保程序正常启动。





## 项目技术栈
1. Go-Zero
2. Gorm
3. Mysql
4. Redis
5. zap
6. ...

## 项目目录结构
本项目遵循ddd项目架构 基于go-zero框架搭建 微服务应用
```text
.
├── Dockerfile
├── LICENSE
├── README.md
├── README_ZH.md
├── application         //微服务应用层
│   └── user-center     // 用户中心
├── common              // 公共模块
│   ├── call            // rpc调用类
│   ├── codex           // 响应包装
│   ├── decorator       // CQRS装饰器，本阶段暂未使用
│   ├── gormx           // gorm操作类
│   ├── interceptors    // 拦截器
│   ├── model           // 数据库表类
│   ├── utils           // 工具类
│   ├── zapx            // zap日志类
│   └── zlog            // 日志服务
├── docker-compose.yaml         // 一键环境安装
├── docs                        // 文档
├── go.mod
├── go.sum
├── logs
├── scripts                 // 脚本类
├── template                // go-zero模板 后续移植github公共仓库
├── test                    // 测试类
└── utils   
    └── path.go             
```
对于每个微服务 我们的目录结构为
```text

.
├── README.md          //微服务说明
├── app                //应用层
│   └── user.go  // 应用
├── domain          // 领域层
│   ├── events      // 领域事件
│   ├── irepository // 仓储层接口
│   ├── model       // 领域模型
│   └── services    // 领域服务
├── infrastructure      // 基础设施层
│   ├── driver          // 驱动层
│   └── repository      // 仓储层
└── interfaces          // 接口层
    ├── api             // API层
    └── rpc             // RPC层



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
<type>/<name>/<description>
```
例如：
- 如果是开发新功能的分支，命名规范如下
```text
feature/<name>/<feature description>
例如：feature/jett/dev_log_system
```

- 如果是修复bug：
```text
bugfix/<name>/<bug name>
e.g.：bugfix/jett/login_error
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
### 项目一期
- [ ] 用户中心
  - [ ] 用户登陆
  - [ ] 用户注册
  - [ ] 用户搜索
- [ ] 面试中心
  - [ ] 简历上传
  - [ ] 简历搜索
  - [ ] 简历筛选
  - [ ] AI面试(纯文本形式)
  - [ ] 题库展示
  - [ ] 格式化批量上传题库
- [ ] AI调研

### 项目二期
- [ ] 用户中心
  - [ ] 用户积分
  - [ ] 积分消耗
  - [ ] 积分购买
- [ ] 面试中心
  - [ ] AI面试(语音形式/视频形式)