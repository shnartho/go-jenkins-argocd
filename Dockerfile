FROM golang:1.21.0 AS builder
WORKDIR /app
COPY . .
RUN go build -o main
FROM nginx:latest
COPY --from=builder /app /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["sh", "-c", "nginx -g 'daemon off;' & /usr/share/nginx/html/main"]
