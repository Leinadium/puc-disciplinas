import type { ErrorApi, ListaDisciplinasApi } from "../types/api";
import type { DisciplinaBasica } from "../types/disciplinas";

const API_URL = 'http://localhost:8080/disciplinas/pesquisa?query='

export async function pesquisarDisciplinas(texto: string): Promise<DisciplinaBasica[]> {
    let res = await fetch(API_URL + texto);
    if (res.status != 200) {
        let body: ErrorApi | undefined = await res.json();
        if (body && body.message) {
            throw new Error(body.message);
        } else {
            throw new Error("Erro ao acessar a API");
        }
    }

    let body: ListaDisciplinasApi = await res.json();
    let disciplinas: DisciplinaBasica[] = body.data;
    return disciplinas
}
