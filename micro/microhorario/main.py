# Serviço de sincronização do microhorario

from os import getenv
from argparse import ArgumentParser

from dotenv import load_dotenv
from sqlalchemy.orm import Session
from sqlalchemy.engine import Engine
from microhorario_dl import Microhorario

from models import Base
from models import Departamento, Disciplina, Horario, Prerequisitos, Professor, Turma, Vaga


load_dotenv()

parser = ArgumentParser(
    prog="Serviço de sincronização do microhorario",
    description="Sinconiza o microhorario com o banco de dados"
)
parser.add_argument("-f", "--full",
                    action="store_true",
                    help="Baixa todos os dados do microhorario novamente")


def baixa_microhorario(full: bool = False) -> Microhorario:
    """Baixa todos os dados do microhorario"""
    m: Microhorario = Microhorario.download()
    if full:
        # m.coletar_extra()
        pass
    return m


def configura_banco(eng: Engine):
    """Cria as tabelas necessarias para instalar o microhorario"""
    Base.metadata.create_all(eng)


def adiciona_no_banco(m: Microhorario, eng: Engine, is_full: bool = False):
    from sqlalchemy import delete

    with Session(eng) as session:
        if is_full:
            session.execute(delete(Departamento))
            session.execute(delete(Disciplina))
            session.execute(delete(Prerequisitos))
            session.execute(delete(Professor))
            session.execute(delete(Turma))

            # adiciona os departamentos
            for d in m.departamentos:
                session.add(Departamento(
                    cod_depto=d.codigo,
                    nome_depto=d.nome
                ))

            # cria o set de professores
            professores = set()
            for disc in m.disciplinas:
                for turma in disc.turmas:
                    professores.add(turma.professor)

            # itera sobre os professores
            for prof in professores:
                session.add(Professor(
                    nome_professor=prof
                ))

            # atualiza o banco para poder adicionar o resto
            session.commit()

            # adiciona as disciplinas
            for d in m.disciplinas:
                session.add(Disciplina(
                    cod_disciplina=d.codigo,
                    cod_depto=d.departamento.codigo,
                    nome_disciplina=d.nome,
                    ementa=d.ementa,
                    creditos=d.creditos
                ))

            # atualiza o banco para poder adicionar o resto
            session.commit()

            # adiciona os prerequisitos
            for d in m.disciplinas:
                # aborta caso prerequisitos seja None
                if d.prerequisitos is None:
                    continue
                for p in d.prerequisitos:
                    session.add(Prerequisitos(
                        cod_disc_orig=d.codigo,
                        cod_disc_depen=p
                    ))

            # atualiza o banco para poder adicionar o resto
            session.commit()

            # adiciona as turmas
            for disc in m.disciplinas:
                for turma in disc.turmas:
                    session.add(Turma(
                        cod_disciplina=disc.codigo,
                        cod_turma=turma.codigo,
                        nome_professor=turma.professor,
                        shf=turma.shf
                    ))

            session.commit()

        # deleta os horarios e vagas
        session.execute(delete(Vaga))
        session.execute(delete(Horario))

        session.commit()


def main(full: bool = False):
    from sqlalchemy import create_engine
    conn_str = getenv("POSTGRES_CONN")

    # conecta ao banco
    engine = create_engine(conn_str, echo=True)

    # cria as tabelas
    configura_banco(engine)

    # baixa os dados
    m = baixa_microhorario(full=full)

    # adiciona os dados
    adiciona_no_banco(m, engine, is_full=full)


if __name__ == "__main__":
    args = parser.parse_args()
    main(args.full)
