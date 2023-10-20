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
    <Grade />
    <ListaSelecao 
        {disciplinas}
        {faltaCursar}
        {podeCursar}
    />
    <ListaRecomendacao
        {disciplinas}
        {escolhidas}
        {faltaCursar}
        {podeCursar}
    />
</div>


<style>
    #grade-container {
        width: min(80vw, 1000px);
        height: 500px;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: stretch;
        gap: 1rem;

        overflow-y: scroll;
    }
</style>