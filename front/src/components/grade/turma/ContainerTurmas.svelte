<script lang="ts">
	import type { DisciplinaTurmasInfo } from "../../../types/data";
	import TurmaTurma from "./TurmaTurma.svelte";

    export let turmas: DisciplinaTurmasInfo[];
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