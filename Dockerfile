# 使用 ARM64 架构的 Alpine Linux 基础镜像
FROM arm64v8/alpine:latest

# 设置工作目录
WORKDIR /app

# 复制编译好的二进制文件到容器
COPY page-redirector-linux-arm64 /app/page-redirector

# 赋予执行权限
RUN chmod +x /app/page-redirector

# 暴露默认端口 2445（可通过环境变量 PR_PORT 修改）
EXPOSE 2445

# 设置环境变量（可选，可以在运行时覆盖）
ENV PR_PORT=2445

# 运行应用
CMD ["/app/page-redirector"]
