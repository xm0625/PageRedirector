# PageRedirector

一个轻量级的 Go HTTP 重定向服务，适用于内网部署场景。

## 功能特性

- 通过 URL 参数进行 302 重定向
- 支持通过环境变量配置端口
- 内置安全防护：
  - 仅允许 HTTP/HTTPS 协议
  - URL 长度限制 (2048 字节)
  - 仅支持 GET 请求

## 快速开始

### 使用 Go 运行

```bash
go run main.go
```

### 编译运行

```bash
go build -o pageredirector
./pageredirector
```

### Docker 部署

```bash
docker build -t pageredirector .
docker run -d -p 2445:2445 --env PR_PORT=2445 pageredirector
```

## 使用方法

### 基本用法

访问以下格式的 URL：

```
http://localhost:2445/?url=<URL编码后的目标地址>
```

### 示例

重定向到百度搜索：

```
http://localhost:2445/?url=http%3A%2F%2Fwww.baidu.com%2Fs%3Fwd%3Dtest
```

这会自动解码 URL 并重定向到：

```
http://www.baidu.com/s?wd=test
```

## 配置

| 环境变量 | 说明 | 默认值 |
|----------|------|--------|
| `PR_PORT` | 服务监听端口 | 2445 |

如果 `PR_PORT` 未设置或无效，将使用默认端口 2445。

## 环境要求

- Go 1.20 或更高版本

## 安全说明

- 本服务设计用于内网部署环境
- 仅允许 HTTP/HTTPS 协议，阻止 `javascript:`、`data:` 等危险协议
- 实施了 URL 长度限制以防止潜在的拒绝服务攻击
- 针对内网环境进行了优化，移除了不必要的复杂安全措施

## 许可证

MIT License

## Powered By

Powered By [智谱GLM](https://www.bigmodel.cn/invite?icode=K%2BSSh5%2Be4aMJI4Nd%2BNVTh%2BZLO2QH3C0EBTSr%2BArzMw4%3D)
非常好用，点击注册马上体验。一个字，丝滑
