# puc-disciplinas

Repositorio para o projeto final do curso de Engenharia de Computação PUC-Rio.

Esse projeto consiste em um sistema, chamado **Macro**, que permite que um aluno da PUC-Rio simule a sua matrícula na universidade, e que receba *recomendações* de disciplinas para fazer, com base nos seus dados e nos de outros alunos. O sistema também permite que o aluno avalie disciplinas e professores.

O relatório completo está disponível no repositório
[Leinadium/puc-tcc-disciplinas](https://github.com/Leinadium/puc-disciplinas)

## Funcionamento

O sistema consiste em um frontend (pasta `front`) escrito em Svelte, que interage com um backend (pasta `api`) desenvolvido em Go. Algumas funcionalidades foram projetadas como microserviços, como é o caso do login com o portal da universadade, e o processamento do histórico. Esses microserviços (pasta `micro`) foram escritos em Python. Por fim, existe um serviço (pasta `services`) que é executado em intervalos regulares para sincronizar as disciplinas com o [microhorario](www.puc-rio.br/microhorario).

No total, o sistema consiste em 5 serviços separados: `front`, `api`, `micro/login`, `micro/historico`, `services/microhorario`, e um banco de dados `data`.

## Como executar

É possível executar o sistema utilizando docker.

```sh
git clone https://github.com/Leinadium/puc-disciplinas

docker compose up
```

O sistema estará disponível na porta `5173`. A API está exposta na porta `8080`, e o banco na porta `5432`. Na primeira execução, espere a execução completa do container `sync-once`, que sincroniza a base de dados com o microhorario.

## Licensa

MIT
