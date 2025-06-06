# Go Gossip Gloomers CLI

A simple CLI tool to help run the [Gossip Gloomers](https://fly.io/dist-sys/) distributed systems challenges with ease using Docker. It provides a more user-friendly interface and allows you to manage system resources efficiently.

## 🚀 Features

- Docker-based setup for easy execution and isolation, the image contains everything needed to run it(eg maelstrom)
- User-friendly menu to run each one of the challenges from [Fly.io](https://fly.io/dist-sys/)
- Restart tests pressing by just `R`
  ![image](https://github.com/user-attachments/assets/ec58b665-5268-4e73-9cec-6a80269c0708)
  ![image](https://github.com/user-attachments/assets/f032e155-dd45-4e6d-8284-f861837d23b4)

## 🐳 How to Run

Make sure your code goes on `challenges` folder. So if you wanna run the `echo` workload for the Echo challenge you must provide the handler on the `challenges/main.go`. See the [challenge page](https://fly.io/dist-sys/1/) for more details
Example:

```go
func main() {
	n := maelstrom.NewNode()
	n.Handle("echo", func(msg maelstrom.Message) error {
    // Unmarshal the message body as an loosely-typed map.
    var body map[string]any
    if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
    }

    // Update the message type to return back.
    body["type"] = "echo_ok"

    // Echo the original message back with the updated message type.
    return n.Reply(msg, body)
  })

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
```

#### 1. Build the Docker Image

```sh
docker build -t gossip-cli .
```

#### 2. Run the container

```sh
docker run -it -v ./challenges:/app/challenges gossip-cli
```

> **Note**
> Some tests are resource-intensive. Running them inside Docker helps manage CPU and memory usage.

If your machine slows down during tests, run with limits:

```sh
docker run -it \
  -m 6GB \
  --cpus=6 \
  -v ./challenges:/app/challenges \
  gossip-cli
```
