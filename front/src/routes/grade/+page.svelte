<script lang="ts">
	import { checkLogin, colegarGrade } from "$lib/api";
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import type { UIDisciplinaResumo } from "../../types/ui";
	import type { EscolhaInfoExtra, EscolhasSimples, GradeAtualExtra, RemoveDisciplinaEvent, SubmitTurmaEvent } from "../../types/data";
	import { adicionarTurmaNaGrade, loadAllInfos, removeDisciplinaNaGrade } from "$lib/grade";
	import TurmaSelecao from "../../components/grade/turma/TurmaSelecao.svelte";
	import ListaRecomendacao from "../../components/grade/recomendacao/ListaRecomendacao.svelte";
	import Grade from "../../components/grade/Grade.svelte";
	import ListaPesquisa from "../../components/grade/lateral/ListaPesquisa.svelte";
	import GrupoBotoes from "../../components/grade/botoes/GrupoBotoes.svelte";

    // variaveis principais
	let codigoGrade: string | null = null;
    let disciplinas: Map<string, UIDisciplinaResumo> = new Map<string, UIDisciplinaResumo>();
    let gradeAtual: GradeAtualExtra = {escolhas: []}
    let faltaCursar: Set<string> = new Set<string>();
    let podeCursar: Set<string> = new Set<string>();

    // dinamicas
    let escolhidas: EscolhasSimples;
    $: escolhidas = gradeAtual.escolhas.map(e => {
        return {disciplina: e.codigo, turma: e.turma}
    })
    let enableSalvar: boolean;
    $: enableSalvar = gradeAtual.escolhas.length > 0;

    // popup da selecao de turma
    let turmaCodigo: string = "";
    let isTurmaOpen: boolean = false;
    const openPopup = (e: CustomEvent<string>) => {
        turmaCodigo = e.detail;
        isTurmaOpen = true;
    };
    const closePopup = () => isTurmaOpen = false;

    // callbacks
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

	onMount(async () => {
        // carrega as informacoes das disciplinas
        let infos = await loadAllInfos();
        if (infos != null) {
            disciplinas = infos.disciplinasMap;
            faltaCursar = infos.faltaCursar;
            podeCursar = infos.podeCursar;
        }

		// verifica se a pagina possui um codigo de grade
		const params = $page.url.searchParams;
		const codigo = params.get("g");
		if (codigo !== null) {
			// salva o codigo e apaga da historia
			codigoGrade = codigo;
			window.history.replaceState({}, "", $page.url.pathname);
		}
        // carrega a grade se existir
        if (codigoGrade !== null) {
            try {
                let r = await colegarGrade(codigoGrade);
                if (r !== null) {
                    gradeAtual = r;
                } else {
                    console.log("Grade nao encontrada");
                }
            } catch (e) {
                console.log(e);
            }
        }
	});
</script>

<div id="grade-page">
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
</div>

<style>
	#grade-page {
		width: 100%;
		height: 100%;

		display: flex;
		flex-flow: column nowrap;
		justify-content: center;
		align-items: center;
		gap: 1rem;
	}

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