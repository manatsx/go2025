# HandlerFunc

[http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)

```Go
type HandlerFunc func(ResponseWriter, *Request)
```

```Go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

**Esto es solo un dato interesante para conocer. Probablemente no lo usarías en código de producción.**

---

## Pregunta

¿Podrías hacer que `http.Handle` acepte una función con esta firma?  
`func(ResponseWriter, *Request)`
