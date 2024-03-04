# Build stage
FROM golang:1.22.0-alpine3.18 AS builder   
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct; go build -o main main.go 

# Run stage
FROM alpine:3.18
WORKDIR /app  
COPY --from=builder /app/main .

EXPOSE 8080 
CMD [ "/app/main" ]  


