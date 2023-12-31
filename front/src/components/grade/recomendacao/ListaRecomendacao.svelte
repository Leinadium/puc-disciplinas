<script lang="ts">
	import { coletarRecomendacoes } from "$lib/api";
    import { userStore, hasCurriculo, userEvent } from "$lib/stores";
    import DisciplinaBox from "../../common/DisciplinaBox.svelte";
    import { filtrarPesos, filtrarRecomendacoes, type ModoRecomendacao } from "$lib/recomendacao";
    import type { UIDisciplinaResumo } from "../../../types/ui";
    import type { DisciplinaRecomendacao, EscolhasSimples } from "../../../types/data";
	import ModoBotaoRecomendacao from "./ModoBotaoRecomendacao.svelte";
	import { onMount } from "svelte";
	import { fly } from "svelte/transition";

    type Texto = "Carregando recomendações..." | "Sem recomendação para o tipo selecionado." | "Erro ao carregar as recomendações";

    export let disciplinas: Map<string, UIDisciplinaResumo>;
    export let escolhidas: EscolhasSimples;
    export let faltaCursar: Set<string> | null;
    export let podeCursar: Set<string> | null;

    /** flags para a mensagem */
    let isLogged: boolean;
    $: isLogged = $userStore !== null;
    let showPedido: boolean;
    $: showPedido = !isLogged || !hasCurriculo;

    /** lista de disciplinas recomendadas */
    let disciplinasRecomendadas: DisciplinaRecomendacao[] = [];
    let disciplinasExibidas: DisciplinaRecomendacao[] = [];
    let modoRecomendacao: ModoRecomendacao = "eletivas";
    let texto: Texto = "Carregando recomendações...";
    const qtdRecomendacao = 10;

    /** converte a EscolhasSimples em um set com as disciplinas escolhidas */
    let escolhidasSet: Set<string>;
    $: escolhidasSet = new Set<string>(escolhidas.map(e => e.disciplina));

    // atualiza as recomendacoes, se alguma das variaveis mudar
    // $: disciplinas, escolhidas, faltaCursar, isLogged, atualizarRecomendacoes();
    let hasMounted = false;
    $: escolhidas, atualizarRecomendacoes();
    
    // atualiza o filtro se o modo de recomendacao mudar,
    $: disciplinasExibidas = filtrarRecomendacoes(
            disciplinasRecomendadas,
            qtdRecomendacao,
            podeCursar,
            faltaCursar,
            escolhidasSet,
            modoRecomendacao
        );

    /**
     * Atualiza a lista de recomendacoes.
     * Se o usuario nao estiver logado, nao faz nada.
     */
    async function atualizarRecomendacoes() {
        if (!hasMounted) return;    // evita rodar no servidor. solucao feia mas funciona
        disciplinasRecomendadas = [];
        texto = "Carregando recomendações...";
        let req = await coletarRecomendacoes(escolhidas);
        if (req && req.length > 0) {
            disciplinasRecomendadas = req;
            // console.log("disciplinasRecomendadas: ", disciplinasRecomendadas.length);
            // console.log("req", req.length);
            // console.log("podeCursar", podeCursar?.size);
            // console.log("faltaCursar", faltaCursar?.size);
            // console.log("escolhidasSet", escolhidasSet.size);
            disciplinasExibidas = filtrarRecomendacoes(
                req, 
                qtdRecomendacao, 
                podeCursar, 
                faltaCursar,
                escolhidasSet, 
                modoRecomendacao
            );
            texto = "Sem recomendação para o tipo selecionado.";
        } else {
            console.log("Erro ao carregar as recomendacoes");
            texto = "Erro ao carregar as recomendações";
        }
    }

    const isEletiva = (cod: string): boolean | null => {
        return faltaCursar ? !faltaCursar.has(cod) : null;
    }

    onMount(() => {
        hasMounted = true;
        atualizarRecomendacoes();

        // se o usuario mudar, recalcula as recomendacoes tbm
        userEvent.subscribe(() => {
            atualizarRecomendacoes();
        })
    })

</script>

<div id="lista-recomendacao">
    <div id="recomendacao-upper">
        <span id="titulo-recomendacao">Recomendações</span>
        {#if showPedido}
            <span id="pedido">Carregue seu currículo para melhorar as recomendações.</span>
        {/if}
        <ModoBotaoRecomendacao bind:value={modoRecomendacao}/>
    </div>

    <div id="lista-disciplinas">
        {#if disciplinasExibidas.length > 0}   
            {#each disciplinasExibidas as disciplina}
                {#if disciplinas.has(disciplina.cod)}
                    <div class="disciplina-recomendada" transition:fly|global={{x: 30, duration: 100}}>
                        <DisciplinaBox
                            info={disciplinas.get(disciplina.cod)}
                            pesos={filtrarPesos(disciplina.pes)}
                            eletiva={isEletiva(disciplina.cod)}
                            on:popup
                        />
                    </div>
                {/if}
            {/each}
        {:else}
            <div id="aviso">
                <span>{texto}</span>
            </div>
        {/if}
    </div>

</div>

<style>
    #lista-recomendacao {
        /* posicionamento da lista no container */
        box-sizing: border-box;
        grid-column: 1 / span 1;
        grid-row: 1 / span 1;

        height: 100%;
        position: relative;
        padding: 0.2rem 0;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: center;
        
        background: var(--color-main-2);
        border-radius: var(--border-radius);
    }

    #recomendacao-upper {
        box-sizing: border-box;
        width: 100%;
        padding: 3px 10px 0 10px;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
    }

    #lista-disciplinas {
        box-sizing: border-box;
        height: 100%;
        width: 100%;
        padding: 10px;

        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-start;
        align-items: stretch;
        gap: 1rem;

        overflow-x: scroll;
    }

    .disciplina-recomendada {
        width: fit-content;
        display: flex;
        flex-flow: row nowrap;
        justify-content: stretch;
        align-items: stretch;
        flex-shrink: 0;
    }

    #aviso {
        width: 100%;
        height: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        color: var(--color-mono-2);
    }

    #pedido {
        font-size: 0.85rem;
        font-weight: bold;
        color: var(--color-tria-1);
    }

    #titulo-recomendacao {
        font-size: 1.2rem;
        font-weight: bold;
        color: var(--color-mono-2);
    }
</style>