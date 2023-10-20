export type UIDisciplinaResumo = {
    nome: string,
    codigo: string,
    creditos: number,
    qtdVagas: number,
    qtdTurmas: number,
};

export type UIDisciplinaCodigo = {
    codigo: string,
};

export type UIEscolha = {
    nome: string,
    codigo: string,
    turma: string,
    professor: string,
};

export type UIDiaDisciplina = "SEG" | "TER" | "QUA" | "QUI" | "SEX" | "SAB" | "SHF";
export type UIHoraDisciplina = 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20;
