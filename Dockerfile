FROM golang:1.24.3

WORKDIR /app
RUN apt-get update && apt-get install -y \
    default-jdk \
    graphviz \
    gnuplot \
    curl \
    bzip2

# Download tar from https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 
# and extract it to /app
RUN curl -L https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 -o maelstrom.tar.bz2
RUN tar -xvf maelstrom.tar.bz2

COPY . .

RUN go build -o main main.go

CMD ["./main", "test"]