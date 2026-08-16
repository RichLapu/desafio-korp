# ==========================================
# Etapa 1: Build (Onde compilamos o código)
# ==========================================
# Usamos a imagem mais recente do Go baseada no Alpine para suportar o Prometheus
FROM golang:alpine AS builder

WORKDIR /app

# Copia o nosso arquivo Go para dentro do container
COPY main.go .

# Inicia o módulo, baixa as dependências (tidy) e compila o código
RUN go mod init projeto-korp && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp main.go

# ==========================================
# Etapa 2: Execução (A imagem final de produção)
# ==========================================
# Usamos uma imagem Alpine pura apenas para rodar o binário
FROM alpine:latest

WORKDIR /root/

# Copia apenas o executável da Etapa 1
COPY --from=builder /app/http-server-projeto-korp .

# Comando de inicialização do serviço
CMD ["./http-server-projeto-korp"]