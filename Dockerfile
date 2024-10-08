FROM golang:1.23-alpine as builder
WORKDIR /app

RUN go install github.com/joho/godotenv/cmd/godotenv@latest
# Copiando o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixando as dependências
RUN go mod download

RUN apk update && \
    apk add --no-cache \
    coreutils \
    git \
    make
# Copiando o código para o diretório de trabalho
COPY . .

# Compilando a aplicação
RUN make build

# Expondo a porta da aplicação
EXPOSE 8080

FROM alpine:3.11
COPY --from=builder /app/bin/linux_amd64/shortener /usr/bin

COPY --from=builder /app/.env . 

# Definindo o comando de inicialização
CMD ["/usr/bin/shortener"]