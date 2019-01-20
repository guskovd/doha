FROM alpine

ARG HAB_BINARY_PATH

ADD $HAB_BINARY_PATH /usr/bin/
RUN apk add sudo
RUN apk add bash
RUN apk add acl

# nopasswd sudo:
RUN echo '%wheel  ALL=(ALL)       NOPASSWD: ALL'  >> /etc/sudoers
RUN sed -i "s/^.*requiretty/#Defaults requiretty/" /etc/sudoers

RUN deluser guest
RUN delgroup users 

ENV PATH /hab/local/bin:/hab/bin:$PATH

