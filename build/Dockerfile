FROM alpine:3.9

ENV OPERATOR=/usr/local/bin/kubedge-ecds-operator \
    USER_UID=1001 \
    USER_NAME=kubedge-ecds-operator

# install operator binary
COPY build/_output/bin/kubedge-ecds-operator ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

COPY build/ecds-templates /opt/kubedge-operators/ecds-templates

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
