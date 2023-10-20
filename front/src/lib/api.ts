import type { ErrorApi, ListaCodigosApi, ListaDisciplinasApi, ListaRecomendacoesApi } from "../types/api";
import type { DisciplinaBasica, DisciplinaCodigo, DisciplinaRecomendacao } from "../types/disciplinas";
import type { DisciplinaEscolha } from "../types/grade";

import { userStore } from "./stores";

const BASE_API_URL = 'http://localhost:8080'
const DISCIPLINA_INFO_URL = BASE_API_URL + '/pesquisa/info';
const DISCIPLINA_PODE_CURSAR_URL = BASE_API_URL + '/pesquisa/podecursar';
const DISCIPLINA_FALTA_CURSAR_URL = BASE_API_URL + '/pesquisa/faltacursar';
const RECOMENDACAO_URL = BASE_API_URL + '/grade/recomendacao';
const LOGIN_URL = BASE_API_URL + '/login';


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

/**
 * Coleta as informações básicas de todas as disciplinas 
 * @returns {DisciplinaBasica[] | null} Lista de disciplinas
 */
export async function coletarDisciplinasInfo(): Promise<DisciplinaBasica[] | null> {
    let body: ListaDisciplinasApi = await genericFetch(DISCIPLINA_INFO_URL);
    return body.data;
}

/**
 * Coleta os códigos de todas as disciplinas que o usuário pode cursar
 * @returns {DisciplinaCodigo[] | null} Lista de disciplinas
 */
export async function coletarDisciplinasPodeCursar(): Promise<DisciplinaCodigo[] | null> {
    let body: ListaCodigosApi = await genericFetch(DISCIPLINA_PODE_CURSAR_URL);
    return body.data;
}

/**
 * Coleta os códigos de todas as disciplinas que o usuário ainda não cursou
 * @returns {DisciplinaCodigo[] | null} Lista de disciplinas
 */
export async function coletarDisciplinasFaltaCursar(): Promise<DisciplinaCodigo[] | null> {
    let body: ListaCodigosApi = await genericFetch(DISCIPLINA_FALTA_CURSAR_URL);
    return body.data;
}

/**
 * Coleta as recomendações de disciplinas para o usuário
 * @param {DisciplinaEscolha[]} escolhas Lista de disciplinas escolhidas
 * @returns {DisciplinaRecomendacao[] | null} Lista de disciplinas recomendadas
 */
export async function coletarRecomendacoes(escolhas: DisciplinaEscolha[]): Promise<DisciplinaRecomendacao[] | null> {
    // nao da para usar o generic fetch, pois precisa ser um post
    console.log("escolhas: ");
    console.log(escolhas);
    try {
        let res = await fetch(RECOMENDACAO_URL, {
            method: "POST",
            headers: {
                'Content-Type': "application/json"
            },
            body: JSON.stringify({escolhas: escolhas}),
            credentials: 'include',
        });

        let data = await res.json();
        return data.data;
    } catch (e: any) {
        throw new Error("Erro ao acessar a API de recomendação");
    }

}

export async function fazerLogin(usuario: string, senha: string): Promise<string> {
    // nao da pra usar o genericFetch, pois precisa configurar outras coisas
    try {
        let res = await fetch(LOGIN_URL, {
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
