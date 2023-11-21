<script lang="ts">
	import { onMount } from "svelte";
	import Popup from "../../common/Popup.svelte";
	import type { DisciplinaInfo, Modificacao, TabelaHorarios } from "../../../types/data";
	import { coletarDisciplinaInfo } from "$lib/api";
	import ContainerTurmas from "./ContainerTurmas.svelte";
	import TurmaTextos from "./TurmaTextos.svelte";
	import TurmaInfo from "./TurmaInfo.svelte";
	import TextoModificacao from "../../common/TextoModificacao.svelte";

    export let codigoDisciplina: string;
    export let tabelaHorariosUsados: TabelaHorarios;
    export let modificacao: Modificacao | null = null;

    let infoApi: Promise<DisciplinaInfo | null>;

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
                        codigo={info.codigo}
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
                </div>
                <span id="atualizacao-turma">
                    <TextoModificacao {modificacao} />
                </span>
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
        height: 650px;
        padding: 2%;

        background: var(--color-main-2);
        border-radius: var(--border-radius);

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

        border-radius: var(--border-radius);
        background-color: var(--color-mono-1);
        color: var(--color-whitef);
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
        height: 50%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: flex-start;
    }

    

    /* ------- turmas --------- */
    #turmas {
        box-sizing: border-box;
        padding: 0.1rem 1rem;

        width: 100%;
        height: 30%;
        flex-shrink: 0;

        overflow-y: scroll;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
        align-items: stretch;

        border-radius: var(--border-radius);
        background-color: var(--color-mono-1);
        color: var(--color-whitef);
    }

    #selecione {
        width: 100%;
        text-align: center;
        margin-bottom: 0.1rem;
    }


    #no-info {
        padding: 2%;
        background: #aaa;
        border-radius: var(--border-radius);
        text-align: center;
    }

    #atualizacao-turma {
        width: 100%;
        text-align: right;
        font-size: 0.8em;
        font-style: italic;
        color: var(--color-whiteff);
    }
</style>