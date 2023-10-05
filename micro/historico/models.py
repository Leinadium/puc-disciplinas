from sqlalchemy.types import SMALLINT, VARCHAR
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped, mapped_column


class Base(DeclarativeBase):
    pass


class Historico(Base):
    __tablename__ = "historicos"

    cod_usuario: Mapped[str] = mapped_column("cod_usuario", primary_key=True)
    cod_disciplina: Mapped[str] = mapped_column("cod_disciplina", primary_key=True)
    semestre: Mapped[int] = mapped_column("semestre", SMALLINT, primary_key=True)
    grau: Mapped[int] = mapped_column("grau", SMALLINT)


class Usuario(Base):
    __tablename__ = "usuarios"

    cod_usuario: Mapped[str] = mapped_column("cod_usuario", primary_key=True)
    nome_usuario: Mapped[str] = mapped_column("nome_usuario", nullable=False)
    cod_curriculo: Mapped[str] = mapped_column("cod_curriculo", VARCHAR(11), nullable=True)
