# FROM alpine/git
FROM  golang

WORKDIR /play

COPY ./a-git .
COPY ./tests ./tests

RUN ls

CMD ["bash", "./tests/runner.sh"]