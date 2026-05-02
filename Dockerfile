FROM ubuntu:24.04

RUN apt-get update && apt-get install -y golang-go

COPY go.deb /tmp/go.deb

RUN dpkg -i /tmp/go.deb

RUN chmod +x /usr/bin/factor_dev

COPY run_loop.sh /usr/bin/run_loop.sh
RUN chmod +x /usr/bin/run_loop.sh

CMD ["/usr/bin/run_loop.sh"]