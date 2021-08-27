FROM --platform=x86_64 golang:1.17.0-buster AS development

# ================= SAP RFC Configuration ======================
RUN mkdir -p /usr/local/sap
COPY ./sap-nw-sdk.tar.gz .
RUN tar -zxf sap-nw-sdk.tar.gz --directory /usr/local/sap
RUN rm sap-nw-sdk.tar.gz
ENV SAPNWRFC_HOME=/usr/local/sap/nwrfcsdk
RUN echo "/usr/local/sap/nwrfcsdk/lib" > /etc/ld.so.conf.d/nwrfcsdk.conf
RUN ldconfig -v | grep sap
# ==============================================================

RUN mkdir -p /app
RUN mkdir -p /data
WORKDIR /app


FROM --platform=x86_64 golang:1.17.0-buster AS production
# ================= SAP RFC Configuration ======================
RUN mkdir -p /usr/local/sap
COPY ./sap-nw-sdk.tar.gz .
RUN tar -zxf sap-nw-sdk.tar.gz --directory /usr/local/sap
RUN rm sap-nw-sdk.tar.gz
ENV SAPNWRFC_HOME=/usr/local/sap/nwrfcsdk
RUN echo "/usr/local/sap/nwrfcsdk/lib" > /etc/ld.so.conf.d/nwrfcsdk.conf
RUN ldconfig -v | grep sap
# ==============================================================

RUN useradd -ms /bin/bash user
USER user

RUN mkdir -p /home/user/app
WORKDIR /home/user/app
COPY --chown=user:user go.mod ./
COPY --chown=user:user go.sum ./
RUN go mod download

COPY --chown=user:user . .
RUN go build -tags=jsoniter -o /home/user/sap-bridge .
RUN rm -rf /home/user/app
ENV GIN_MODE=release
EXPOSE 80
ENTRYPOINT [ "/home/user/sap-bridge" ]
