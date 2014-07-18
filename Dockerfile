FROM ubuntu:14.04
MAINTAINER Brian Rosner <brosner@eldarion.com>
ENV PATH /app/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games
ENV GOPATH /app
RUN apt-get update
RUN apt-get install -y ruby2.0 wget git make
RUN gem install einhorn
RUN wget -qO- http://golang.org/dl/go1.3.linux-amd64.tar.gz | tar -C /usr/local -xzf -
ADD . /app
WORKDIR /app
RUN make
EXPOSE 80
CMD einhorn -b :80,r -m manual eighty -bind=einhorn@0
