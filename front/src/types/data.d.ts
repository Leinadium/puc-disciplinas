import type { UIDiaDisciplina, UIHoraDisciplina } from "./ui"

export type EscolhaInfoExtra = {
    nome: string,
    codigo: string,
    turma: string,
    professor: string,
    horarios: DisciplinaHorariosInfo[],
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

export type GradeAtualExtra = {
    escolhas: EscolhaInfoExtra[],
}

export type GradeAtual = {
    escolhas: EscolhaSimples[], 
}

export type EscolhaSimples = {
    disciplina: string,
    turma: string,
}

export type EscolhasSimples = EscolhaSimples[]

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
    turmas: DisciplinaTurmaInfo[],
}

export type DisciplinaPreRequisitosInfo = {
    grupoId: number,
    preReqs: string[],
}

export type DisciplinaTurmaInfo = {
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

export type SubmitTurmaEvent = {
    disciplina: string,
    turma: DisciplinaTurmaInfo
}