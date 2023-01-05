





### ¿Qué es el API de Searchella? 

Es un maravilloso script para poder indexar correo electrónico a Zincsearch ✨

### ¿Como está construida? 💻

Su Backend está construido con ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
Utilizando Chi como router.


Utiliza como base de datos Zincsearch

### ¿Como utilizarlo? 💻
- Posiciona el script indexer.go en la carpeta raiz donde se encuentran tus carpetas de correo y ejecutalo utilizando:

```
go run server.go
```


- Corre el script usando

```
go run indexer.go
```

- Espera que se genere un archivo llamado

```
output.json
```

### Siguientes mejoras a la aplicación 🧾
- Crear un api mas robusto.
- Crear un sistema de login.
- Desplegar en Terraform


