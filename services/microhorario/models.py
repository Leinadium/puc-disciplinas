from sqlalchemy import ForeignKey
from sqlalchemy import String
from sqlalchemy import DateTime
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped, mapped_column


class Base(DeclarativeBase):
    pass


class Departamento(Base):
    __tablename__ = "departamentos"

    cod_depto: Mapped[str] = mapped_column("cod_depto", String(3), primary_key=True)
    nome_depto: Mapped[str] = mapped_column("nome_depto", String(), nullable=False)

    def __repr__(self):
        return f"<Departamento {self.cod_depto}>"


class Disciplina(Base):
    __tablename__ = "disciplinas"

    cod_disciplina: Mapped[str] = mapped_column("cod_disciplina", primary_key=True)
    cod_depto = mapped_column("cod_depto",
                              ForeignKey("departamentos.cod_depto"), nullable=False)
    nome_disciplina: Mapped[str] = mapped_column("nome_disciplina", String(100), nullable=False)
    ementa: Mapped[str] = mapped_column("ementa", String(1000), nullable=True)
    creditos: Mapped[int] = mapped_column("creditos", nullable=False)

    def __repr__(self):
        return f"<Disciplina {self.cod_disciplina}>"


class Prerequisitos(Base):
    __tablename__ = "prerequisitos"

    cod_disc_orig: Mapped[str] = mapped_column("cod_disc_orig",
                                               ForeignKey("disciplinas.cod_disciplina"), primary_key=True)
    grupo_prereq: Mapped[int] = mapped_column("grupo_prereq", primary_key=True)
    cod_disc_depen = mapped_column("cod_disc_depen", String(7), primary_key=True)


class Professor(Base):
    __tablename__ = "professores"

    nome_professor: Mapped[str] = mapped_column("nome_professor", String(100), nullable=False,
                                                primary_key=True)


class Turma(Base):
    __tablename__ = "turmas"

    cod_turma: Mapped[str] = mapped_column("cod_turma", String(3), primary_key=True)
    cod_disciplina: Mapped[str] = mapped_column("cod_disciplina",
                                                ForeignKey("disciplinas.cod_disciplina"), primary_key=True)
    nome_professor: Mapped[str] = mapped_column("nome_professor",
                                                ForeignKey("professores.nome_professor"), nullable=False)
    shf: Mapped[int] = mapped_column("shf", nullable=False)

    __mapper_args__ = {
        "primary_key": [cod_turma, cod_disciplina]
    }


class Horario(Base):
    __tablename__ = "horarios"

    cod_disciplina: Mapped[str] = mapped_column("cod_disciplina", primary_key=True)
    # ForeignKey("turmas.cod_disciplina"), primary_key=True)
    cod_turma: Mapped[str] = mapped_column("cod_turma", primary_key=True)
    # ForeignKey("turmas.cod_turma"), primary_key=True)
    dia: Mapped[str] = mapped_column("dia", String(3), primary_key=True)
    hora_inicio: Mapped[int] = mapped_column("hora_inicio", primary_key=True)
    hora_fim: Mapped[int] = mapped_column("hora_fim", nullable=True)


class Alocacao(Base):
    __tablename__ = "alocacoes"

    cod_disciplina: Mapped[str] = mapped_column("cod_disciplina", primary_key=True)
    # ForeignKey("disciplinas.cod_disciplina"), primary_key=True)
    cod_turma: Mapped[str] = mapped_column("cod_turma", primary_key=True)
    # ForeignKey("turmas.cod_turma"), primary_key=True)
    destino: Mapped[str] = mapped_column("destino", String(80), primary_key=True)
    vagas: Mapped[int] = mapped_column("vagas", nullable=False)


class Modificacao(Base):
    __tablename__ = "modificacao"

    data_ementa = mapped_column("data_ementa", DateTime, primary_key=True)
    data_geral = mapped_column("data_geral", DateTime, primary_key=True)
