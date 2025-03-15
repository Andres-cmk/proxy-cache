# 🚀 Caching Proxy

## 📌 Descripción
`Caching Proxy` es un servidor proxy con caché implementado en **Go**, que almacena las respuestas de las solicitudes HTTP para mejorar el rendimiento y reducir la carga en el servidor de origen.

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
 git clone 
 cd caching-proxy
