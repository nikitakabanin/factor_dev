FROM ubuntu:24.04

RUN apt-get update && apt-get install -y golang-go

COPY go.deb /tmp/go.deb

RUN dpkg -i /tmp/go.deb

RUN chmod +x /usr/local/bin/factor_dev

COPY loop.sh /usr/bin/loop.sh
RUN chmod +x /usr/bin/loop.sh

CMD ["/usr/bin/loop.sh"]