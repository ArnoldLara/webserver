#!/bin/bash
#Punto 4. Script Despliegue de imagen en local
echo "Clonando repositorio"
git clone https://github.com/ArnoldLara/webserver.git
cd webserver/
echo "Construyendo Imagen Docker"
docker build -t go-webserver .
echo "Corriendo contenedor"
echo "Puede acceder a la aplicación web accediendo al siguiente enlace: http://localhost/"
docker run -p 80:8080 -it go-webserver
