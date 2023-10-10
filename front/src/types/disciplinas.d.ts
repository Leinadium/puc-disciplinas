export type DisciplinaBasica = {
    nome: string,
    codigo: string,
    creditos: number,
    qtdVagas: number,
    qtdTurmas: number,
}

export type Escolha = {
    nome: string,
    codigo: string,
    turma: string,
    professor: string,
}

export type DiaDisciplina = "SEG" | "TER" | "QUA" | "QUI" | "SEX" | "SAB" | "SHF";
export type HoraDisciplina = 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20;

export type EscolhaInfoExtra = {
    nome: string,
    codigo: string,
    turma: string,
    professor: string,
    dia: DiaDisciplina,
    inicio: HoraDisciplina,
    horas: number,
}

