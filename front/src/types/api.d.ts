import type { DisciplinaBasica, DisciplinaCodigo } from "./disciplinas"

export type ErrorApi ={
    message: string
}

export type ListaDisciplinasApi = {
    data: DisciplinaBasica[]
}

export type ListaCodigosApi = {
    data: DisciplinaCodigo[]
}