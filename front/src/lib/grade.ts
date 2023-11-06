import type { UIDiaDisciplina, UIHoraDisciplina, UIDisciplinaResumo, UIDisciplinaCodigo, UIEscolha } from "../types/ui";
import type { EscolhaInfoExtra, EscolhaSimples, GradeAtualExtra, LoadDisciplinasResponse} from "../types/data";
import { coletarDisciplinasCursadas, coletarDisciplinasFaltaCursar, coletarDisciplinasInfo, coletarDisciplinasPodeCursar } from "./api";


export function formatHora(h: UIHoraDisciplina): string {
    // const f = (h: number) => h.toString().padStart(2, '0') + 'h'
    // return f(h) + '-' + f(h + 1);
    return h.toString().padStart(2, '0') + 'h';
}

export function removeExtraFromEscolha(e: EscolhaInfoExtra): UIEscolha {
    return {
        nome: e.nome,
        codigo: e.codigo,
        professor: e.professor,
        turma: e.turma,
    };
}

export function getDuracao(e: EscolhaInfoExtra, dia: UIDiaDisciplina): number {
    for (let h of e.horarios) {
        if (h.dia === dia)
            return h.fim - h.inicio;
    }
    return 0;
}

export function generateGrade(
    dias: UIDiaDisciplina[],
    horas: UIHoraDisciplina[],
    escolhas: EscolhaInfoExtra[]) {
    // a resposta deve ser um objeto de objetos de escolhas
    // o primeiro nivel representa as horas
    // o segundo nivel representa as datas
    // cada valor pode ser null ou uma escolha
    // mas a quantidade de dias no segundo nivel nao é necessariamente igual
    // pois uma disciplina pode ocupar várias horas no mesmo dia

    let res: any = {};

    const checkInicio = (e: EscolhaInfoExtra, dia: UIDiaDisciplina, hora: UIHoraDisciplina) => {
        for (let h of e.horarios) {
            if (h.dia === dia && h.inicio === hora)
                return true;
        }
        return false;
    }

    const checkDurante = (e: EscolhaInfoExtra, dia: UIDiaDisciplina, hora: UIHoraDisciplina) => {
        for (let h of e.horarios) {
            if (h.dia === dia && h.inicio <= hora && h.fim > hora)
                return true;
        }
        return false;
    }

    for (let hora of horas) {
        let linha: any = {};
        for (let dia of dias) {
            // primero, verifica se precisa adicionar uma disciplina que começa naquela hora
            // se sim, adiciona na lista e vai pro proximo
            let escolha = escolhas.find(e => checkInicio(e, dia, hora));
            if (escolha !== undefined) {
                linha[dia] = escolha;
                continue
            }
            // se nao, verifica se tem alguma disciplina que ja esta ocupando aquele horario
            // se nao tiver, adiciona um nulo. Se tiver, não faz nada
            escolha = escolhas.find(e => checkDurante(e, dia, hora));
            if (escolha === undefined) {
                linha[dia] = null;
            }
        }
        res[hora] = linha;
    }
    return res;
}

async function getDisciplinasInfo(): Promise<Map<string, UIDisciplinaResumo>> {
    let disciplinas = new Map<string, UIDisciplinaResumo>();
    let r = await coletarDisciplinasInfo();
    // TODO: armazenar as modificacoes
    let fDisciplinas = r?.disciplinas;
    if (fDisciplinas) 
        fDisciplinas.forEach((d: UIDisciplinaResumo) => disciplinas.set(d.codigo, d));
    return disciplinas
}

export async function getDisciplinasFaltaCursar(): Promise<Set<string> | null> {
    let faltaCursar = new Set<string>();
    let fFaltaCursar = await coletarDisciplinasFaltaCursar();
    if (fFaltaCursar) 
        fFaltaCursar.forEach((c: UIDisciplinaCodigo) => faltaCursar.add(c.codigo));
    else
        return null
    return faltaCursar;
}

export async function getDisciplinasPodeCursar(): Promise<Set<string> | null> {
    let podeCursar = new Set<string>();
    let fPodeCursar = await coletarDisciplinasPodeCursar();
    if (fPodeCursar) 
        fPodeCursar.forEach((c: UIDisciplinaCodigo) => podeCursar.add(c.codigo));
    else
        return null;
    return podeCursar;
}

export async function getDisciplinasCursadas(): Promise<Set<string> | null> {
    let cursadas = new Set<string>();
    let fCursadas = await coletarDisciplinasCursadas();
    if (fCursadas) 
        fCursadas.forEach((c: UIDisciplinaCodigo) => cursadas.add(c.codigo));
    else
        return null;
    return cursadas;
}

export async function loadAllInfos(): Promise<LoadDisciplinasResponse | null> {
    try {
        return {
            disciplinasMap: await getDisciplinasInfo(),
            faltaCursar: await getDisciplinasFaltaCursar(),
            podeCursar: await getDisciplinasPodeCursar(),
            cursadas: await getDisciplinasCursadas()
        }
    } catch (e: any) {
        console.log(e.message);     // TODO
        return null;
    }
}

/**
 * Adiciona uma escolha na grade, retornando a própria grade
 * @param escolha nova escolha
 * @param grade grade atual
 * @returns a propria grade
 */
export function adicionarTurmaNaGrade(escolha: EscolhaInfoExtra, grade: GradeAtualExtra): GradeAtualExtra {
    grade.escolhas.push(escolha);
    return grade;
}

/**
 * Remove uma disciplina da grade, retornando a própria grade
 * @param escolha disciplina a ser removida
 * @param grade a propria grade
 */
export function removeDisciplinaNaGrade(disciplina: string, grade: GradeAtualExtra): GradeAtualExtra {
    grade.escolhas = grade.escolhas.filter(e => e.codigo !== disciplina);
    return grade;
}