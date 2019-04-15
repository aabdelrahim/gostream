FROM alpine

RUN apk update
RUN apk add ca-certificates
ADD server ./server
VOLUME [ "/musicLibrary" ]
ENTRYPOINT ["./server"]