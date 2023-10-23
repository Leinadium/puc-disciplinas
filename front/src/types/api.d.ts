import type { DisciplinaInfo, DisciplinaRecomendacao } from "./data"
import type { UIDisciplinaCodigo, UIDisciplinaResumo } from "./ui"

export type ErrorApi = {
    message: string
}

export type ListaDisciplinasApi = {
    data: UIDisciplinaResumo[]
}

export type ListaCodigosApi = {
    data: UIDisciplinaCodigo[]
}

export type ListaRecomendacoesApi = {
    data: DisciplinaRecomendacao[]
}

export type DisciplinaInfoApi = {
    data: DisciplinaInfo
}