FROM golang

RUN mkdir /app
COPY dugdugsehat-backend /app/
WORKDIR /app

RUN [ "chmod", "u+x", "dugdugsehat-backend" ]
ENTRYPOINT [ "./dugdugsehat-backend" ]