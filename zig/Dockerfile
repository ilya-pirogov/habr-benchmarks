FROM alpine
MAINTAINER Ilya Pirogov <ilja.pirogov@gmail.com>

ENV ZIG_VERSION=x86_64-0.3.0+85edf55b
ENV PATH=/opt/zig-linux-$ZIG_VERSION:$PATH

RUN wget https://ziglang.org/builds/zig-linux-$ZIG_VERSION.tar.xz -O /zig.tar.xz \
    && tar -xf /zig.tar.xz -C /opt \
    && rm /zig.tar.xz
