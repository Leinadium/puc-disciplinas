export type EscolhaInfoExtra = {
    nome: string,
    codigo: string,
    turma: string,
    professor: string,
    dia: DiaDisciplina,
    inicio: HoraDisciplina,
    horas: number,
}

export type LoadDisciplinasResponse = {
    disciplinasMap: Map<string, DisciplinaBasica>,
    faltaCursar: Set<string>,
    podeCursar: Set<string>
}

export type DisciplinaRecomendacao = {
    cod: string,
    val: number,
    pes: PesosRecomendacao,
}

export type PesosRecomendacao = {
    c: number,
    h: number,
    o: number,
    p: number,
    a: number,
}

export type GradeAtual = {
    escolhas: DisciplinaEscolha[], 
}

export type Escolha = {
    disciplina: string,
    turma: string,
}