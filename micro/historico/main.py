from os import getenv
from dotenv import load_dotenv
from sqlalchemy.orm import Session
from sqlalchemy.exc import IntegrityError
from sqlalchemy.engine import create_engine
from flask import Flask, request, render_template

from logic import parse_historico
from database import inserir_historico
from database import inserir_curriculo
from database import remover_historico


load_dotenv()
conn_str = getenv("POSTGRES_CONN")

engine = create_engine(conn_str, echo=True)

app = Flask(__name__)
app.config['MAX_CONTENT_LENGTH'] = 1 * 1024 * 1024


@app.get("/")
def index():
    return render_template("exemplo.html")


@app.post("/")
def submit():
    usuario = request.args.get("usuario")
    if usuario is None:
        print('usuario nao informado')
        return {"message": "Usuário não informado"}, 400

    # vendo se o arquivo foi submetido
    file = request.files.get('historico')
    if file is None:
        print('arquivo nao enviado')
        return {"message": "Arquivo não enviado"}, 400

    # lendo o arquivo
    try:
        file.stream.seek(0)
        conteudo: bytes = file.stream.read()
    except Exception as e:
        return {"message": f"Erro ao ler o arquivo: {e}"}, 400

    # processando o arquivo
    try:
        lista, cod_curriculo = parse_historico(conteudo)
    except Exception as e:
        print("erro ao processar o arquivo" + e)
        return {"message": f"Erro ao processar o arquivo: {e}"}, 400

    # inserindo no banco
    with Session(engine) as session:
        try:
            remover_historico(session, cod_usuario=usuario)
            inserir_historico(session, lista=lista, cod_usuario=usuario)
            session.commit()
        except Exception as e:
            session.rollback()
            print('erro ao inserir no banco', e)
            return {"message": f"Erro ao inserir no banco: {e}"}, 500

        try:
            if cod_curriculo is not None:
                inserir_curriculo(session, cod_usuario=usuario, cod_curriculo=cod_curriculo)
                session.commit()
        except IntegrityError:
            cod_curriculo = None

    print('sucesso')
    return {"inseridos": len(lista), "curriculo": cod_curriculo}, 200


if __name__ == "__main__":
    app.run(debug=True, port=5001)
