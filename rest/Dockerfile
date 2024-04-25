# Definir una variable para especificar la versión de go
ARG GO_VERSION=1.19.9

# Se hace un pull de las imágenes para que se rendericen
# alpine es una versión de construcción muy pequeña para
# que contruya el binario de las aplicaciones
FROM golang:${GO_VERSION}-alpine AS builder

# Se definen algunas variables de entorno
# No se establece ningún proxy para definir de dónde descargar los módulos
RUN go env -w GOPROXY=direct
# Se instala git, para permitir que los modulos de go.mod sean descargadas
RUN apk add --no-cache git
# Se instalan algunos certificados para que la app pueda correr
# Al mismo tiempo también se hace una actualización de ellos
RUN apk --no-cache add ca-certificates && update-ca-certificates

# Se define un directorio de trabajo nuevo para la ejecución de comandos
WORKDIR /src

# Se copian los módulos de las dependencias
COPY ./go.mod ./go.sum ./
# Se instalan las dependencias
RUN go mod download

# Luego se copian todos los directorios
COPY ./ ./

# CONTRUIR LA IMAGEN
# Se especifica que no sea habilitado MinGw
# Se le especifica a go que no utilice ese compilador
# Flags
# Por ultimo se define cómo se va a llamar el ejecutable
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /platzi-rest-ws .


# Se llama a otra imagen encargada de ejecutar la aplicación
FROM scratch AS runner

# Se copian los certificados
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Se copia el archivo de las variables de entorno
COPY .env ./

# Se copia la imagen del builder a la imagen del runner
COPY --from=builder /platzi-rest-ws /platzi-rest-ws

# Se especifica el puerto que se va a exponer
EXPOSE 5050

# Configura el contenedor para ser corrido como un ejecutable
ENTRYPOINT ["/platzi-rest-ws"]