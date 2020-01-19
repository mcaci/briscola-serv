FROM scratch

WORKDIR /app
COPY briscolad /app

EXPOSE 8080 8081

ENTRYPOINT ["./briscolad"]