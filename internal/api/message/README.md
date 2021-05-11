# 消息中心

> 消息通知：
> 为其它API提供最底层的统一消息网关，比如 工单，监控，任务异常等

## 需求

- 发送渠道类型
- 发送内容及相关信息保存数据库，方便页面展示分析
- API接口设计可统一，通过类型来区分不同的发送渠道，尽量让接口用起来更简单
- 扩展: 发送内容，服务端支持模版格式选择

### 发送消息
```text
POST /api/v1/message
Header
    Content-Type: application/json
    Authorization: JWT xxxxx

{
    "channel": "email/ short_letter/ phone/ ding/ wechat"
    "data": {
       
    }
}
```

注意
1. 如果需要用到第三方ak/sk，需要本地建表存储。

### 查询消息
```text
GET /api/v1/message
Header
    Content-Type: application/json
    Authorization: JWT xxxxx

{
   "code": 0,
   "message": nil,
   "data": []
}
```
