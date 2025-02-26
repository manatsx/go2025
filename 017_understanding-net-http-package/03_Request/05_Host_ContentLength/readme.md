Los **headers** en HTTP son **metadatos** que se envían en una solicitud (request) o en una respuesta (response). Contienen información sobre el cliente, el servidor, la autenticación, el formato de los datos, entre otros.

---

## **📌 Tipos de Headers y Ejemplos**

Los headers se pueden clasificar en varias categorías:

### **1️⃣ Headers de Autenticación y Seguridad**

Se usan para autenticar usuarios y proteger datos.

| Header             | Descripción                                  | Ejemplo                                          |
| ------------------ | -------------------------------------------- | ------------------------------------------------ |
| `Authorization`    | Envía credenciales (Token, Basic Auth)       | `Authorization: Bearer eyJhbGciOiJIUzI1...`      |
| `WWW-Authenticate` | Indica el esquema de autenticación requerido | `WWW-Authenticate: Basic realm="User Login"`     |
| `Cookie`           | Envía cookies almacenadas en el cliente      | `Cookie: sessionId=abc123`                       |
| `Set-Cookie`       | Configura cookies desde el servidor          | `Set-Cookie: sessionId=xyz456; HttpOnly; Secure` |
| `X-CSRF-Token`     | Token para prevenir ataques CSRF             | `X-CSRF-Token: f23af23a...`                      |

---

### **2️⃣ Headers de Control de Caché**

Sirven para indicar cómo deben manejarse las respuestas en caché.

| Header          | Descripción                                  | Ejemplo                                              |
| --------------- | -------------------------------------------- | ---------------------------------------------------- |
| `Cache-Control` | Controla la caché del navegador              | `Cache-Control: no-cache, no-store, must-revalidate` |
| `Pragma`        | Versión antigua de `Cache-Control`           | `Pragma: no-cache`                                   |
| `Expires`       | Fecha de expiración de la respuesta en caché | `Expires: Wed, 21 Oct 2025 07:28:00 GMT`             |

---

### **3️⃣ Headers de Contenido**

Especifican el tipo y la codificación de los datos enviados.

| Header                | Descripción                       | Ejemplo                                                |
| --------------------- | --------------------------------- | ------------------------------------------------------ |
| `Content-Type`        | Tipo de contenido enviado         | `Content-Type: application/json`                       |
| `Content-Length`      | Tamaño del cuerpo en bytes        | `Content-Length: 348`                                  |
| `Content-Encoding`    | Especifica la compresión de datos | `Content-Encoding: gzip`                               |
| `Content-Disposition` | Controla la descarga de archivos  | `Content-Disposition: attachment; filename="file.pdf"` |

---

### **4️⃣ Headers de Cliente (User-Agent)**

Proporcionan información sobre el cliente que realiza la solicitud.

| Header            | Descripción                                            | Ejemplo                                     |
| ----------------- | ------------------------------------------------------ | ------------------------------------------- |
| `User-Agent`      | Identifica el navegador o cliente HTTP                 | `User-Agent: Mozilla/5.0 (Windows NT 10.0)` |
| `Referer`         | Indica la URL de la que proviene la solicitud          | `Referer: https://example.com/login`        |
| `Origin`          | Indica el origen de la solicitud (importante en CORS)  | `Origin: https://example.com`               |
| `Accept`          | Especifica qué formatos de respuesta acepta el cliente | `Accept: application/json, text/html`       |
| `Accept-Encoding` | Especifica qué compresiones soporta el cliente         | `Accept-Encoding: gzip, deflate`            |

---

### **5️⃣ Headers de Control de Conexión**

Definen el manejo de la conexión entre cliente y servidor.

| Header       | Descripción                                                                | Ejemplo                          |
| ------------ | -------------------------------------------------------------------------- | -------------------------------- |
| `Connection` | Controla si la conexión debe mantenerse abierta                            | `Connection: keep-alive`         |
| `Keep-Alive` | Define parámetros de persistencia de conexión                              | `Keep-Alive: timeout=5, max=100` |
| `Upgrade`    | Indica que la conexión debe actualizarse a otro protocolo (Ej. WebSockets) | `Upgrade: websocket`             |

---

### **6️⃣ Headers de CORS (Cross-Origin Resource Sharing)**

Se usan para permitir o restringir solicitudes entre dominios diferentes.

| Header                         | Descripción                        | Ejemplo                                                     |
| ------------------------------ | ---------------------------------- | ----------------------------------------------------------- |
| `Access-Control-Allow-Origin`  | Indica qué dominios pueden acceder | `Access-Control-Allow-Origin: *`                            |
| `Access-Control-Allow-Methods` | Métodos HTTP permitidos            | `Access-Control-Allow-Methods: GET, POST, PUT`              |
| `Access-Control-Allow-Headers` | Headers permitidos en la solicitud | `Access-Control-Allow-Headers: Authorization, Content-Type` |

---

## **📌 Ejemplo Completo de Request y Response con Headers**

### **📤 Ejemplo de Request (Cliente → Servidor)**

```http
POST /api/login HTTP/1.1
Host: example.com
Authorization: Bearer eyJhbGciOiJIUzI1...
Content-Type: application/json
User-Agent: Mozilla/5.0 (Windows NT 10.0)
Accept: application/json
```

### **📥 Ejemplo de Response (Servidor → Cliente)**

```http
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: sessionId=xyz456; HttpOnly; Secure
Access-Control-Allow-Origin: *
Cache-Control: no-cache
Content-Length: 123
```

---

## **🔹 Resumen**

✔ Los **headers** contienen **metadatos** sobre la solicitud y respuesta.  
✔ Se dividen en **autenticación, control de caché, contenido, cliente, conexión y CORS**.  
✔ Son clave en **seguridad, rendimiento y compatibilidad** en aplicaciones web.

🔹 ¿Necesitas un header específico para una implementación en Go? 🚀
