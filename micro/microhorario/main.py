# Serviço de sincronização do microhorario

from os import getenv
from argparse import ArgumentParser

from dotenv import load_dotenv
from sqlalchemy.orm import Session
from sqlalchemy.engine import Engine
from microhorario_dl import Microhorario

from models import Base
from models import Departamento, Disciplina, Prerequisitos, Professor, Turma
from models import Horario, Alocacao


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
    from sqlalchemy import delete, text

    with Session(eng) as session:
        if is_full:
            print("[ADICIONA_NO_BANCO] Deletando os dados antigos")
            session.execute(text(f'TRUNCATE TABLE {Departamento.__tablename__} CASCADE'))
            session.execute(text(f'TRUNCATE TABLE {Disciplina.__tablename__} CASCADE'))
            session.execute(text(f'TRUNCATE TABLE {Prerequisitos.__tablename__} CASCADE'))
            session.execute(text(f'TRUNCATE TABLE {Professor.__tablename__} CASCADE'))
            session.execute(text(f'TRUNCATE TABLE {Turma.__tablename__} CASCADE'))
            session.commit()

            # adiciona os departamentos
            print("[ADICIONA_NO_BANCO] Adicionando os departamentos")
            for d in m.departamentos:
                session.add(Departamento(
                    cod_depto=d.codigo,
                    nome_depto=d.nome
                ))

            # cria o set de professores
            print("[ADICIONA_NO_BANCO] Adicionando os professores")
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

            print("[ADICIONA_NO_BANCO] Adicionando os disciplinas")
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

            print("[ADICIONA_NO_BANCO] Adicionando os prerequisitos")
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

            print("[ADICIONA_NO_BANCO] Adicionando as turmas")
            for disc in m.disciplinas:
                for turma in disc.turmas:
                    session.add(Turma(
                        cod_disciplina=disc.codigo,
                        cod_turma=turma.codigo,
                        nome_professor=turma.professor,
                        shf=turma.shf
                    ))

            session.commit()

        # deleta os horarios e alocacoes
        print("[ADICIONA_NO_BANCO] Deletando as alocacoes e horarios")
        session.execute(text(f'TRUNCATE TABLE {Alocacao.__tablename__} CASCADE'))
        session.execute(text(f'TRUNCATE TABLE {Horario.__tablename__} CASCADE'))

        # adiciona as alocacoes para cada turma
        print("[ADICIONA_NO_BANCO] Adicionando as alocacaoes")
        for disc in m.disciplinas:
            for turma in disc.turmas:
                session.add(Alocacao(
                    cod_disciplina=disc.codigo,
                    cod_turma=turma.codigo,
                    destino=f"{turma.alocacao.destino.codigo} ({turma.alocacao.destino.nome})",
                    vagas=turma.alocacao.vagas,
                ))

        session.commit()

        # adiciona os horarios para cada turma
        print("[ADICIONA_NO_BANCO] Adicionando os horarios")
        for disc in m.disciplinas:
            for turma in disc.turmas:
                # uma mesma turma pode ter mais de um horario igual
                # devido tbm a sua alocação
                buffer: dict[str, Horario] = {}
                # por isso, joga no buffer, vendo sua unicidade
                for horario in turma.horarios:
                    dia = horario.dia if horario is not None else "xxx"
                    inicio = horario.inicio if horario is not None else 0
                    fim = horario.fim if horario is not None else None
                    key = f"{dia}-{inicio}-{fim}"
                    if key not in buffer:
                        buffer[key] = Horario(
                            cod_disciplina=disc.codigo,
                            cod_turma=turma.codigo,
                            dia=dia,
                            hora_inicio=inicio,
                            hora_fim=fim
                        )
                # depois de tudo isso, adiciona no banco
                for horario in buffer.values():
                    session.add(horario)

        session.commit()
    return


def main(full: bool = False):
    from sqlalchemy import create_engine
    conn_str = getenv("POSTGRES_CONN")

    # conecta ao banco
    print("[MAIN] Conectando no banco")
    engine = create_engine(conn_str, echo=False)

    # cria as tabelas
    print("[MAIN] Criando as tabelas")
    configura_banco(engine)

    # baixa os dados
    print("[MAIN] Baixando o microhorario")
    m = baixa_microhorario(full=full)

    # adiciona os dados
    print("[MAIN] Adicionando os dados")
    adiciona_no_banco(m, engine, is_full=full)


if __name__ == "__main__":
    args = parser.parse_args()
    main(args.full)
