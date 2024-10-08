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

`make start`

Docker is required to run application using the command above.

Execute the command to short url.

Example:

`curl --location 'http://localhost:8080/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url_original" : "https://www.mercadolivre.com.br/cafeteira-italiana-de-aluminio-600ml-para-6-xicaras-cor-prateado/p/MLB19846937#polycard_client=recommendations_home_navigation-recommendations&reco_backend=machinalis-homes-univb-equivalent-offer&wid=MLB5060623436&reco_client=home_navigation-recommendations&reco_item_pos=5&reco_backend_type=function&reco_id=20337492-08e5-4735-a4d3-95b384fbc207&sid=recos&c_id=/home/navigation-recommendations/element&c_uid=46d88259-5ef0-4a6d-88df-62f2accf6453"
}`

Execute the command to get the original url.

Exemplo:

`curl --location 'http://localhost:8080/getOriginalUrl/bed8a7f5`

Execute the command to delete the url.

Exemplo:

`curl --location --request DELETE 'http://localhost:8080/delete/bed8a7f5`

## Design Decisions

### Files organization

To keep it simple, the chosen file structure was the **Flat Structure**. This is a good choise for small applications and microservices.

##### References:
- [Go Project Structure Best Practices](https://tutorialedge.net/golang/go-project-structure-best-practices/)
- [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0&t=245s)
