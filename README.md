# url-shortener
url shortener application

### Service Architecture

![](/docs/URL_SHORTENER.png)

## Running

### Prerequisites

- [Golang v1.22+](https://golang.org/) 
- [Docker]([https://www.docker.com/) - Docker is required to execute integration test and to run application.
- [Make]([https://www.gnu.org/software/make/)

### Running lint

`make lint`

### Running staticcheck

`make staticcheck`

### Running unit test

`make test`

### Building

`make build`

The binary will be created at `<ROOT_DIR>/bin/<OS>_<ARCH>/` directory.

### Building docker image

`make docker-build`

Docker is required to execute this command

### Running integration test

Docker is required to execute the integration test.

`make integration-tests`

### Running application

To run application you must pass file as argument, as you can see below:

/*TODO*/

## Design Decisions

### Files organization

To keep it simple, the chosen file structure was the **Flat Structure**. This is a good choise for small applications and microservices.
##### References:
- [Go Project Structure Best Practices](https://tutorialedge.net/golang/go-project-structure-best-practices/)
- [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0&t=245s)

### calculation package

The calculation packages contais all domain business logics.


### application layer

Is responsable for parsing data entry and trigger the calculation service functions accordind to entry type.