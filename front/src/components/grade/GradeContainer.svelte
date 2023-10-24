<script lang="ts">
	import { onMount } from "svelte";
	import Grade from "./Grade.svelte";
	import ListaRecomendacao from "./recomendacao/ListaRecomendacao.svelte";
	import ListaPesquisa from "./lateral/ListaPesquisa.svelte";
	import { adicionarTurmaNaGrade, loadAllInfos, removeDisciplinaNaGrade } from "$lib/grade";
	import type { UIDisciplinaResumo } from "../../types/ui";
    import type { EscolhaInfoExtra, EscolhasSimples, GradeAtualExtra, RemoveDisciplinaEvent, SubmitTurmaEvent } from "../../types/data";
	import TurmaSelecao from "./turma/TurmaSelecao.svelte";
	import GrupoBotoes from "./botoes/GrupoBotoes.svelte";
	import { armazenarGrade } from "$lib/api";

    let disciplinas: Map<string, UIDisciplinaResumo> = new Map<string, UIDisciplinaResumo>();
    let gradeAtual: GradeAtualExtra = {escolhas: []}
    let faltaCursar: Set<string> = new Set<string>();
    let podeCursar: Set<string> = new Set<string>();

    let escolhidas: EscolhasSimples;
    $: escolhidas = gradeAtual.escolhas.map(e => {
        return {disciplina: e.codigo, turma: e.turma}
    })

    let enableSalvar: boolean;
    $: enableSalvar = gradeAtual.escolhas.length > 0;

    onMount(async () => {
        let infos = await loadAllInfos();
        if (infos != null) {
            disciplinas = infos.disciplinasMap;
            faltaCursar = infos.faltaCursar;
            podeCursar = infos.podeCursar;
        }
    })

    // popup da selecao de turma
    let turmaCodigo: string = "";
    let isTurmaOpen: boolean = false;
    const openPopup = (e: CustomEvent<string>) => {
        turmaCodigo = e.detail;
        isTurmaOpen = true;
    };
    const closePopup = () => isTurmaOpen = false;

    function submitTurma(e: CustomEvent<SubmitTurmaEvent>) {
        const escolha = e.detail;
        const escolhaInfoExtra: EscolhaInfoExtra = {
            codigo: escolha.disciplina,
            nome: disciplinas.get(escolha.disciplina)?.nome ?? "",
            turma: escolha.turma.codigo,
            professor: escolha.turma.professor,
            horarios: escolha.turma.horarios,
            shf: escolha.turma.shf
        }

        gradeAtual = adicionarTurmaNaGrade(escolhaInfoExtra, gradeAtual);
        
        closePopup();
    }

    function removeDisciplina(e: CustomEvent<RemoveDisciplinaEvent>) {
        gradeAtual = removeDisciplinaNaGrade(e.detail.disciplina, gradeAtual);
    }

    function handleLimpar() {
        gradeAtual.escolhas = [];
    }

</script>

<div id="grade-container">
    {#if isTurmaOpen}
        <TurmaSelecao 
            codigoDisciplina={turmaCodigo}
            on:close={closePopup}
            on:submit={submitTurma}
        />
    {/if}

    <ListaRecomendacao
        {disciplinas}
        {escolhidas}
        {faltaCursar}
        {podeCursar}
        on:popup={openPopup}
    />

    <Grade
        info={gradeAtual.escolhas}
        on:popup={openPopup}
        on:remove={removeDisciplina}
    />

    <ListaPesquisa 
        {disciplinas}
        {faltaCursar}
        {podeCursar}
        on:popup={openPopup}
    />

    <GrupoBotoes
        {enableSalvar}
        {gradeAtual}
        on:limpar={handleLimpar}
    />
</div>


<style>
    #grade-container {
        box-sizing: border-box;
        max-width: min(80vw, 1000px);
        height: 700px;
        width: 100%;

        display: inline-grid;
        grid-template-columns: 84% 15%;
        grid-template-rows: 19% 80%;
        place-content: stretch;
        gap: 1%;

        border: 3px solid blue;
    }
</style>