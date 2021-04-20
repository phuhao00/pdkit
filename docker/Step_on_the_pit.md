





dokerfile

```dockerfile
FROM golang:alpine3.13 

RUN mkdir ./src/test

WORKDIR ./src/test

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
RUN go env
ADD ./main.go   .
RUN go mod init test
RUN go mod tidy
RUN go mod vendor
RUN ls
RUN pwd
RUN CGO_ENABLED=0 GO111MODULE=on  go build -ldflags  "-s -w"  -o docker_test
CMD ["ls","./docker_test"]

```





删除 标签为 <none> 的 image 标签



```shell
 docker rmi --force $(docker images | grep none | awk '{print $3}')
```

