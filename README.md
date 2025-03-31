# go-Conorch

## _Container orchestrator coupled with AI recommendation_

go-Conorch is a CLi tool for performing container orchestration via AI recommendation.

- Written in Golang
- Uses golang's self sufficient standard library, net/http
- Uses [urfave/cli](https://cli.urfave.org/)
- Quick, easy to use

## Features

- AI-driven Task Orchestration – Uses an LLM to analyze and recommend the best containerized task execution strategy.
- Dynamic Task Execution – Runs and manages containerized tasks based on AI recommendations.
- Manual Task Control – Supports manual task execution via CLI using [urfave/cli](https://cli.urfave.org/).
- Lightweight Backend – Built using the Go standard library for high efficiency and minimal dependencies.



## Tech

go-Conorch uses a number of projects to work properly:

- [Golang] - Golang language
- [net/http] - mux, servering
- [urfave/cli] - CLI handling

## Installation

go-Conorch requires [Golang](https://go.dev/) to run.

Install the dependencies and devDependencies and start the server.

```sh
# Clone the repository
git clone https://github.com/briheet/go-Conorch.git
cd go-Conorch

# Get GROQ API Key and place it .env file just like .env.example

# Normal build and run
make
./cmd/bin run -f prompt.txt
```

## Design

Desing has been kept minimal, different directories are present in tools. All these are main packages and 
are hence independent. This independence helps us in developing alone repos and updating them via git modules.

[Link to Diagram (Click me)](https://excalidraw.com/#json=h_3ohAoRMT6PZg1jTEm7y,jwbMN5hgFIirNoTmWK0S7A)
