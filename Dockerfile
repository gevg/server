
 FROM golang:latest

# EXPOSE 8088


# # FROM alpine:3.4

# # RUN apk add --no-cache ca-certificates

# # ENV GOLANG_VERSION 1.6.4
# # ENV GOLANG_SRC_URL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz
# # ENV GOLANG_SRC_SHA256 8796cc48217b59595832aa9de6db45f58706dae68c9c7fbbd78c9fdbe3cd9032

# # # https://golang.org/issue/14851
# # # COPY no-pic.patch /
# # # # https://golang.org/issue/17847
# # # COPY 17847.patch /

# # RUN set -ex \
# # 	&& apk add --no-cache --virtual .build-deps \
# # 		bash \
# # 		gcc \
# # 		musl-dev \
# # 		openssl \
# # 		go \
# # 	\
# # 	&& export GOROOT_BOOTSTRAP="$(go env GOROOT)" \
# # 	\
# # 	&& wget -q "$GOLANG_SRC_URL" -O golang.tar.gz \
# # 	&& echo "$GOLANG_SRC_SHA256  golang.tar.gz" | sha256sum -c - \
# # 	&& tar -C /usr/local -xzf golang.tar.gz \
# # 	&& rm golang.tar.gz \
# # 	&& cd /usr/local/go/src \
# # 	# && patch -p2 -i /no-pic.patch \
# # 	# && patch -p2 -i /17847.patch \
# # 	&& ./make.bash \
# # 	\
# # 	&& rm -rf /*.patch \
# # 	&& apk del .build-deps

# # ENV GOPATH /go
# # ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# # RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
# # WORKDIR $GOPATH

# # # COPY go-wrapper /usr/local/bin/


# # golang image where workspace (GOPATH) configured at /go. 

# FROM debian:jessie

# # Install Go
# # IMPORTANT: If the version of Go is updated, the Windows to Linux CI machines
# #            will need updating, to avoid errors. Ping #docker-maintainers on IRC
# #            with a heads-up.

# # Packaged dependencies
# RUN apt-get update && apt-get install -y \
# 	apparmor \
# 	apt-utils \
# 	aufs-tools \
# 	automake \
# 	bash-completion \
# 	binutils-mingw-w64 \
# 	bsdmainutils \
# 	btrfs-tools \
# 	build-essential \
# 	clang \
# 	cmake \
# 	createrepo \
# 	curl \
# 	dpkg-sig \
# 	gcc-mingw-w64 \
# 	git \
# 	iptables \
# 	jq \
# 	less \
# 	libapparmor-dev \
# 	libcap-dev \
# 	libltdl-dev \
# 	libnl-3-dev \
# 	libprotobuf-c0-dev \
# 	libprotobuf-dev \
# 	libsqlite3-dev \
# 	libsystemd-journal-dev \
# 	libtool \
# 	libzfs-dev \
# 	mercurial \
# 	net-tools \
# 	pkg-config \
# 	protobuf-compiler \
# 	protobuf-c-compiler \
# 	python-dev \
# 	python-mock \
# 	python-pip \
# 	python-websocket \
# 	tar \
# 	ubuntu-zfs \
# 	vim \
# 	vim-common \
# 	xfsprogs \
# 	zip \



# ENV GO_VERSION 1.7.4
# RUN curl -fsSL "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
# 	| tar -xzC /usr/local

# ENV PATH /go/bin:/usr/local/go/bin:$PATH
# ENV GOPATH /go


# # Dependency for golint
# ENV GO_TOOLS_COMMIT 823804e1ae08dbb14eb807afc7db9993bc9e3cc3
# RUN git clone https://github.com/golang/tools.git /go/src/golang.org/x/tools \
# 	&& (cd /go/src/golang.org/x/tools && git checkout -q $GO_TOOLS_COMMIT)

# # Grab Go's lint tool
# ENV GO_LINT_COMMIT 32a87160691b3c96046c0c678fe57c5bef761456
# RUN git clone https://github.com/golang/lint.git /go/src/github.com/golang/lint \
# 	&& (cd /go/src/github.com/golang/lint && git checkout -q $GO_LINT_COMMIT) \
# 	&& go install -v github.com/golang/lint/golint


# Copy the local package files to the container’s workspace. 
ADD . /go/src/github.com/goa-fhir/server 
# Setting up working directory 
WORKDIR /go/src/github.com/goa-fhir/server 
# Get godeps for managing and restoring dependencies 
RUN go get github.com/kardianos/govendor
# Restore godep dependencies 
RUN govendor install +local
# RUN godep restore 
# Build the taskmanager command inside the container. 
RUN go install github.com/goa-fhir/server 
# Run the taskmanager command when the container starts. 
ENTRYPOINT /go/bin/server
# Service listens on port 8080. 
EXPOSE 8080

# COPY . /go/src/github.com/goa-fhir/server

