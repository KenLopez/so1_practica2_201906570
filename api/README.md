[Regresar](../README.md)
# API Desarrollada en Node (16)
API REST diseñada para consultar los datos cargados a una instancia de SQL en la nube, montada en docker.

## Puerto Expuesto: 3000

## Dependencias
- express
- mysql
- cors
- nodemon

## Rutas

### GET /data
Retorna los datos generales del los registros.
```
{
    "id": number,
    "fecha": string,
    "cpu": number,
    "ram": number
}
```

### GET /data/process
Retorna la información de procesos.
```
{
    "ejecucion": number,
    "suspendidos": number,
    "detenidos": number,
    "zombie": number,
    "total": string,
    "procs": [
        {
            "id": number,
            "pid": number
            "nombre": string
            "usuario": number
            "estado": string
            "ram": number
            "padre": null
            "children": [
                {
                    "id": number,
                    "pid": number
                    "nombre": string
                    "usuario": number
                    "estado": string
                    "ram": number
                    "padre": number
                }
            ]
        }
    ]
}
```