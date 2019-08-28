# Nous démarrons d'une image golang:alpine avec la dernière version de Go
# Elle va nous servir d'environnement de build
FROM golang:alpine3.8 as builder

# Nous chageons de répertoire courant pour aller là ou le code a été cloné
WORKDIR /root/glmf

# build du code source
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/glmf

# Nous transférons le binaire vers une image scratch pour diminuer le footprint
FROM scratch
COPY --from=builder /go/bin/glmf .

# Le binaire est placé comme entrypoint pour l'image
ENTRYPOINT ["./glmf"]

# Nous documentons que le conteneur écoute sur le port 8080
EXPOSE 8080
