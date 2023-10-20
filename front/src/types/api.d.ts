import type { DisciplinaBasica, DisciplinaCodigo, DisciplinaRecomendacao } from "./disciplinas"

export type ErrorApi ={
    message: string
}

export type ListaDisciplinasApi = {
    data: DisciplinaBasica[]
}

export type ListaCodigosApi = {
    data: DisciplinaCodigo[]
}

export type ListaRecomendacoesApi = {
    data: DisciplinaRecomendacao[]
}