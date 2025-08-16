FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend

COPY mvc_frontend/package.json mvc_frontend/package-lock.json ./

RUN npm install

COPY mvc_frontend ./

RUN npm run build

FROM golang:1.24-alpine AS backend-builder

WORKDIR /app/backend

COPY mvc_backend/go.mod mvc_backend/go.sum ./

RUN go mod download

COPY mvc_backend ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

FROM golang:1.24-alpine AS migrate-builder
RUN apk --no-cache add git
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add mysql-client

RUN mkdir -p /app/static

COPY --from=backend-builder /server .

COPY --from=frontend-builder /app/frontend/dist /app/static

COPY --from=migrate-builder /go/bin/migrate /migrate

COPY ./mvc_backend/database/migrations /migrations

COPY ./migrate.sh /migrate.sh

RUN chmod +x /migrate.sh

COPY mvc_backend/.env.sample ./.env

EXPOSE 8080

ENTRYPOINT ["/migrate.sh"]

CMD ["./server"]