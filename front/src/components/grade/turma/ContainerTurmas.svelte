<script lang="ts">
	import type { DisciplinaHorariosInfo, DisciplinaTurmaInfo, SubmitTurmaEvent, TabelaHorarios } from "../../../types/data";
    import { createEventDispatcher } from "svelte";
    import TurmaTurma from "./TurmaTurma.svelte";
	import { hasConflitoHorario } from "$lib/utils";

    export let disciplina: string = "";
    export let turmas: DisciplinaTurmaInfo[];
    export let tabelaHorariosUsados: TabelaHorarios;


    let dispatch = createEventDispatcher<{submit:SubmitTurmaEvent}>();

    const handleClick = (turma: DisciplinaTurmaInfo) => {
        dispatch("submit", {
            disciplina: disciplina,
            turma: turma
        });
    }

    const invalida = (hs: DisciplinaHorariosInfo[], v: number) => {
        return hasConflitoHorario(hs, tabelaHorariosUsados) || v == 0;
    }

</script>

{#if turmas.length == 0}
    <div id="nenhuma-turma">
        Disciplina não oferecida no momento
    </div>
{:else}
    <div id="container-turmas">
        {#each turmas as turma}
            {#each turma.alocacoes as alocacao}
                <TurmaTurma
                    codigo={turma.codigo}
                    professor={turma.professor}
                    nota={turma.notaProfessor}
                    destino={alocacao.destino}
                    vagas={alocacao.vagas}
                    horarios={turma.horarios}
                    shf={turma.shf}
                    invalida={invalida(turma.horarios, alocacao.vagas)}
                    on:click={() => {handleClick(turma)}}
                />
            {/each}
        {/each}
    </div>
{/if}

<style>
    #container-turmas {
        width: 100%;
        height: 100%;
        overflow-x: scroll;

        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-start;
        align-items: stretch;
        gap: 1%;
    }

    #nenhuma-turma {
        width: 100%;
        height: 100%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
    }
</style>