FROM golang

ENV GO111MODULE=on

WORKDIR  /app

COPY . .

EXPOSE $APP_PORT

RUN [ "go", "build", "-o","app", "main.go"]

RUN [ "chmod", "777", "app" ]

CMD [ "./app" ]