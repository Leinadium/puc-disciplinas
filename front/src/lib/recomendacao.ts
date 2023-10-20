import type { DisciplinaRecomendacao } from "../types/data";

export type ModoRecomendacao = "todas" | "eletivas" | "faltacursar";

function filtroRecomendacoes(
    r: DisciplinaRecomendacao,
    podeCursar: Set<string> | null | undefined,
    faltaCursar: Set<string> | null | undefined,
    modo: ModoRecomendacao = "todas" 
) {
    const pode = podeCursar ? podeCursar.has(r.cod) : true;
    const falta = faltaCursar ? faltaCursar.has(r.cod) : true;
    const naoFalta = faltaCursar ? !faltaCursar.has(r.cod) : true;

    // se o usuario nao selecionou nenhum filtro especifico
    // entao filtra pelo o que ele pode cursar, se for definido
    if (modo == "todas") {
        return pode;
    }
    // se o usuario selecionou o filtro de eletivas
    // entao filtra pelo o que ele pode cursar mas nao falta cursar
    else if (modo == "eletivas") {
        return pode && naoFalta;
    }
    // se o usuario selecionou o filtro de falta cursar
    // entao filtra pelo o que ele pode cursar e falta cursar
    else {
        return pode && falta;
    }
}

/**
 * Filtra as recomendações recebidas, ranqueando de acordo com o valor.
 * Retorna os primeiros n valores.
 * 
 * Caso faltaCursar seja fornecido, então as recomendações 
 * 
 * @param recomendacoes Lista contendo os dados das recomendações 
 * @param n Quantidade de recomendações a serem retornadas
 * @param podeCursar Códigos das disciplinas que o usuário pode cursar. Se for nulo, então todas as disciplinas são consideradas
 * @param faltaCursar Código das disciplinas que o usuário falta cursar. Se for nulo, então todas as disciplinas são consideradas
 */
export function filtrarRecomendacoes(
    recomendacoes: DisciplinaRecomendacao[],
    n: number = 20,
    podeCursar: Set<string> | null | undefined,
    faltaCursar: Set<string> | null | undefined,
    modo: ModoRecomendacao = "todas"
) {
    // filtra as recomendações
    let res = recomendacoes.filter(r => filtroRecomendacoes(r, podeCursar, faltaCursar, modo));
    // ordena as recomendações
    res.sort((a, b) => b.val - a.val);
    // retorna as primeiras n recomendações
    return res.slice(0, n);
}