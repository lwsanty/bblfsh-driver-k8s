FROM ubuntu:latest

ADD test /usr/bin/test

ENTRYPOINT ["/usr/bin/test", "-test.v", "-test.run=TestK8s"]
