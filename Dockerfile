# 构建前端项目
FROM node:lts as front
WORKDIR /app
COPY console/package.json .
COPY console/yarn.lock .
RUN yarn install
COPY console/ .
RUN yarn build

# 构建后端项目
FROM golang:1.23 as back
WORKDIR /app
COPY . .
COPY --from=front /app/dist /app/internal/controller/static
RUN GOOS=linux GOARCH=amd64 go build -o x-seek ./cmd
# 运行环境
FROM ubuntu:latest

WORKDIR /app
COPY --from=back /app/x-seek .
COPY conf/server.yaml config.yaml
# 安装CA证书
RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 8080
CMD ["./x-seek", "server", "-c", "config.yaml"]
