[Regresar](../README.md)
# Backend
Contiene los módulos a cargar en el sistema operativo y el listener a implementar para guardar los datos periódicamente.

* [Módulos de Kernel](./modules/README.md)
* [Backend en Go](./src/README.md)

Usa el archivo docker compose para montar la imagen del backend, debes cargar los módulos antes de correr la imagen de go. 
```
$ docker-compose up
```