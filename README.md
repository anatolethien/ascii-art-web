# ascii-art-web

Complete web application that allows users to convert text into an ascii-art representation of that text.

![Desktop version](https://raw.githubusercontent.com/anatolethien/anatolethien/main/images/ascii-art-web/app-desktop.png)

The user can choose between three different **fonts**: _standard_, _shadow_ and _thinkertoy_. The application supports **line feeds**, juste type the `\n` escape character to indicate a line feed. The user can also choose to **copy** the output to their clipboard, or even **save** it locally in a `.txt` file.

## Classic run

First, build the `ascii-art-web` binaries in the `bin/` directory using the `go` compiler.

    go build -o bin/ascii-art-web cmd/main.go

Then, execute those binaries.

    ./bin/ascii-art-web

## Docker run

First, build the docker image using the provided `Dockerfile`.

    docker build -t ascii-art-web .

Then, run the docker image on port `3000`.

    docker run -dp 3000:80 ascii-art-web
