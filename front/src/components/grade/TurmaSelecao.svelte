<script lang="ts">
	import { onMount, createEventDispatcher } from "svelte";
	import Popup from "../common/Popup.svelte";

    export let codigoDisciplina: string;

    
    onMount(() => {});
</script>

<Popup on:close>
    <div id="turma-selecao">
        <span>codigo</span>
        <span>nome</span>
        <span>ementa</span>
        <span>creditos</span>
        <span>pre-requisitos</span>
        <span>avaliacao</span>
        <span>grau médio</span>
        <span>quantidade de alunos</span>

        <div id="turmas">
            <div class="turma">
                <span>codigo</span>         <!-- turmas    -->
                <span>professor</span>      <!-- turmas    -->
                <span>shf</span>            <!-- turmas    -->
                <span>horario</span>        <!-- horarios  -->
                <span>destino</span>        <!-- alocacoes -->
                <span>qtdVagas</span>       <!-- alocacoes -->
            </div>
        </div>


    </div>
</Popup>