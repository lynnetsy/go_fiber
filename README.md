## API

Inicializar un paquete
```bash
go mod init <module_name>
```

Instalar paquete
```bash
go get <package_name>
```

Limpiar dependencias
```bash
go mod tidy
```

Ejecutar un paquete
```bash
go run .
```

CMD para instalar el paquete que sirve para detectar cambios y reiniciar
el paquete
```bash
go install github.com/cosmtrek/air@latest
```

Inicializar air (para esto, se necesita inicializar dentro de un pckg)
```bash
air init
```

Iniciar watcher
```bash
air
```
