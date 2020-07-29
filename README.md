# Gingo

Gin Web 项目脚手架

## 项目结构

参考 [Standard Go Project Layout](https://github.com/golang-standards/project-layout) 目录结构

- api: 接口
- cmd： 程序启动入口
- configs： 配置文件
- deployments： 部署文件
- initialize： 初始化
- internal： 内部包
- middleware： 中间件
- model： 模型
- router： 路由
- service： 服务层
- test： 测试包
- third_party： 三方件
- tools： 工具包
- vendor： 依赖包

## RBAC 权限

测试： [casbin editor](https://casbin.org/zh-CN/editor)

RBAC

```text
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
```

```text
p, Admin, /ping*, (GET)|(POST)|(DELETE)|(PUT)
p, Normal, /ping*, (GET)|(POST)|(DELETE)|(PUT)
p, Guest, /ping*, (GET)
```


