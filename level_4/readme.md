# Enunciados level 4

## Ex_1

Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.

## Ex_2

- Usando uma literal composta:
  - Crie uma slice de tipo int
  - Atribua 10 valores a ela

## Ex_3

- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
- Do quinto ao último item do slice (incluindo o último item!)
- Do segundo ao sétimo item do slice (incluindo o sétimo item!)
- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

## Ex_4

Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.

## Ex_5

- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]

## Ex_6

- Crie uma slice usando make que possa conter todos os estados do Brasil.
    - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.*

## Ex_7

- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
- "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

## Ex_8

Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.

## Ex_9

- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.

## Ex_10

- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.