FROM centos:6.9

RUN set -eux; \
    yum groupinstall -y "Development Tools"; \
    yum install -y wget git glibc-static
ENV GOLANG_VERSION 1.9.2
RUN set -eux; \
    wget -O go.tar.gz "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"; \
    tar -C /usr/local -xf go.tar.gz; \
    rm go.tar.gz; \
    export PATH="/usr/local/go/bin:$PATH"; \
    go version
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go get -u -v github.com/naoina/kocha/cmd/...
RUN go get -u -v github.com/golang/dep/cmd/dep
