Claro, aquí tienes el **README corregido** para el problema de Hackerrank:

---

# Electronics Shop

## Descripción del problema

El objetivo de este ejercicio es determinar cuál es la combinación más cara de un teclado y una memoria USB que una persona puede comprar con un presupuesto dado. Se te proporcionan dos listas de precios: una para teclados y otra para memorias USB. La tarea es encontrar la combinación de un teclado y una memoria USB cuyo precio total sea lo más alto posible sin exceder el presupuesto. Si no es posible comprar ambos productos dentro del presupuesto, se debe devolver un valor de `-1`.

## Función

La función que debe ser completada es `getMoneySpent`, la cual tiene los siguientes parámetros:

- **keyboards[n]**: Un array que contiene los precios de los teclados disponibles.
- **drives[m]**: Un array que contiene los precios de las memorias USB disponibles.
- **b**: El presupuesto disponible para realizar la compra.

### Parámetros:

- **keyboards[n]**: Un array de enteros que representa los precios de los teclados (con `n` teclados).
- **drives[m]**: Un array de enteros que representa los precios de las memorias USB (con `m` memorias USB).
- **b**: El presupuesto, un entero que limita el gasto total para comprar tanto un teclado como una memoria USB.

### Retorno:

La función debe devolver un único entero:

- El valor máximo que se puede gastar en la combinación de un teclado y una memoria USB sin superar el presupuesto.
- Si no es posible comprar ambos productos dentro del presupuesto, la función debe devolver `-1`.

## Formato de entrada

El formato de entrada es el siguiente:

1. La primera línea contiene tres números enteros separados por espacios: el presupuesto `b`, el número de teclados `n` y el número de memorias USB `m`.
2. La segunda línea contiene `n` enteros separados por espacios que representan los precios de los teclados.
3. La tercera línea contiene `m` enteros separados por espacios que representan los precios de las memorias USB.

### Ejemplo de entrada

```
10 2 3
3 1
5 2 8
```

### Explicación del ejemplo:

- El presupuesto es 10.
- Hay 2 teclados, con precios de 3 y 1.
- Hay 3 memorias USB, con precios de 5, 2 y 8.

La combinación más cara que se puede comprar es un teclado que cuesta 1 y una memoria USB que cuesta 8, lo que da un total de 9, que es el valor más cercano al presupuesto sin excederlo.

## Ejemplo de salida

```
9
```

## Caso en que no sea posible comprar ambos productos

En algunos casos, no será posible comprar tanto un teclado como una memoria USB debido a que la combinación de precios excede el presupuesto. En estos casos, la función debe devolver `-1`.

### Ejemplo de entrada

```
5 1 1
4
5
```

### Explicación del ejemplo:

- El presupuesto es 5.
- Hay 1 teclado con un precio de 4.
- Hay 1 memoria USB con un precio de 5.

El total sería 4 + 5 = 9, lo que excede el presupuesto de 5. Por lo tanto, no es posible realizar la compra, y la salida debe ser `-1`.

## Restricciones

- El precio de cada artículo (teclado y memoria USB) estará en el rango de 1 a 1000 (inclusive).
- El presupuesto `b` estará en el rango de 1 a 1000 (inclusive).
