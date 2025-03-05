# Uso de Bucles `for` en Go

Diferentes usos del bucle `for` en Go:

## 1. **Bucle `for` con condición y actualización**

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
fmt.Println("Sum:", sum)
```

### **Descripción:**

Este es el uso más común del bucle `for`, similar a `for` en C, Java y JavaScript. Tiene tres partes:

1. **Inicialización (`i := 0`)** → Se ejecuta una sola vez antes de que empiece el bucle.
2. **Condición (`i < 10`)** → Evalúa si se debe seguir ejecutando el bucle.
3. **Actualización (`i++`)** → Modifica el valor de `i` después de cada iteración.

### **Casos de uso:**

- Iterar sobre un rango de valores numéricos.
- Realizar una acción un número fijo de veces.
- Recorrer arrays cuando se necesita acceso al índice.

---

## 2. **Bucle `for` con solo una condición (similar a `while`)**

```go
sum = 1
for sum < 1000 {
    sum++
}
fmt.Println(sum)
```

### **Descripción:**

Este `for` se ejecuta mientras la condición `sum < 1000` sea verdadera. Es equivalente a un `while` en otros lenguajes.

### **Casos de uso:**

- Cuando no se sabe cuántas iteraciones se necesitarán.
- Esperar una condición externa (por ejemplo, una entrada del usuario).
- Implementar ciclos con tiempos de espera.

---

## 3. **Bucle `for` infinito con `break`**

```go
sum = 0
for {
    if sum > 1000 {
        break
    }
    sum++
}
fmt.Println(sum)
```

### **Descripción:**

Este `for` no tiene ninguna condición de salida explícita. Se ejecutará indefinidamente hasta que se cumpla la condición dentro del bloque (`if sum > 1000 { break }`).

### **Casos de uso:**

- Servidores o procesos que deben ejecutarse indefinidamente.
- Bucle de juego en videojuegos.
- Programas que esperan eventos externos (como una señal de interrupción).

---

## 4. **Bucle `for` sobre un `slice` con índice**

```go
arr := []int{1, 2, 3, 1, 2, 3}
for i := range arr {
    fmt.Println("Index:", i, "- Value:", arr[i])
}
```

### **Descripción:**

Este `for` usa `range`, lo que permite recorrer una lista (slice) de manera sencilla.

- `i` representa el índice del elemento.
- `arr[i]` es el valor del elemento en ese índice.

### **Casos de uso:**

- Cuando se necesita tanto el índice como el valor.
- Recorrer `slices`, `arrays`, `maps` o `channels`.

---

## 5. **Bucle `for` sobre un `slice` ignorando el índice**

```go
for _, v := range arr {
    fmt.Println("Value:", v)
}
```

### **Descripción:**

- Se usa `_` para ignorar el índice y solo obtener el valor (`v`).
- Es útil cuando no se necesita el índice.

### **Casos de uso:**

- Recorrer listas sin necesidad de modificar los valores.
- Mejor legibilidad cuando no se usa el índice.

---

## 6. **Bucle `for` sobre un `map`**

```go
map2 := map[string]float64{
    "A": 12.3,
    "Z": 23.1,
    "C": 34,
}

for key, value := range map2 {
    fmt.Println("Key:", key, "Value:", value)
}
```

### **Descripción:**

Este `for` recorre un `map`, obteniendo:

- `key` → Clave del mapa.
- `value` → Valor asociado a la clave.

### **Casos de uso:**

- Leer datos almacenados en `maps`.
- Iterar sobre configuraciones o estructuras clave-valor.

---

## 7. **Bucle `for` anidado para recorrer un `map` con `slices` como valores**

```go
map3 := map[string][]int{
    "A": nil,
    "B": {2, 34, 1, 2, 4},
    "C": {4, 5, 3, 2, 1},
}

for key, value := range map3 {
    fmt.Println("Key:", key)
    for _, v := range value {
        fmt.Println("Value:", v)
    }
    fmt.Println()
}
```

### **Descripción:**

Este código tiene un `for` anidado:

1. **Primer `for`** → Itera sobre las claves (`key`) y los valores (`value`) del `map`.
2. **Segundo `for`** → Itera sobre los elementos dentro del `slice` de cada `key`.

### **Casos de uso:**

- Recorrer estructuras de datos complejas como `map[string][]int`.
- Procesar información jerárquica.
- Leer configuraciones anidadas.

---

## **Conclusión**

Go usa `for` como su única estructura de bucle, pero con distintas variaciones:

- **`for` con condición y actualización** → Para iteraciones con conteo definido.
- **`for` con solo condición (`while` en otros lenguajes)** → Para iteraciones sin número fijo.
- **Bucle infinito con `break`** → Para procesos continuos.
- **`for` con `range`** → Para recorrer `slices`, `arrays` y `maps`.
- **Bucles anidados** → Para estructuras de datos más complejas.

Cada variante se usa en escenarios específicos, optimizando legibilidad y eficiencia en Go. 🚀
