import type { UIDiaDisciplina, UIDisciplinaResumo, UIHoraDisciplina, UITipoAvaliacao } from "./ui"

export type EscolhaInfoExtra = {
    nome: string,
    codigo: string,
    turma: string,
    professor: string,
    horarios: DisciplinaHorariosInfo[],
    shf: number,
}

export type Modificacao = {
    dataEmenta: string,
    dataGeral: string,
    modoFallback: boolean,
}

export type DisciplinasComModificacao = {
    disciplinas: UIDisciplinaResumo[],
    modificacao: Modificacao,
}

export type LoadDisciplinasResponse = {
    disciplinasMap: Map<string, DisciplinaBasica>,
    faltaCursar: Set<string> | null,
    podeCursar: Set<string> | null,
    cursadas: Set<string> | null,
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
    notaProfessor: number,
    shf: number,
    horarios: DisciplinaHorariosInfo[],
    alocacoes: DisciplinaAlocacoesInfo[],
}

export type TabelaHorarios = Map<UIDiaDisciplina, Set<UIHoraDisciplina>>

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

export type RemoveDisciplinaEvent = {
    disciplina: string,
}

export type SalvarGradeEvent = {
    codigo: string,
}

export type PostHistorico = {
    inseridos: number,
    curriculo: string,
}

export type ItemDisciplina = {
    nome: string,
    codigo: string,
    nota: number | null,
    media: number | null,
    qtd: number,
}

export type ItemProfessor = {
    nome: string,
    nota: number | null,
    media: number | null,
    qtd: number,
}

export interface ItemGenerico {
    codigo: string,
    nome: string,
    nota: number | null,
    media: number | null,
    qtd: number,
}

export type SelectAvaliacaoEvent = {
    conteudo: ItemGenerico,
    tipo: UITipoAvaliacao,
}

export type SubmitAvaliacaoEvent = {
    avaliacao: number;
}

export type ItemsCompletos = {
    disciplinas: ItemDisciplina[],
    professores: ItemProfessor[],
}