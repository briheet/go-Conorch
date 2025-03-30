# go-Conorch

## _Container orchestrator coupled with AI recommendation

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
cd tkgo

# Get GROQ API Key and place it .env file just like .env.example

# Normal build and run
make
./cmd/bin run -f now.txt
```

<!-- ## Docker -->
<!---->
<!-- Tkgo is very easy to use and deploy in a Docker container. -->
<!---->
<!-- By default, the Docker will expose port 8080, so change this within the -->
<!-- Dockerfile if necessary. When ready, simply use the Dockerfile to -->
<!-- build the image. -->
<!---->
<!-- ```sh -->
<!-- # Enter the project directory -->
<!-- cd Tkgo -->
<!---->
<!-- # Directly build the image -->
<!-- docker build -t tkgo:multistage -f Dockerfile.multistage . -->
<!---->
<!-- # Or use Makefile -->
<!-- make docker-build -->
<!-- ``` -->
<!---->
<!-- This will create the Tkgo image and pull in the necessary dependencies. -->
<!---->
<!-- Once done, run the Docker image and map the port to whatever you wish on -->
<!-- your host. For now, we simply map port 8080 of the host to -->
<!-- port 8080 of the Docker (or whatever port was exposed in the Dockerfile): -->
<!---->
<!-- ```sh -->
<!-- # Directly run the image -->
<!-- docker run -p 8080:8080 tkgo:multistage -->
<!---->
<!-- # Or use the Makefile -->
<!-- make docker-run -->
<!-- ``` -->
<!---->
<!-- Verify the deployment by navigating to your server address in -->
<!-- your preferred browser. -->
<!---->
<!-- ```sh -->
<!-- http://localhost:8080/health -->
<!-- ``` -->
<!---->
<!-- ## Insights -->
<!---->
<!-- ![Diagram](./docs/dianew.png) -->
<!---->
<!-- Create User Request Body -->
<!---->
<!-- ``` -->
<!-- { -->
<!--     "userInfo": { -->
<!--         "userName": "John", -->
<!--         "userId": "1234" -->
<!--     }, -->
<!--     "simulationTime": 5, -->
<!--     "tokenNumbers": 5 -->
<!-- } -->
<!-- ``` -->
<!---->
<!-- GetToken Request Body -->
<!---->
<!-- ``` -->
<!-- { -->
<!--     "userId": "1234" -->
<!-- } -->
<!-- ``` -->
<!---->
<!-- ## Development -->
<!---->
<!-- Going on. Want to contribute? Make a pr :) -->
<!---->
<!-- [//]: # "These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax" -->
<!-- [net/http]: https://pkg.go.dev/net/http -->
<!-- [Zap (Logging)]: https://github.com/uber-go/zap -->
<!-- [Golang]: http://go.dev -->
