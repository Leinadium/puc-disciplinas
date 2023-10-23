import type { UIDiaDisciplina, UIHoraDisciplina } from "./ui"

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

export type DisciplinaInfo = {
    codigo: string,
    nome: string,
    creditos: number,
    ementa: string,
    preRequisitos: DisciplinaPreRequisitosInfo[],
    grauMedio: number,
    qtdAlunos: number,
    avaliacaoMedia: number,
    qtdAvaliacoes: number,
    turmas: DisciplinaTurmasInfo[],
}

export type DisciplinaPreRequisitosInfo = {
    grupoId: number,
    preReqs: string[],
}

export type DisciplinaTurmasInfo = {
    codigo: string,
    professor: string,
    shf: number,
    horarios: DisciplinaHorariosInfo[],
    alocacoes: DisciplinaAlocacoesInfo[],
}

export type DisciplinaHorariosInfo = {
    dia: UIDiaDisciplina,
    inicio: UIHoraDisciplina,
    fim: UIHoraDisciplina,
}

export type DisciplinaAlocacoesInfo = {
    destino: string,
    vagas: number,
}