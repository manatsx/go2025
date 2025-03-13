# Counting Valleys

## Descripción

Un excursionista entusiasta lleva un registro meticuloso de sus caminatas. Durante su última caminata, que constó de un número exacto de pasos, registró si cada paso fue cuesta arriba (`U`) o cuesta abajo (`D`).

Las caminatas siempre comienzan y terminan al nivel del mar, y cada paso representa un cambio de una unidad en la altitud.

### Definiciones:

- Una **montaña** es una secuencia de pasos consecutivos por encima del nivel del mar, que comienza con un paso hacia arriba (`U`) desde el nivel del mar y termina con un paso hacia abajo (`D`) hasta el nivel del mar.
- Un **valle** es una secuencia de pasos consecutivos por debajo del nivel del mar, que comienza con un paso hacia abajo (`D`) desde el nivel del mar y termina con un paso hacia arriba (`U`) hasta el nivel del mar.

Dado un número de pasos y una secuencia de movimientos en la caminata, se debe determinar cuántos valles fueron atravesados.

---

## Ejemplo

### Entrada:

```
steps = 8
path = "UDDDUDUU"
```

### Proceso:

El excursionista primero baja en un valle, desciende 1 unidad y luego sube de nuevo al nivel del mar. Luego, sube a una montaña y desciende nuevamente hasta el nivel del mar.

### Salida:

```
1
```

---

## Descripción de la Funcionalidad

Se debe implementar la función `countingValleys`, que recibe los siguientes parámetros:

### **Parámetros:**

- `steps` (_int_): El número total de pasos en la caminata.
- `path` (_string_): Una cadena de caracteres que representa la secuencia de pasos (`U` para subida y `D` para bajada).

### **Salida:**

- Retorna un **entero** que representa el número de valles recorridos durante la caminata.

---

## Formato de Entrada

1. La primera línea contiene un entero `steps`, que indica el número de pasos en la caminata.
2. La segunda línea contiene una cadena `path` de longitud `steps`, con caracteres `U` (subida) y `D` (bajada).

---

## Formato de Salida

- Un entero que indica el número total de valles recorridos.

---

## Restricciones

- `2 <= steps <= 10^6`
- `path[i]` pertenece al conjunto `{U, D}`

---

## Ejemplo de Entrada y Salida

### Entrada 1:

```
12
DDUUDDUDUUUD
```

### Salida 1:

```
2
```

### Explicación:

El excursionista baja a un valle (1) y sube de regreso. Luego, vuelve a bajar a otro valle (2) y sube nuevamente.

---

## Notas Adicionales

- La caminata siempre comienza y termina en el nivel del mar.
- Se deben contar los valles completos, es decir, aquellos en los que se baja por debajo del nivel del mar y se vuelve a subir a él.
- Se asume que la entrada es válida según las restricciones dadas.
