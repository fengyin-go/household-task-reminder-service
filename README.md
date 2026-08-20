# 待办清单 (todo-list)

纯 Go 标准库实现的待办清单后端服务，零第三方依赖，开箱即跑。

## 运行

```bash
go run ./cmd/server
# 自定义端口
PORT=9090 go run ./cmd/server
```

环境变量：

| 变量 | 说明 | 默认 |
|------|------|------|
| `PORT` / `ADDR` | 监听地址 | `:8080` |
| `MAX_PAGE_SIZE` | 分页最大条数 | `100` |
| `LOG_LEVEL` | 日志级别（debug/info/warn/error） | `info` |

## 架构分层

```
cmd/server            入口：配置加载、依赖装配、优雅关闭
internal/app          依赖装配 store -> service -> handler
internal/config       环境变量配置
internal/model        领域模型 + 校验 + 状态机
internal/store        数据访问接口 + 内存实现
internal/service      业务逻辑 + 状态流转 + 完成率统计
internal/handler      HTTP 路由与处理器
pkg/httpx             统一响应、分页、JSON 解析
pkg/idgen             ID 生成
pkg/logger            分级日志
```

## 核心业务

- 任务：CRUD、状态机 `todo -> doing -> done` / `cancelled`，支持状态流转校验。
- 清单：CRUD、按清单统计任务完成率（总数/已完成/进行中/待办/已取消）。
- 标签：CRUD、名称唯一性校验。
- 提醒：CRUD、关联任务外键校验。

## API 一览

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/tasks` | 创建任务 |
| GET | `/api/tasks` | 任务列表（task_list_id/status/keyword 筛选，分页） |
| GET/PUT/DELETE | `/api/tasks/{id}` | 任务详情 / 更新 / 删除 |
| POST | `/api/tasks/{id}/transition` | 任务状态流转 |
| POST | `/api/task-lists` | 创建清单 |
| GET | `/api/task-lists` | 清单列表（keyword 筛选，分页） |
| GET/PUT/DELETE | `/api/task-lists/{id}` | 清单详情 / 更新 / 删除 |
| GET | `/api/task-lists/{id}/stats` | 清单任务完成率统计 |
| POST | `/api/tags` | 创建标签 |
| GET | `/api/tags` | 标签列表（keyword 筛选，分页） |
| GET/PUT/DELETE | `/api/tags/{id}` | 标签详情 / 更新 / 删除 |
| POST | `/api/reminders` | 创建提醒 |
| GET | `/api/reminders` | 提醒列表（task_id 筛选，分页） |
| GET/PUT/DELETE | `/api/reminders/{id}` | 提醒详情 / 更新 / 删除 |

统一响应结构：`{"code":0,"message":"ok","data":...}`；错误码 400/404/409/500。
