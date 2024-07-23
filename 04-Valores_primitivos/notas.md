# Valores primitivos
Pueden o no declararse explicitamente, pero hacerlo mejora el rendimiento al ahorrar el trabajo de averiguar le tipo de dato.

## Enteros

- int: Depende del OS 32 o 64
- int8: de -128 a 127
- int 16: de -2^15 a 2^15 -1
- int 32: de -2^31 a 2^31 -1
- int 64: de -2^63 a 2^63 -1

## Enteros positivos

- uint: Depende del OS 32 o 64
- uint8: de 0 a 2^8 -1
- uint 16: de 0 a 2^16 -1
- uint 32: de 0 a 2^32 -1
- uint 64: de 0 a 2^64 -1


## Decimales

- float32: 32 bits = +/- 1.18x10^-38 a +/- 3.4x10^38
- float64: 64 bits = +/- 2.23x10^-308 a +/- 1.8x10^308


## Textos y bool

- string = ""
- bool \in {true, false}



## Complejos

- Complex64 = Real e imaginario float32
- Comlex128 = Real e imaginario float64
- Ejemplo: c := 10 + 8i

