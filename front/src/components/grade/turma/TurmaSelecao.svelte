<script lang="ts">
	import { onMount, createEventDispatcher } from "svelte";
	import Popup from "../../common/Popup.svelte";
	import type { DisciplinaInfo, TabelaHorarios } from "../../../types/data";
	import { coletarDisciplinaInfo } from "$lib/api";
	import ContainerTurmas from "./ContainerTurmas.svelte";
	import TurmaTextos from "./TurmaTextos.svelte";
	import TurmaInfo from "./TurmaInfo.svelte";

    export let codigoDisciplina: string;
    export let tabelaHorariosUsados: TabelaHorarios;

    let infoApi: Promise<DisciplinaInfo | null>;
    let dispatch = createEventDispatcher();

    function close() {
        dispatch("close");
    }

    onMount(() => {
        infoApi = coletarDisciplinaInfo(codigoDisciplina);
    });
</script>
<Popup on:close>
    {#await infoApi}
        <div id="loading">Carregando...</div>
    {:then info}
        {#if info != null}
            <div id="turma-selecao">
                <div id="titulo">
                    <span id="codigo">{info.codigo}</span>
                    <span id="nome">&nbsp;- {info.nome}</span>
                </div>

                <div id="meio">
                    <TurmaTextos
                        ementa={info.ementa}
                        preReqs={info.preRequisitos}
                    />
                    <TurmaInfo
                        creditos={info.creditos}
                        qtdTurmas={info.turmas.length}
                        qtdAlunos={info.qtdAlunos}
                        qtdAvaliacoes={info.qtdAvaliacoes}
                        avaliacaoMedia={info.avaliacaoMedia}
                        grauMedio={info.grauMedio}
                    />
                    
                </div>
                <div id="turmas">
                    <span id="selecione">Selecione uma das turmas abaixo</span>
                    <ContainerTurmas
                        disciplina={info.codigo}
                        turmas={info.turmas}
                        {tabelaHorariosUsados}
                        on:submit
                    />
                    <a id="voltar" href="/#" on:click|preventDefault={close}>Voltar à grade</a>
                </div>
            </div>
        {:else}
            <div id="no-info">Disciplina {codigoDisciplina} não encontrada</div>
        {/if}
    {:catch}
        <div id="error">Não foi possível encontrar as informações da disciplina {codigoDisciplina}</div>
    {/await}
</Popup>


<style>
    div {
        box-sizing: border-box;
    }

    #turma-selecao {
        width: min(80vw, 1000px);
        height: 600px;
        padding: 2%;

        background: cyan;
        border-radius: 30px;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
        align-items: stretch;
    }
    /* ------- titulo --------*/
    #titulo {
        width: 100%;
        height: 10%;
        padding: 2%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-start;
        align-items: center;

        border: 3px solid red;
    }

    #codigo {
        font-size: 1.5em;
        font-weight: bold;
    }

    #nome {
        font-size: 1.5em;
        white-space: nowrap;
        overflow: hidden;
    }

    /* ------- meio --------*/
    #meio {
        width: 100%;
        height: 60%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: flex-start;
        border: 3px solid red;
    }

    

    /* ------- turmas --------- */
    #turmas {
        width: 100%;
        height: 30%;
        flex-shrink: 0;

        /* temporario */
        overflow-y: scroll;
        border: 3px solid red;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
        align-items: stretch;
    }

    #no-info {
        padding: 2%;
        background: #aaa;
        border-radius: var(--border-radius);
        text-align: center;
    }
</style>