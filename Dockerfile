# Etapa de build
FROM golang:1.22.8-bookworm AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia o go.mod e go.sum e instala as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código da aplicação
COPY . .

# Compila o binário da aplicação para produção
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main/goecommerce-golang-model-api ./main/main.go

# Etapa final: usa uma imagem leve para produção
FROM debian:bookworm-slim

# Define o diretório de trabalho na imagem final
WORKDIR /app

# Copia o binário compilado da etapa de build
COPY --from=builder /app/main/goecommerce-golang-model-api .
COPY --from=builder /app/.env .
RUN echo "Fake go.mod file" > ./go.mod
RUN mkdir .tmp

# Expõe a porta usada pela API (ajuste conforme a configuração do Gin)
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./goecommerce-golang-model-api"]
