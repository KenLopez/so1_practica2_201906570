[Regresar](../README.md)
# Backend Desarrollado en Go (19.1)
Listener que se encarga de leer los módulos cargados en /proc y guarda la información recopilada en una base de datos MySQL en la nube cada dos segundos.

## Structs
### Parent
Structura para relacionar los procesos padre e hijo a guardar en base de datos.
```
type Parent struct {
	Value    *Proc
	Children []*Proc
}
```

### Ram
Estructura JSON guardada por el módulo ram_201906570.
```
type Ram struct {
	Totalram int `json:"totalram"`
	Freeram  int `json:"freeram"`
}
```

### Proc 
Estructura JSON guardada por el módulo cpu_201906570.
```
type Proc struct {
	Pid      int    `json:"pid",omitempty`
	Nombre   string `json:"nombre",omitempty`
	Usuario  int    `json:"usuario",omitempty`
	Estado   int    `json:"estado",omitempty`
	Ram      int    `json:"ram",omitempty`
	Children []int  `json:"children",omitempty`
}
```
### CPU 
Estructura JSON que recopila todos los datos de los módulos para fácil acceso.
```
type Proc struct {
	Pid      int    `json:"pid",omitempty`
	Nombre   string `json:"nombre",omitempty`
	Usuario  int    `json:"usuario",omitempty`
	Estado   int    `json:"estado",omitempty`
	Ram      int    `json:"ram",omitempty`
	Children []int  `json:"children",omitempty`
}
```

## Tablas Afectadas en cada Registro
* LOG: 1 fila.
* PROCESO: 1 fila por proceso leído de /proc.