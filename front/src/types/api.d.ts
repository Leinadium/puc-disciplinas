import type { DisciplinaBasica } from "./disciplinas"

export type ErrorApi ={
    message: string
}

export type ListaDisciplinasApi = {
    data: DisciplinaBasica[]
}