# Laços - Loop For

Laços ou Loops são estruturas de controle importantes e presentes em todas as linguagens de programação. Em Go não é diferente, porém, devido a sua simplicidade, Go possue apenas o laço `for`. Em Go, o laço `for` é extremamente flexível e é capaz de substituir facilmente as diversas estruturas existentes em outras linguangens.

O laço `for` tradicional, pode ser escrito da seguinte forma:

```go
for i := 0; i < 10; i++ {
	fmt.Println("Qual é o valor de i?", i)
} 
```

Já um laço `while`, pode ser simulado da seguinte forma usando o `for`:

```go
valor := 0
teste := true
for teste {
    valor++
    if valor%5 == 0 {
        teste = false
    }
    fmt.Println("O número é:", valor)
}
```

O `for` em Go é capaz de avaliar uma expressão booleana e executar até que a expressão seja avaliada como falsa.

Também é possível simular um laço infinito e utilizar a palavra chave `break` para sair do laço:

```go
for {
    valor--
    fmt.Println("O número é:", valor)
    if valor == 0 {
        break
    }
}
```

Go também disponibiliza um `for` especial capaz de interar sobre `arrays` e `slices`. Esta estrutura utiliza a palavra chave `range`:

```go
texto := "Eu ado escrever prdogramas em Go"
for indice, letra := range texto {
    fmt.Printf("Texto[%d] = %q\r\n", indice, letra)
}
```

OBS.: Uma string em Go é um array de bytes.
