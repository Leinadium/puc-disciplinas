from requests import post, get
from requests.cookies import RequestsCookieJar

from typing import TypedDict


class Escolhas(TypedDict):
    disciplina: str
    turma: str


class Grade(TypedDict):
    escolhas: list[Escolhas]


class Resposta(TypedDict):
    cod: str
    val: float


def executar_login(usuario: str, senha: str) -> RequestsCookieJar:
    url = "http://localhost:8080/login"
    data = {"usuario": usuario, "senha": senha}
    response = post(url, data=data)

    if response.status_code == 200:
        return response.cookies
    else:
        print(response.status_code, response.text)
        raise Exception("Erro ao fazer login")


def pegar_recomendacoes(grade: Grade, cookies: RequestsCookieJar) -> list[Resposta]:
    url = "http://localhost:8080/recomendacao"
    response = get(url, json=grade, cookies=cookies)

    if response.status_code == 200:
        return response.json()['data']
    else:
        print(response.status_code, response.text)
        raise Exception("Erro ao pegar recomendações")


def filtra_resposta(resposta: list[Resposta]) -> list[Resposta]:
    resposta.sort(key=lambda x: x['val'], reverse=True)
    return resposta[:20]


def main():
    from getpass import getpass
    u = input('Usuario: ')
    p = getpass(prompt='Senha: ')

    cookies = executar_login(u, p)
    grade = {
        "escolhas": [
            {"disciplina": "INF1350", "turma": "3WA"}
        ]
    }

    resposta = pegar_recomendacoes(grade, cookies)
    resposta = filtra_resposta(resposta)

    for x in resposta:
        print(x)


if __name__ == "__main__":
    main()
