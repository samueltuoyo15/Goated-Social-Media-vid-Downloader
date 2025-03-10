FROM node:20-alpine AS client-builder

WORKDIR /client
COPY client/package.json client/package-lock.json ./
RUN npm ci

COPY client ./
RUN npm run build


FROM golang:1.23.5-alpine AS server-builder

WORKDIR /server
COPY server/go.mod server/go.sum ./
RUN go mod tidy

COPY server ./
RUN CGO_ENABLED=0 go build -o main


FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache python3 py3-pip libwebp libwebp-dev build-base ca-certificates

RUN python3 -m venv /opt/venv
RUN /opt/venv/bin/pip install --no-cache-dir yt-dlp
ENV PATH="/opt/venv/bin:$PATH"

COPY --from=client-builder /client/dist /app/client/dist
COPY --from=server-builder /server/main /app/main

RUN chmod +x /app/main

EXPOSE 10000

CMD ["/app/main"]