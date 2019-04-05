# FROM alpine/git
FROM  golang

WORKDIR /go/src/github.com/ritwik310/a-git

# TODO: Find some better solution to get packages (or not to)
RUN go get gopkg.in/ini.v1

COPY ./a-git .
COPY ./tests ./tests
COPY ./src ./src

RUN ls

CMD ["bash", "./tests/runner.sh"]