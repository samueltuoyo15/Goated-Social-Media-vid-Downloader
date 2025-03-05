FROM golang:1.23.4-alpine  
RUN apk update && apk add --no-cache python3 py3-pip libwebp libwebp-dev build-base  
RUN python3 -m venv /opt/venv  
RUN /opt/venv/bin/pip install --no-cache-dir yt-dlp  
ENV PATH="/opt/venv/bin:$PATH"  
ENV CGO_ENABLED=1
WORKDIR /app  