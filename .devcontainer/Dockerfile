FROM golang:alpine3.10

ARG USERNAME=vscode
ARG DEVGROUP=developer


# Configure apt and install packages
RUN apk update \
	&& apk add git \
    && apk add bash \
    && apk add gcc \
    && apk add g++ \
    #
    #
    # Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
    && addgroup -S $DEVGROUP && adduser -S $USERNAME -G $DEVGROUP \
    # [Optional] Add sudo support for the non-root user
    && apk add sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME\
    && chmod 0440 /etc/sudoers.d/$USERNAME 