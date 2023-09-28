from sqlalchemy import ForeignKey
from sqlalchemy import String
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped, mapped_column
from sqlalchemy.types import ARRAY


class Base(DeclarativeBase):
    pass


class Departamento(Base):
    __tablename__ = "departamentos"

    cod_depto: Mapped[str] = mapped_column("cod-depto", String(3), primary_key=True)
    nome_depto: Mapped[str] = mapped_column("nome-depto", String(100), nullable=False)

    def __repr__(self):
        return f"<Departamento {self.cod_depto}>"


class Disciplina(Base):
    __tablename__ = "disciplinas"

    cod_disciplina: Mapped[str] = mapped_column("cod-disciplina", primary_key=True)
    cod_depto = mapped_column("cod-depto",
                              ForeignKey("departamentos.cod-depto"), nullable=False)
    nome_disciplina: Mapped[str] = mapped_column("nome-disciplina", String(100), nullable=False)
    ementa: Mapped[str] = mapped_column("ementa", String(1000), nullable=True)
    creditos: Mapped[int] = mapped_column("creditos", nullable=False)

    def __repr__(self):
        return f"<Disciplina {self.cod_disciplina}>"


class Prerequisitos(Base):
    __tablename__ = "prerequisitos"

    cod_disc_orig: Mapped[str] = mapped_column("cod-disc-orig",
                                               ForeignKey("disciplinas.cod-disciplina"), primary_key=True)
    cod_disc_depen = mapped_column("cod-disc-depen", ARRAY(String(7)), primary_key=True)


class Professor(Base):
    __tablename__ = "professores"

    nome_professor: Mapped[str] = mapped_column("nome-professor", String(100), nullable=False,
                                                primary_key=True)


class Turma(Base):
    __tablename__ = "turmas"

    cod_turma: Mapped[str] = mapped_column("cod-turma", String(3), primary_key=True)
    cod_disciplina: Mapped[str] = mapped_column("cod-disciplina",
                                                ForeignKey("disciplinas.cod-disciplina"), primary_key=True)
    nome_professor: Mapped[str] = mapped_column("nome-professor",
                                                ForeignKey("professores.nome-professor"), nullable=False)
    shf: Mapped[int] = mapped_column("shf", nullable=False)

    __mapper_args__ = {
        "primary_key": [cod_turma, cod_disciplina]
    }


class Horario(Base):
    __tablename__ = "horarios"

    cod_disciplina: Mapped[str] = mapped_column("cod-disciplina", primary_key=True)
    # ForeignKey("turmas.cod-disciplina"), primary_key=True)
    cod_turma: Mapped[str] = mapped_column("cod-turma", primary_key=True)
    # ForeignKey("turmas.cod-turma"), primary_key=True)
    dia: Mapped[str] = mapped_column("dia", String(3), primary_key=True)
    hora_inicio: Mapped[int] = mapped_column("hora-inicio", nullable=False)
    hora_fim: Mapped[int] = mapped_column("hora-fim", nullable=False)


class Vaga(Base):
    __tablename__ = "vagas"

    cod_disciplina: Mapped[str] = mapped_column("cod-disciplina", primary_key=True)
    # ForeignKey("disciplinas.cod-disciplina"), primary_key=True)
    cod_turma: Mapped[str] = mapped_column("cod-turma", primary_key=True)
    # ForeignKey("turmas.cod-turma"), primary_key=True)
    destino: Mapped[str] = mapped_column("destino", String(3), primary_key=True)
    vagas: Mapped[int] = mapped_column("vagas", nullable=False)
