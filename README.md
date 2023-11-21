# puc-disciplinas

Repositorio para o projeto final do curso de Engenharia de Computação PUC-Rio.

O relatório completo está disponível no repositório [Leinadium/puc-tcc-disciplinas](https://github.com/Leinadium/puc-disciplinas)

## Como executar

É possível executar o sistema utilizando docker.

```sh
git clone https://github.com/Leinadium/puc-disciplinas

docker compose up
```

O sistema estará disponível na porta `5173`. A API está exposta na porta `8080`, e o banco na porta `5432`. Na primeira execução, espere a execução completa do container `sync-once`, que sincroniza a base de dados com o microhorario

## Licensa

MIT