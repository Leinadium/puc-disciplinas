<script lang="ts">
	import { onMount } from "svelte";
	import Grade from "./Grade.svelte";
	import ListaRecomendacao from "./ListaRecomendacao.svelte";
	import ListaSelecao from "./ListaSelecao.svelte";
	import { loadAllInfos } from "$lib/grade";
	import type { UIDisciplinaResumo } from "../../types/ui";
    import type { Escolha } from "../../types/data";
	import Popup from "../common/Popup.svelte";
	import TurmaSelecao from "./turma/TurmaSelecao.svelte";

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

        console.log('disciplinas', disciplinas);
        console.log('faltaCursar', faltaCursar);
        console.log('podeCursar', podeCursar);
    })

    // popup da selecao de turma
    let turmaCodigo: string = "";
    let isTurmaOpen: boolean = false;
    const openPopup = (e: any) => {
        turmaCodigo = e.detail;
        isTurmaOpen = true;
    };
    const closePopup = () => isTurmaOpen = false;

</script>

<div id="grade-container">
    {#if isTurmaOpen}
        <TurmaSelecao 
            codigoDisciplina={turmaCodigo}
            on:close={closePopup} 
        />
    {/if}


    <ListaRecomendacao
        {disciplinas}
        {escolhidas}
        {faltaCursar}
        {podeCursar}
        on:popup={openPopup}
    />
    <Grade />
    <ListaSelecao 
        {disciplinas}
        {faltaCursar}
        {podeCursar}
        on:popup={openPopup}
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