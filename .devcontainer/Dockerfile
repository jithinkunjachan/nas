FROM mcr.microsoft.com/devcontainers/base:debian

RUN wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz 
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz


RUN wget https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_3.13.2-stable.tar.xz
RUN  tar -C /usr/local -xf flutter_linux_3.13.2-stable.tar.xz

RUN echo "export PATH=$PATH:/usr/local/flutter/bin:/usr/local/go/bin" >> /home/vscode/.zshrc