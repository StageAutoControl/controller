FROM alpine

MAINTAINER Alexander Pinnecke <alexander.pinnecke@googlemail.com>

ADD bin/controller_linux /controller

ENTRYPOINT ["/controller"]