import type { DisciplinaInfoApi, ErrorApi, GradeGetApi, ListaCodigosApi, ListaDisciplinasApi, ListaRecomendacoesApi } from "../types/api";
import type { DisciplinaInfo, DisciplinaRecomendacao, EscolhasSimples, GradeAtualExtra } from "../types/data";
import type { UIDisciplinaCodigo, UIDisciplinaResumo } from "../types/ui";

import { userStore } from "./stores";

const BASE_API_URL = 'http://localhost:8080';
const DISCIPLINAS_INFO_URL          = BASE_API_URL + '/pesquisa/info';
const DISCIPLINAS_PODE_CURSAR_URL   = BASE_API_URL + '/pesquisa/podecursar';
const DISCIPLINAS_FALTA_CURSAR_URL  = BASE_API_URL + '/pesquisa/faltacursar';
const RECOMENDACAO_URL              = BASE_API_URL + '/recomendacao';
const DISCIPLINA_INFO_URL           = BASE_API_URL + '/disciplina/info';
const GRADE_URL                     = BASE_API_URL + "/grade";
const LOGIN_URL                     = BASE_API_URL + '/login';
const LOGOUT_URL                    = BASE_API_URL + '/logout';


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
 * @returns {UIDisciplinaResumo[] | null} Lista de disciplinas
 */
export async function coletarDisciplinasInfo(): Promise<UIDisciplinaResumo[] | null> {
    let body: ListaDisciplinasApi = await genericFetch(DISCIPLINAS_INFO_URL);
    return body.data;
}

/**
 * Coleta os códigos de todas as disciplinas que o usuário pode cursar
 * @returns {DisciplinaCodigo[] | null} Lista de disciplinas
 */
export async function coletarDisciplinasPodeCursar(): Promise<UIDisciplinaCodigo[] | null> {
    let body: ListaCodigosApi = await genericFetch(DISCIPLINAS_PODE_CURSAR_URL);
    return body.data;
}

/**
 * Coleta os códigos de todas as disciplinas que o usuário ainda não cursou
 * @returns {DisciplinaCodigo[] | null} Lista de disciplinas
 */
export async function coletarDisciplinasFaltaCursar(): Promise<UIDisciplinaCodigo[] | null> {
    let body: ListaCodigosApi = await genericFetch(DISCIPLINAS_FALTA_CURSAR_URL);
    return body.data;
}

/**
 * Coleta as recomendações de disciplinas para o usuário
 * @param {DisciplinaEscolha[]} escolhas Lista de disciplinas escolhidas
 * @returns {DisciplinaRecomendacao[] | null} Lista de disciplinas recomendadas
 */
export async function coletarRecomendacoes(escolhas: EscolhasSimples): Promise<DisciplinaRecomendacao[] | null> {
    // nao da para usar o generic fetch, pois precisa ser um post
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
        console.log(res);

        let data = await res.json();
        if ('nome' in data) {
            return data.nome;
        }
        console.log(data);
        console.log(res.status);
        console.log(data.message);
        console.log('message' in data);

        if (res.status > 200 && 'message' in data) {
            throw new Error(data.message);
        }
        throw new Error();
    
    } catch (e: any) {
        throw new Error(e.message || "Erro ao acessar a API de login");
    }
}

export async function fazerLogout(): Promise<boolean> {
    try {
        let res = await fetch(LOGOUT_URL, {
            method: "POST",
            credentials: 'include',
        })
        return res.status === 200;
    } catch (e) {
        console.log(e);
        return false;
    }
}

/**
 * Coleta as informações completas de uma disciplina
 * @param {string} codigo Código da disciplina
 * @returns {DisciplinaInfo | null} Informações da disciplina
 */
export async function coletarDisciplinaInfo(codigo: string): Promise<DisciplinaInfo | null> {
    let body: DisciplinaInfoApi = await genericFetch(DISCIPLINA_INFO_URL + `?c=${codigo}`);
    return body.data;
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

/**
 * Armazena uma grade horária
 * @param grade Grade horaria para salvar
 */
export async function armazenarGrade(grade: GradeAtualExtra): Promise<string | null> {
    try {
        let res = await fetch(GRADE_URL, {
            method: "POST",
            headers: {
                'Content-Type': "application/json"
            },
            body: JSON.stringify(grade),
            credentials: 'include',
        });

        let data = await res.json();
        return data.id;
    } catch (e: any) {
        throw new Error("Erro ao acessar a API de login");
    }
}

/**
 * Coleta as informações de uma grade
 * @param codigo Código da grade
 */
export async function colegarGrade(codigo: string): Promise<GradeAtualExtra | null> {
    let body: GradeGetApi = await genericFetch(GRADE_URL + `?id=${codigo}`);
    return JSON.parse(body.data);
}