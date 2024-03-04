FROM golang:1.22.0-alpine3.18         
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main main.go 

EXPOSE 8080 

CMD [ "/app/main" ]  