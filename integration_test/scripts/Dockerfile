FROM irisnet/irishub:v2.0.3

COPY . /scripts

RUN sh /scripts/setup.sh

EXPOSE 26657
EXPOSE 9090

CMD iris start