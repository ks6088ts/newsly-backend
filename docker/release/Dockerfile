FROM newsly-backend_go:latest AS builder

ARG WORK_DIR=""
ARG HTTP_PROXY=""
ARG HTTPS_PROXY=""

WORKDIR ${WORK_DIR}

COPY . .

RUN make build \
        GOBUILD='GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build' \
        BIN_PATH='outputs/newsly-backend'

# ---

FROM alpine:3.12.1

ARG WORK_DIR=""

COPY --from=builder ${WORK_DIR}/outputs/newsly-backend ./cmd
