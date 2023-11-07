package main

import(
    "fmt"
    "reflect"
)

func main() {
    /* Código original
    var nome string = "Reuben"
    var idade int = 10
    var versao float32 = 1.0
    */

    /*Inferência de tipos. Go pode identificar o tipo através do valor que atribuimos para as variáveis
    var nome = "Reuben"
    var idade = 10
    var versao = 1.0 //Nesse caso ele vai inferir float64 (que é o maior) mas pode ser que a gente quisesse float32 então é importante usar o correto.
    */

    //Declaração ainda mais curta
    nome := "Reuben"
    idade := 10
    versao := 1.0
    fmt.Println("Olá Sr. ",nome,". Sua idade é ",idade)
    fmt.Println("A versão desse programa é: ",versao)

    fmt.Println("O tipo de variável versão é ", reflect.TypeOf(versao))

    fmt.Println("1 - Iniciar monitoramento")
    fmt.Println("2 - Exibir logs")
    fmt.Println("0 - Sair")

    var comando int
    //fmt.Scanf("%d", &comando) //Através do %d indica que só pode receber inteiros. O "&" indica que é um ponteiro para a variável. É nosso ByRef do VB.
    fmt.Scan(&comando) // Receberia qualquer coisa. Caso o usuário digitasse algo diferente de inteiro a variável seguiria sendo 0 que é o valor inicializado dela.

    fmt.Println("O comando escolhido foi", comando)   

    /* Usando IF
    if comando == 1 {
        fmt.Println("Monitorando...")
    } else if comando == 2 {
        fmt.Println("Exibindo logs...")
    } else if comando == 0 {
        fmt.Println("Saindo do programa...")
    } else {
        fmt.Println("Comando não reconhecido")
    }
    */

    //Não precisa de break pois o Go só executa o primeiro switch que satisfizer a condição
    switch(comando){
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo logs...")
    case 0:
        fmt.Println("Saindo...")
    default:
        fmt.Println("Comando não reconhecido")
    }
}