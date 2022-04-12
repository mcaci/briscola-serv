FROM scratch

LABEL author="Michele Caci <michele.caci@gmail.com>"

WORKDIR /app
COPY briscolad /app

EXPOSE 8080 8081

ENTRYPOINT ["./briscolad", "-d"]
