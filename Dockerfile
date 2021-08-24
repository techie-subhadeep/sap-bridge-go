FROM --platform=x86_64 golang:1.17.0-buster

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
