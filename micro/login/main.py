import bs4
import flask
import requests


USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:98.0) Gecko/20100101 Firefox/98.0"

SESSAO_BASE = "VmluY3Vsbz1BJlNpc3RMb2dpbj1QVUNPTkxJTkVfQUxVTk8mQXBwTG9naW49TE9HSU4mRnVuY0xvZ" \
              "2luPUxPR0lOJlNpc3RNZW51PVBVQ09OTElORV9BTFVOTyZBcHBNZW51PU1FTlUmRnVuY01lbnU9TUVOVQ__"

URL_BASE = f"https://sau.puc-rio.br/WebLoginPucOnline/Default.aspx?sessao={SESSAO_BASE}"


app = flask.Flask(__name__)


@app.post("/")
def login():
    # pega o usuario e senha do form
    usuario = flask.request.form.get("usuario")
    senha = flask.request.form.get("senha")
    if not usuario or not senha:
        return {"message": "Usuário ou senha não informados"}, 400

    # faz o primeiro request
    try:
        r = requests.get(
            url=URL_BASE,
            headers={'User-Agent': USER_AGENT}
        )
    except requests.exceptions.ConnectionError:
        return {"message": "Não foi possível se conectar ao servidor"}, 500

    # pegando os cookies
    cookies = r.cookies.get_dict()

    # pegando propriedades do ASP.NET
    soup = bs4.BeautifulSoup(r.text, features='html.parser')
    view_state_generator: bs4.Tag = soup.find(id='__VIEWSTATEGENERATOR')
    event_validation: bs4.Tag = soup.find(id='__EVENTVALIDATION')
    view_state: bs4.Tag = soup.find(id='__VIEWSTATE')

    if view_state_generator is None or event_validation is None or view_state is None:
        return {"message": "Não foi possível encontrar os campos do ASP.NET"}, 500

    # faz o segundo request
    try:
        r = requests.post(
            url=URL_BASE,
            headers={
                'User-Agent': USER_AGENT,
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            cookies=cookies,
            data={
                "__EVENTTARGET": "",
                "__EVENTARGUMENT": "",
                "__VIEWSTATE": view_state.get("value"),
                "__VIEWSTATEGENERATOR": view_state_generator.get("value"),
                "__EVENTVALIDATION": event_validation.get("value"),
                "txtLogin": usuario,
                "txtSenha": senha,
                "btnOk": "Logar"
            }
        )
    except requests.exceptions.ConnectionError:
        return {"message": "Não foi possível se conectar ao SAU"}, 500

    # tenta pegar a tag do nome do usuario
    soup = bs4.BeautifulSoup(r.text, features='html.parser')
    nome_usuario: bs4.Tag = soup.find(id='lblNomeUsuario')
    if nome_usuario is None:
        return {"message": "Usuário ou senha incorretos"}, 404

    # tenta reduzir pro nome
    nome_usuario = nome_usuario.find("b")
    if nome_usuario is None:
        return {"message": "Nome do usuário não encontrado"}, 500

    return {"nome": nome_usuario.text.strip()}, 200


if __name__ == '__main__':
    app.run(debug=True)
