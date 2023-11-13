import type { DisciplinaHorariosInfo, GradeAtualExtra, TabelaHorarios } from "../types/data";
import type { UIDiaDisciplina, UIDisciplinaCodigo, UIHoraDisciplina } from "../types/ui";

/**
 * Função simples de pluralização
 * @param n numero de elementos
 * @param text texto a ser pluralizado
 * @returns 
 */
export function pluralize(n: number, text: string): string {
    if (n == 1) return text;
    if (text.endsWith("s")) return text + "es";
    if (text.endsWith("ão")) return text.substring(0, text.length - 2) + "ões";
    return text + "s";
}

/**
 * Extrai todos os horarios das disciplinas escolhidas
 * @param grade grade horaria atual
 */
export function extractHorarios(grade: GradeAtualExtra): DisciplinaHorariosInfo[] {
    return grade.escolhas.map(e => e.horarios).flat()
}

/**
 * Cria uma tabela com os horarios ocupados da grade horaria atual
 * @param horarios os horarios da grade atual
 * @returns Um set contendo os horarios ocupados da grade atual
 */
export function criaTabelaHorarios(horarios: DisciplinaHorariosInfo[]): TabelaHorarios {
    const ret: TabelaHorarios = new Map();
    horarios.forEach(h => {
        // se for uma disciplina sem horario, ignora
        if (h.dia == "xxx") return;
        // cria um set do dia se ja nao existir
        if (!ret.has(h.dia)) ret.set(h.dia, new Set());
        // adiciona todos as hora entre os horarios do dia
        for (let i = h.inicio; i < h.fim; i++) ret.get(h.dia)!.add(i);
    });
    return ret;
}

/**
 * Descobre se os horarios de uma turma estão em conflito com os da grade horaria
 * @param meu os horarios da turma
 * @param horarios a tabela de horarios da grade
 * @returns true se houver conflito, false caso contrario
 */
export function hasConflitoHorario(
    meu: DisciplinaHorariosInfo[], 
    horarios: TabelaHorarios
): boolean {
    // verifica se algum dos horarios da turma esta em conflito com os da grade
    for (let i = 0; i < meu.length; i++) {
        const h = meu[i];
        // se for uma disciplina sem horario, ignora
        if (h.dia == "xxx") continue;
        // se o dia nao existir na grade, nao tem como ter conflito
        if (!horarios.has(h.dia)) continue;
        // se algum dos horarios da turma estiver no set de horarios da grade, tem conflito
        for (let j = h.inicio; j < h.fim; j++) if (horarios.get(h.dia)!.has(j)) return true;
    }
    return false;
}

export const codigoCores = {
    c: "C",
    h: "H",
    o: "A",
    p: "P",
    a: "N",
}

export const pesosCores = {
    c: "red",
    h: "blue",
    o: "orange",
    p: "purple",
    a: "pink"
}

export const pesosDescricao = {
    c: "Conteúdo relevante",
    h: "Muitos horários livres",
    o: "Boas avaliações",
    p: "Bons professores",
    a: "Notas altas"
}

/**
 * Calcula a diferença entre a data passada e a hora atual
 * Retorna uma string no formato "x dias" se a diferença for maior que 1 dia,
 * ou "x horas" se a diferença for menor que 1 dia
*/
export function calcularDiferenca(h: Date) {
    const agora = new Date();
    const diferenca = - Math.floor((h.getTime() - agora.getTime()) / 1000 / 60 / 60);
    if (diferenca > 24) return `${Math.floor(diferenca / 24)} dias`;
    return `${diferenca} horas`;
}

/**
 * Converte um dia da semana para um numero
 * @param cod codigo da disciplina
 * @returns um link para a ementa da disciplina
 */
export function geraLinkEmenta(cod: string): string {
    return 'https://www.puc-rio.br/ferramentas/ementas/ementa.aspx?cd=' + cod;
}