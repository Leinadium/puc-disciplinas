<script lang="ts">
	import { onMount } from "svelte";
	import Grade from "./Grade.svelte";
	import ListaRecomendacao from "./ListaRecomendacao.svelte";
	import ListaSelecao from "./ListaSelecao.svelte";
	import { loadAllInfos } from "$lib/grade";
	import type { UIDisciplinaResumo } from "../../types/ui";
    import type { Escolha } from "../../types/data";

    let disciplinas: Map<string, UIDisciplinaResumo> = new Map<string, UIDisciplinaResumo>();
    let escolhidas: Escolha[] = [];
    let faltaCursar: Set<string> = new Set<string>();
    let podeCursar: Set<string> = new Set<string>();

    onMount(async () => {
        let infos = await loadAllInfos();
        if (infos != null) {
            disciplinas = infos.disciplinasMap;
            faltaCursar = infos.faltaCursar;
            podeCursar = infos.podeCursar;
        }
    })
</script>

<div id="grade-container">
    <ListaRecomendacao
        {disciplinas}
        {escolhidas}
        {faltaCursar}
        {podeCursar}
    />
    <Grade />
    <ListaSelecao 
        {disciplinas}
        {faltaCursar}
        {podeCursar}
    />
</div>


<style>
    #grade-container {
        box-sizing: border-box;
        max-width: min(80vw, 1000px);
        height: 600px;
        width: 100%;

        display: inline-grid;
        grid-template-columns: 84% 15%;
        grid-template-rows: 19% 80%;
        place-content: stretch;
        gap: 1%;

        border: 3px solid blue;
    }
</style>