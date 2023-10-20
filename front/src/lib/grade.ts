import type { UIDiaDisciplina, UIHoraDisciplina, UIDisciplinaResumo, UIDisciplinaCodigo, UIEscolha } from "../types/ui";
import type { EscolhaInfoExtra, LoadDisciplinasResponse} from "../types/data";
import { coletarDisciplinasFaltaCursar, coletarDisciplinasInfo, coletarDisciplinasPodeCursar } from "./api";


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
        return e.dia === dia && e.inicio === hora;
    }

    const checkDurante = (e: EscolhaInfoExtra, dia: UIDiaDisciplina, hora: UIHoraDisciplina) => {
        return e.dia === dia && e.inicio <= hora && e.horas + e.inicio > hora;
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

export async function loadAllInfos(): Promise<LoadDisciplinasResponse | null> {
    console.log("loadAllInfos");

    let disciplinas = new Map<string, UIDisciplinaResumo>();
    let faltaCursar = new Set<string>();
    let podeCursar = new Set<string>();

    try {
        // salva as disciplinas
        let fDisciplinas = await coletarDisciplinasInfo();
        if (fDisciplinas) {
            fDisciplinas.forEach((d: UIDisciplinaResumo) => {
                disciplinas.set(d.codigo, d);           // salva a disciplina
            });
        }
        // salva os falta cursar
        let fFaltaCursar = await coletarDisciplinasFaltaCursar();
        if (fFaltaCursar) {
            faltaCursar.clear()
            fFaltaCursar.forEach((c: UIDisciplinaCodigo) => faltaCursar.add(c.codigo));
        }
        // salva os pode cursar
        let fPodeCursar = await coletarDisciplinasPodeCursar();
        if (fPodeCursar) {
            podeCursar.clear()
            fPodeCursar.forEach((c: UIDisciplinaCodigo) => podeCursar.add(c.codigo));
        }
        
        return {
            disciplinasMap: disciplinas,
            faltaCursar: faltaCursar,
            podeCursar: podeCursar,
        }

    } catch (e: any) {
        console.log(e.message);     // TODO
        return null;
    }
}