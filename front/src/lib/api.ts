import type { ErrorApi, ListaCodigosApi, ListaDisciplinasApi } from "../types/api";
import type { DisciplinaBasica, DisciplinaCodigo } from "../types/disciplinas";

import { userStore } from "./stores";

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

export async function fazerLogin(usuario: string, senha: string): Promise<string> {
    // nao da pra usar o genericFetch, pois precisa configurar outras coisas
    try {
        let res = await fetch(BASE_API_URL + '/login', {
            method: "POST",
            headers: {
                'Content-Type': "application/x-www-form-urlencoded"
            },
            body: `usuario=${usuario}&senha=${senha}`,
            credentials: 'include',
        });

        let data = await res.json();
        if ('nome' in data) {
            return data.nome;
        }

        if (res.status > 200 && 'message' in data) {
            throw new Error(data.message);
        }

        throw new Error("Erro ao acessar a API de login");
    } catch (e: any) {
        throw new Error("Erro ao acessar a API de login");
    }
}

/**
 * Verifica se o usuário está logado.
 * Se estiver logado, retorna true e configura a store do usuário.
 * Se não estiver logado, retorna false
 */
export async function checkLogin(): Promise<boolean> {
    try {
        let res = await fetch(BASE_API_URL + '/login', {
            credentials: 'include',
        });
        let data = await res.json();
        // logado
        if ('nome' in data) {
            userStore.set(data.nome);
            return true;
        }
        // nao logado
        if (res.status === 401) {
            userStore.set(null);
            return false;
        }
        // erro
        if (res.status > 200 && 'message' in data) {
            throw new Error(data.message);
        }
        return false;
    } catch (e: any) {
        throw new Error("Erro ao acessar a API de login");
    }
}