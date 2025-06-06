# Go Gossip Gloomers CLI

A simple CLI tool to help run the [Gossip Gloomers](https://fly.io/dist-sys/) distributed systems challenges with ease using Docker. It provides a more user-friendly interface and allows you to manage system resources efficiently.

## 🚀 Features

- Docker-based setup for easy execution and isolation
- Resource limits to prevent your machine from struggling during tests
- Volume mounting for working directly with challenge files
- Restart tests pressing by just `R`
![image](https://github.com/user-attachments/assets/ec58b665-5268-4e73-9cec-6a80269c0708)
![image](https://github.com/user-attachments/assets/f032e155-dd45-4e6d-8284-f861837d23b4)


## 🐳 How to Run

Make sure your code goes on `challenges` folder.

#### 1. Build the Docker Image 

```sh
docker build -t gossip-cli .
```
#### 2. Run the container
```sh
docker run -it -v ./challenges:/app/challenges gossip-cli
```
> **Note**
Some tests are resource-intensive. Running them inside Docker helps manage CPU and memory usage.

If your machine slows down during tests, run with limits:
```sh
docker run -it \
  -m 6GB \
  --cpus=6 \
  -v ./challenges:/app/challenges \
  gossip-cli
```
