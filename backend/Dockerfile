FROM golang:1.15-buster

# non-root user
ARG USERNAME=appuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && mkdir -p $GOPATH/pkg \
    && chmod -R 777 $GOPATH/pkg

RUN go get github.com/cespare/reflex

WORKDIR /go/src/app

EXPOSE 3000
