# Problema: Plus Minus

## Descripción

Se te da un arreglo de enteros. Tu tarea es calcular y mostrar las proporciones de los valores positivos, negativos y cero en el arreglo. Devuelve las proporciones como valores decimales, redondeados a 6 decimales.

### **Firma de la función:**

```go
func plusMinus(arr []int32)
```

### **Entrada:**

- **arr (1 ≤ len(arr) ≤ 100)**: Un arreglo de enteros `arr` de longitud `n`. Los elementos del arreglo son enteros entre `-100` y `100`.

### **Salida:**

Imprime las siguientes 3 líneas:

1. La proporción de números positivos en el arreglo.
2. La proporción de números negativos en el arreglo.
3. La proporción de números cero en el arreglo.

Cada proporción debe ser impresa en una nueva línea con 6 dígitos después del punto decimal.

### **Restricciones:**

- El arreglo `arr` siempre tendrá al menos un elemento.
- El arreglo contendrá una mezcla de números positivos, negativos y ceros.

### **Ejemplo:**

#### Entrada:

```
6
1 1 0 -1 -1 0
```

#### Salida:

```
0.333333
0.333333
0.333333
```

#### Explicación:

- Hay `2` números positivos (`1, 1`), `2` números negativos (`-1, -1`), y `2` ceros (`0, 0`).
- La proporción de números positivos es `2/6 = 0.333333`.
- La proporción de números negativos es `2/6 = 0.333333`.
- La proporción de números cero es `2/6 = 0.333333`.

### **Nota:**

- La salida debe ser impresa con una precisión de 6 decimales, incluso si el número es un número entero.

---
