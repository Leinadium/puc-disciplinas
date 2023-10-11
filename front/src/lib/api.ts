import type { ErrorApi, ListaCodigosApi, ListaDisciplinasApi } from "../types/api";
import type { DisciplinaBasica, DisciplinaCodigo } from "../types/disciplinas";

const BASE_API_URL = 'http://localhost:8080'


async function genericFetch(url: string): Promise<any> {
    let res = await fetch(url, { credentials: 'include' });
    if (res.status !== 200) {
        let body: ErrorApi | undefined = await res.json();
        if (body && body.message) {
            throw new Error(body.message);
        } else {
            throw new Error("Erro ao acessar a API");
        }
    }
    return res.json();
}

export async function coletarDisciplinasInfo(): Promise<DisciplinaBasica[] | null> {
    let body: ListaDisciplinasApi = await genericFetch(BASE_API_URL + '/pesquisa/info');
    return body.data;
}

export async function coletarDisciplinasPodeCursar(): Promise<DisciplinaCodigo[] | null> {
    let body: ListaCodigosApi = await genericFetch(BASE_API_URL + '/pesquisa/podecursar');
    return body.data;
}

export async function coletarDisciplinasFaltaCursar(): Promise<DisciplinaCodigo[] | null> {
    let body: ListaCodigosApi = await genericFetch(BASE_API_URL + '/pesquisa/faltacursar');
    return body.data;
}