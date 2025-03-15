# 🚀 Caching Proxy

## 📌 Descripción
`Caching Proxy` es un servidor proxy con caché implementado en **Go**, que almacena las respuestas de las solicitudes HTTP para mejorar el rendimiento y reducir la carga en el servidor de origen. Este proyecto recibe importantes parametros por la linea de comandos del SO.

Cuando un cliente hace una solicitud, el proxy verifica si ya tiene la respuesta en caché. Si la encuentra (`HIT`), la devuelve directamente sin reenviar la solicitud al servidor. Si no está en caché (`MISS`), obtiene la respuesta del servidor de origen, la almacena en caché y la envía al cliente.

---

## 🎯 Características
✅ Almacena respuestas en caché por un tiempo configurable.  
✅ Maneja concurrencia con `sync.RWMutex` para evitar condiciones de carrera.  
✅ Devuelve respuestas en formato JSON.  
✅ Controla el estado de la caché mediante el encabezado `X-Cache` (`HIT` o `MISS`).  

---

## 🛠️ Instalación

**Clona este repositorio:**
```bash
 git clone https://github.com/Andres-cmk/proxy-cache.git
 cd caching-proxy
```
---
**Compila el proyecto y ejecuta:**
```bash
go run .\main.go -port <port> -origin <url>
```
```bash
go run .\main.go -port <port> -origin <url> clear-cache true/false
```

Donde la **url** para este proyecto se utilizo un generador de json falsos, para mas información presione [aqui](http://dummyjson.com)


--- 
### ⚙️ Uso
Desde la línea de comandos del sistema operativo, en este caso PowerShell, ejecuta:

```bash
curl -Headers @{ "X-Cache" = "MISS/HIT" } http://localhost:<port>/?url=products
```
Donde el **port** es el puerto ingresado por el usuario.

--- 
## Extras:

![image](https://github.com/user-attachments/assets/5cf568e9-814d-4741-aa29-4a3e8ed53ad5)

📢 Este proyecto fue inspirado en [roadmap.sh](https://roadmap.sh/projects/caching-server)



