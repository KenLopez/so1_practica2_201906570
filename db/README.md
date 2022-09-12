# Base de Datos MySQL
Base de datos montada en la nube en una instancia de GCP.

## Nombre: practica2

## Tablas

### Log
Mantiene un registro de la fecha y hora de cada inserción, guarda también los datos de uso de ram y cpu.

### Proceso
Guarda la información de cada proceso asociada a una inserción en la tabla log. También se guardan las relaciones padre-hijo de los procesos y la información de cada proceso.