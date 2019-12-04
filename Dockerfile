FROM golang:1.12

WORKDIR /go-etl
# ENV SRC_DIR=/go/src/github.com/mwebster/go-etl/
# ADD . $SRC_DIR

ENV GOPATH /go
COPY . .

# RUN ["chmod", "+x", "./entrypoint.sh"]
# ENTRYPOINT [ "./entrypoint.sh"]

#RUN cd $SRC_DIR; go build -o myapp
# RUN cp myapp /go-etl/
RUN go build -o myetl
CMD [ "./myetl" ]