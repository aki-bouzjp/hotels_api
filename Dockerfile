FROM --platform=linux/amd64 ubuntu:20.04 AS development

ENV LANG en_US.utf8
ENV ROOT /app
ENV TZ=Asia/Tokyo
ENV ARCH amd64
ENV GOVERSION 1.17.1
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV DEBIAN_FRONTEND=noninteractive


WORKDIR $ROOT
COPY . $ROOT

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
    && apt-get -y update \
    && apt-get -y upgrade \
    && apt -y update \
    && apt -y upgrade \
    && apt -y install mysql-server \
    && apt -y install mysql-client \
    && apt -y install libmysqlclient-dev \
    && apt -y install redis-server \
    && apt-get -y install curl wget vim git-core gcc

RUN curl -s -o /tmp/go.tar.gz https://storage.googleapis.com/golang/go$GOVERSION.linux-$ARCH.tar.gz \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \
    && rm /tmp/go.tar.gz \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# RUN curl https://nodejs.org/dist/$(cat .nvmrc)/node-$(cat .nvmrc)-linux-x64.tar.gz | tar -xz -C /usr/local --strip-components 1 \
RUN tar -zxvf ./node-$(cat .nvmrc)-linux-x64.tar.gz -C /usr/local --strip-components 1


FROM development AS prisma-migrate
RUN npm install
# ENTRYPOINT [ "npm", "run", "prisma:migrate:dev" ]
ENTRYPOINT [ "node" ]

# FROM development AS prisma-generate
# RUN npm install
# ENTRYPOINT [ "npm", "run", "prisma:generate:dev" ]

# FROM development AS app-build
# ADD . .

# # RUN apk add build-base
# RUN go mod download
# RUN mkdir -p /tmp/log/app
# RUN mkdir -p /tmp/app

# ENTRYPOINT go build -o app ./src/main.go
# ENTRYPOINT go build -o app ./src/main.go && ./app -c app.toml



# FROM development AS watch-promoted-js

# COPY --from=build-promoted-js /app/mapbox-promoted-js-$SDK_VERSION.tar.gz /app/mapbox-promoted-js-$SDK_VERSION.tar.gz
# COPY . .

# ADD scripts/run_watch.sh ./scripts/run_watch.sh
# RUN chmod +x ./scripts/run_watch.sh

# ENTRYPOINT ["sh", "./scripts/run_watch.sh"]



# FROM golang:alpine

# WORKDIR /app

# ADD . .

# RUN apk add build-base
# RUN go mod download

# RUN mkdir -p /tmp/log/adserver
# RUN mkdir -p /tmp/adserver

# ENTRYPOINT go build -o adserver ./cmd/*.go  && ./adserver -c adserver.toml
