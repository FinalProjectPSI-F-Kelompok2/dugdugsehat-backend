FROM golang

COPY . /app

WORKDIR /app
RUN [ "rm", "Dockerfile" ]
RUN [ "rm", "docker-compose.yml" ]
RUN [ "rm", ".env"]
RUN [ "go", "mod", "tidy" ]
RUN [ "go", "build", "." ]
RUN [ "chmod", "u+x", "dugdugsehat-backend" ]
ENTRYPOINT [ "./dugdugsehat-backend" ]