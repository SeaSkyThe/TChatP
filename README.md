# TChatP

TChatP is a simple text chat application written in Go.

## Motivation

I just wanted to learn how TCP connections worked in Go and how I could use the Mutex in this scenario.

## Usage

To run the server you have to have go installed on your system.

And then run the following command:
```bash
go run main.go -mode server
```

To run the client, run the following command:

```bash
go run main.go -mode client
```


## Features

- Multiple users can connect to the server
- Messages are broadcasted to all connected users
- Colored output for better readability

