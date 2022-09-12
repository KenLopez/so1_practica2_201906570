[Regresar](../README.md)
# Módulos de Kernel
Dos módulos de kernel desarrollados en c para la obtención de datos del sistema operativo.

## Módulo de RAM
Guarda los datos de memoria RAM total y memoria RAM libre del sistema.

### Nombre del módulo: ram_201906570
### Librerías utilizadas
* <linux/module.h>
* <linux/kernel.h>
* <linux/init.h>
* <linux/proc_fs.h>
* <asm/uaccess.h>	
* <linux/seq_file.h>
* <linux/hugetlb.h>

## Módulo de CPU
Guarda los datos de procesos del sistema
### Nombre del módulo: cpu_201906570
### Librerías empleadas

* <linux/module.h>
* <linux/kernel.h>
* <linux/init.h>
* <linux/proc_fs.h>
* <asm/uaccess.h>	
* <linux/seq_file.h>
* <linux/sched.h>
* <linux/mm.h>

## Makefile
* all: Compila los archivos necesarios para implementar los módulos.
* clean: Limpia los archivos creados por `make all`

## Script
Corre el script para limpiar los módulos, mensajes de consola e implementar los módulos con un solo comando.
```
$ bash module.sh
```


