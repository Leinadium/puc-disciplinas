from sqlalchemy.orm import Session
from sqlalchemy import insert, select, values
from sqlalchemy.exc import IntegrityError
from models import Historico, Usuario

from logic import Historia


def inserir_historico(session: Session, lista: list[Historia], cod_usuario: str):
    for x in lista:
        try:   
            session.add(Historico(
                cod_usuario=cod_usuario,
                cod_disciplina=x.disciplina,
                semestre=x.periodo,
                grau=x.grau
            ))
            session.commit()
        except IntegrityError:
            session.rollback()


def inserir_curriculo(session: Session, cod_usuario: str, cod_curriculo: str):
    session.query(
        Usuario
    ).filter(
        Usuario.cod_usuario == cod_usuario
    ).update(
        {Usuario.cod_curriculo: cod_curriculo}
    )


def remover_historico(session: Session, cod_usuario: str):
    session.query(
        Historico
    ).filter(
        Historico.cod_usuario == cod_usuario
    ).delete()
