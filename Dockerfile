FROM alpine
RUN echo ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir -p /opt/gateway
ADD  gateway /opt/gateway/gateway
ADD  config.yaml /opt/gateway/gateway.yaml
ENV LANG C.UTF-8
ENV LANGUAGE zh_CN.UTF-8
ENV LC_ALL C.UTF-8
ENV TZ Asia/Shanghai
WORKDIR /opt/gateway
RUN chmod -R 755 /opt/gateway
CMD ["./gateway"]