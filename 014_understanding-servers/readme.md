# Programación Web - Términos sinónimos

- Router
- Router de solicitud
- Multiplexor
- Mux
- Servemux
- Servidor
- Router HTTP
- Router de solicitud HTTP
- Multiplexor HTTP
- Mux HTTP
- Servemux HTTP
- Servidor HTTP

---

# Solicitudes y respuestas

Las solicitudes y respuestas son similares. Ambos mensajes constan de:

- una línea de solicitud/respuesta
- cero o más líneas de encabezado (Header)
- una línea en blanco (es decir, CRLF por sí sola)
- un cuerpo de mensaje opcional (Body)

---

## Solicitud HTTP

Solicitud

- Línea de solicitud
- Encabezados (Headers)
- Cuerpo de mensaje opcional (Body)

Línea de solicitud

- Método SP URI-de-solicitud SP Versión-HTTP CRLF

Ejemplo de línea de solicitud:

- GET /api/v1/products HTTP/1.1

---

## Respuesta HTTP

Respuesta

- Línea de estado
- Encabezados (Headers)
- Cuerpo de mensaje opcional (Body)

Línea de estado

- Versión-HTTP SP Código-de-estado SP Frase-de-razón CRLF

Ejemplo de línea de estado:

- HTTP/1.1 200 OK

---

## Encabezados (Headers)

[Listado de campos de encabezado](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields)

---

## Inspección

- Puedes usar las herramientas de desarrollo de Google Chrome / red
- Puedes usar CURL en la línea de comandos

```
curl -v golang.org
```

---

# Documentación

## [Protocolo de Transferencia de Hipertexto (HTTP)](https://es.wikipedia.org/wiki/Protocolo_de_Transferencia_de_Hipertexto)

El desarrollo de HTTP fue iniciado por Tim Berners-Lee en el CERN en 1989. El desarrollo de los estándares de HTTP fue coordinado por el Grupo de Trabajo de Ingeniería de Internet (IETF) y el Consorcio World Wide Web (W3C), culminando en la publicación de una serie de Solicitudes de Comentarios (RFC).

La primera definición de HTTP/1.1, la versión de HTTP más utilizada, ocurrió en la RFC 2068 en 1997, aunque esta fue reemplazada por la RFC 2616 en 1999. En 2014, la RFC2616 fue reemplazada por múltiples RFCs (7230-7237).

Una versión posterior, el sucesor HTTP/2, fue estandarizada en 2015, y ahora es compatible con los principales servidores web.

## [Solicitud de Comentarios (RFC)](https://es.wikipedia.org/wiki/Solicitud_de_comentarios)

Una Solicitud de Comentarios (RFC) es un tipo de publicación del Grupo de Trabajo de Ingeniería de Internet (IETF) y la Sociedad de Internet (ISOC), los principales organismos de desarrollo técnico y establecimiento de estándares para Internet.

Un RFC es escrito por ingenieros y científicos informáticos en forma de un memorando que describe métodos, comportamientos, investigaciones o innovaciones aplicables al funcionamiento de Internet y sistemas conectados a Internet. Se somete a revisión por pares o simplemente para transmitir nuevos conceptos, información o (ocasionalmente) humor ingenieril. El IETF adopta algunas de las propuestas publicadas como RFCs como estándares de Internet.

Los documentos RFC fueron inventados por Steve Crocker en 1969 para ayudar a registrar notas no oficiales sobre el desarrollo de ARPANET. Desde entonces, los RFCs se han convertido en documentos oficiales de especificaciones de Internet, protocolos de comunicación, procedimientos y eventos.

## [Grupo de Trabajo de Ingeniería de Internet (IETF)](https://es.wikipedia.org/wiki/Grupo_de_Trabajo_de_Ingeniería_de_Internet)

El Grupo de Trabajo de Ingeniería de Internet (IETF) desarrolla y promueve estándares voluntarios para Internet, en particular los estándares que componen el conjunto de protocolos de Internet (TCP/IP). Es una organización de estándares abiertos, sin membresía formal ni requisitos de membresía. Todos los participantes y gerentes son voluntarios, aunque su trabajo generalmente es financiado por sus empleadores o patrocinadores.

El IETF comenzó como una actividad respaldada por el gobierno federal de los EE. UU., pero desde 1993 opera como una función de desarrollo de estándares bajo los auspicios de la Sociedad de Internet, una organización internacional sin fines de lucro basada en membresías.

## [Lista de RFCs](https://es.wikipedia.org/wiki/Lista_de_RFCs)

---

# RFC 7230

HTTP fue creado para la arquitectura de la World Wide Web (WWW) y ha evolucionado con el tiempo para satisfacer las necesidades de escalabilidad de un sistema mundial de hipertexto. Gran parte de esa arquitectura está reflejada en la terminología y producciones de sintaxis utilizadas para definir HTTP.

## 2.1. Mensajería Cliente/Servidor

HTTP es un protocolo sin estado de solicitud/respuesta que opera mediante el intercambio de mensajes a través de una **conexión** confiable de capa de **TRANSPORTE** o **SESIONES**.

Un "CLIENTE" HTTP es un programa que establece una **conexión** con un servidor con el propósito de enviar una o más solicitudes HTTP.

Un "SERVIDOR" HTTP es un programa que acepta **conexiones** para atender solicitudes HTTP enviando respuestas HTTP.

Los términos "cliente" y "servidor" se refieren solo a los roles que estos programas desempeñan para una conexión en particular. El mismo programa puede actuar como cliente en algunas conexiones y como servidor en otras. El término "agente de usuario" hace referencia a cualquier programa cliente que inicia una solicitud, incluidos (pero no limitados a) navegadores, arañas (robots basados en la web), herramientas de línea de comandos, aplicaciones personalizadas y aplicaciones móviles. El término "servidor de origen" hace referencia al programa que puede originar respuestas autoritarias para un recurso objetivo dado. Los términos "emisor" y "receptor" se refieren a cualquier implementación que envíe o reciba un mensaje dado, respectivamente.

HTTP depende del estándar de **Identificador Uniforme de Recurso (URI)** para indicar el recurso objetivo y las relaciones entre recursos. Los mensajes se transmiten en un formato similar al utilizado por el correo electrónico de Internet y las **Extensiones Multipropósito de Correo de Internet (MIME)** (consulta el Apéndice A de [RFC7231] para conocer las diferencias entre los mensajes HTTP y MIME).

La mayor parte de la comunicación HTTP consiste en una solicitud de recuperación (GET) para obtener una representación de algún recurso identificado por un URI. En el caso más simple, esto puede lograrse mediante una sola **conexión bidireccional** (===) entre el agente de usuario (UA) y el servidor de origen (O).

**Solicitud del cliente:**

```
GET /hello.txt HTTP/1.1

User-Agent: curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.7l zlib/1.2.3

Host: www.example.com

Accept-Language: en, mi
```

**Respuesta del servidor:**

```
HTTP/1.1 200 OK

Date: Mon, 27 Jul 2009 12:28:53 GMT

Server: Apache

Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT

ETag: "34aa387-d-1568eb00"

Accept-Ranges: bytes

Content-Length: 51

Vary: Accept-Encoding

Content-Type: text/plain

Hello World! My payload includes a trailing CRLF.
```

## 2.2. Diversidad de Implementación

Al considerar el diseño de HTTP, es fácil caer en la trampa de pensar que todos los agentes de usuario son navegadores de propósito general y que todos los servidores de origen son grandes sitios web públicos. No es así en la práctica. Los agentes de usuario HTTP comunes incluyen electrodomésticos, estéreos, básculas, scripts de actualización de firmware, programas de línea de comandos, aplicaciones móviles y dispositivos de comunicación en una multitud de formas y tamaños. De igual manera, los servidores de origen HTTP comunes incluyen unidades de automatización del hogar, componentes de redes configurables, máquinas de oficina, robots autónomos, feeds de noticias, cámaras de tráfico, selectores de anuncios y plataformas de entrega de videos.

El término "AGENTE DE USUARIO" no implica que haya un usuario humano interactuando directamente con el software del agente en el momento de la solicitud. En muchos casos, un agente de usuario está instalado o configurado para ejecutarse en segundo plano y guardar sus resultados para inspección posterior (o guardar solo un subconjunto de esos resultados que podrían ser interesantes o erróneos). Las arañas, por ejemplo, generalmente reciben un URI de inicio y están configuradas para seguir un comportamiento determinado mientras rastrean la web como un gráfico hipertextual.
