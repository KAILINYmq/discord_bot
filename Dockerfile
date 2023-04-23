ARG GO_VERSION

FROM golang:1.16-alpine as builder


WORKDIR /builder/

COPY . .

ENV PATH="/go/bin:${PATH}"
ENV TZ=Asia/Shanghai \
    CGO_ENABLED=1   \
    GO111MODULE=on  \
    GOOS=linux

# install cgo
RUN apk -U add ca-certificates
RUN apk add build-base
RUN apk add --no-cache bash openssh-client git sudo

# Add ssh private key, to download dependencies from private repositories
RUN mkdir /root/.ssh/
ADD .ssh/id_ed25519 /root/.ssh/id_ed25519
RUN chmod 400 /root/.ssh/id_ed25519

# make sure github domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

#RUN go get ./...
RUN go mod download
RUN go build -tags musl --ldflags "-extldflags -static" -o Discord-Roles-Bot .

FROM scratch
#FROM alpine:latest

ENV TZ=Asia/Shanghai \
    GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    PROGRAM_ENV=pro

WORKDIR /app

COPY --from=builder /builder/Discord-Roles-Bot .
COPY ./conf ./conf

EXPOSE 8091

# 启动服务
CMD ["./Discord-Roles-Bot"]