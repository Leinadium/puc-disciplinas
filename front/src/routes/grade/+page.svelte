<script lang="ts">
	import { colegarGrade, coletarModificacao } from "$lib/api";
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import type { UIDisciplinaResumo } from "../../types/ui";
	import type { EscolhaInfoExtra, EscolhasSimples, GradeAtualExtra, Modificacao, RemoveDisciplinaEvent, SubmitTurmaEvent, TabelaHorarios } from "../../types/data";
	import { adicionarTurmaNaGrade, getDisciplinasCursadas, getDisciplinasFaltaCursar, getDisciplinasPodeCursar, loadAllInfos, removeDisciplinaNaGrade } from "$lib/grade";
	import TurmaSelecao from "../../components/grade/turma/TurmaSelecao.svelte";
	import ListaRecomendacao from "../../components/grade/recomendacao/ListaRecomendacao.svelte";
	import Grade from "../../components/grade/Grade.svelte";
	import ListaPesquisa from "../../components/grade/lateral/ListaPesquisa.svelte";
	import GrupoBotoes from "../../components/grade/botoes/GrupoBotoes.svelte";
	import { criaTabelaHorarios, extractHorarios } from "$lib/utils";
	import { userEvent, userStore } from "$lib/stores";
	import TextosDescricao from "../../components/grade/TextosDescricao.svelte";

    // variaveis principais
	let codigoGrade: string | null = null;
    let disciplinas: Map<string, UIDisciplinaResumo> = new Map<string, UIDisciplinaResumo>();
    let gradeAtual: GradeAtualExtra = {escolhas: []}
    let faltaCursar: Set<string> | null = null;
    let podeCursar: Set<string> | null = null;
    let cursadas: Set<string> | null = null;
    let modificacao: Modificacao | null = null;

    // dinamicas
    // escolhidas: grade simplificada.
    let escolhidas: EscolhasSimples;
    $: escolhidas = gradeAtual.escolhas.map(e => {
        return {disciplina: e.codigo, turma: e.turma}
    })
    // enableSalvar: flag para liberar o botão de salvar
    let enableSalvar: boolean;
    $: enableSalvar = gradeAtual.escolhas.length > 0;
    // tabelaHorariosUsados: uma tabela com os horarios usados na grade
    let tabelaHorariosUsados: TabelaHorarios;
    $: tabelaHorariosUsados = criaTabelaHorarios(extractHorarios(gradeAtual));

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

    // roda na primeira vez que carrega a pagina:
    // pega todas as disciplinas,falta cursar e pode cursar
	onMount(async () => {
        // carrega as informacoes das disciplinas
        let infos = await loadAllInfos();
        if (infos != null) {
            disciplinas = infos.disciplinasMap;
            faltaCursar = infos.faltaCursar;
            podeCursar = infos.podeCursar;
            cursadas = infos.cursadas;
        }

        // carrega a modificacao
        try {
            modificacao = await coletarModificacao();
        } catch (e) {
            console.log("Erro ao carregar modificacao");
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
        // callback caso o usuario mude
        userEvent.subscribe(async () => {
            // pega as informacoes atualizadas,
            // mesmo que o usuario nao esteja mais logado
            console.log("userStore subscribe");
            faltaCursar = await getDisciplinasFaltaCursar();
            podeCursar = await getDisciplinasPodeCursar();
            cursadas = await getDisciplinasCursadas();
        })
	});
</script>

<div id="grade-page">

    <TextosDescricao {modificacao}/>

    <div id="grade-container">
        {#if isTurmaOpen}
            <TurmaSelecao 
                codigoDisciplina={turmaCodigo}
                tabelaHorariosUsados={tabelaHorariosUsados}
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
            {cursadas}
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
		min-height: 100%;

		display: flex;
		flex-flow: column nowrap;
		justify-content: center;
		align-items: center;
		gap: 1rem;
	}

    #grade-container {
        box-sizing: border-box;
        max-width: min(80vw, 1000px);
        height: 1200px;
        width: 100%;

        display: inline-grid;
        grid-template-columns: 84% 15%;
        grid-template-rows: 12% 87%;
        place-content: stretch;
        gap: 1%;

        background: var(--color-mono-2);
        padding: 1%;
        border-radius: var(--border-radius);

    }
</style>