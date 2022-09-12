[Regresar](../README.md)
# Frontend Desarrollado con React
Aplicación de React desplegada en un contenedor de docker en conjunto con la aplicación API.

## Puerto Expuesto: 4200
## Componentes
### App
Componente padre de la aplicación, en él se contienen los demás componentes.

### Usage
Componente que hace uso de Chart.js para desplegar una gráfica de dona para mostrar el porcentaje de uso dado el parámetro y el título.
Props:
- usage
- title

### Tree
Componente de tabla que hace uso de DevExtreme para desplegar la información de los procesos en una tabla expandible y jerárquica.
Props:
- rows

### Counts
Componente que muestra los datos de conteo de procesos.
Props:
- ejecucion
- detenidos
- suspendidos
- zombie
- total

