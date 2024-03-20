#!/bin/bash
#Punto 3. Script Generación Imagen de Docker
echo "Construyendo Imagen Docker"
cd webserver/
docker build -t go-webserver .
echo "Tageando contenedor local y remoto"
docker tag go-webserver:latest alarav/weathercomp-devops:latest
echo "Haciendo push imagen docker"
docker push alarav/weathercomp-devops
