FROM ubuntu:24.04

RUN apt-get update && apt-get install -y golang-go

COPY deb-package /tmp/deb-package

RUN dpkg -i /tmp/deb-package

CMD ["/usr/bin/factor_app"]