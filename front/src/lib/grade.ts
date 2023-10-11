import type { DiaDisciplina, Escolha, EscolhaInfoExtra, HoraDisciplina } from "../types/disciplinas";


function formatHora(h: HoraDisciplina): string {
    // const f = (h: number) => h.toString().padStart(2, '0') + 'h'
    // return f(h) + '-' + f(h + 1);
    return h.toString().padStart(2, '0') + 'h';
}

function removeExtraFromEscolha(e: EscolhaInfoExtra): Escolha {
    return {
        nome: e.nome,
        codigo: e.codigo,
        professor: e.professor,
        turma: e.turma,
    };
}

function generateGrade(
    dias: DiaDisciplina[],
    horas: HoraDisciplina[],
    escolhas: EscolhaInfoExtra[]) {
    // a resposta deve ser um objeto de objetos de escolhas
    // o primeiro nivel representa as horas
    // o segundo nivel representa as datas
    // cada valor pode ser null ou uma escolha
    // mas a quantidade de dias no segundo nivel nao é necessariamente igual
    // pois uma disciplina pode ocupar várias horas no mesmo dia

    let res: any = {};

    const checkInicio = (e: EscolhaInfoExtra, dia: DiaDisciplina, hora: HoraDisciplina) => {
        return e.dia === dia && e.inicio === hora;
    }

    const checkDurante = (e: EscolhaInfoExtra, dia: DiaDisciplina, hora: HoraDisciplina) => {
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
export { generateGrade, removeExtraFromEscolha, formatHora };