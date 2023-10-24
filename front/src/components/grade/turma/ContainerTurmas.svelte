<script lang="ts">
	import type { DisciplinaTurmaInfo, SubmitTurmaEvent } from "../../../types/data";
    import { createEventDispatcher } from "svelte";
    import TurmaTurma from "./TurmaTurma.svelte";

    export let disciplina: string = "";
    export let turmas: DisciplinaTurmaInfo[];

    let dispatch = createEventDispatcher<{submit:SubmitTurmaEvent}>();

    const handleClick = (turma: DisciplinaTurmaInfo) => {
        dispatch("submit", {
            disciplina: disciplina,
            turma: turma
        });
    }
</script>

<div id="container-turmas">
    {#each turmas as turma}
        {#each turma.alocacoes as alocacao}
            <TurmaTurma
                codigo={turma.codigo}
                professor={turma.professor}
                destino={alocacao.destino}
                vagas={alocacao.vagas}
                horarios={turma.horarios}
                shf={turma.shf}
                on:click={() => {handleClick(turma)}}
            />
        {/each}
    {/each}
</div>

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
</style>