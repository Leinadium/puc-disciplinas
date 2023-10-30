<script lang="ts">
	import { coletarRecomendacoes } from "$lib/api";
    import { userStore } from "$lib/stores";
    import DisciplinaBox from "../../common/DisciplinaBox.svelte";
    import { filtrarPesos, filtrarRecomendacoes, type ModoRecomendacao } from "$lib/recomendacao";
    import type { UIDisciplinaResumo } from "../../../types/ui";
    import type { DisciplinaRecomendacao, EscolhasSimples } from "../../../types/data";
	import ModoBotaoRecomendacao from "./ModoBotaoRecomendacao.svelte";
	import { onMount } from "svelte";

    type Texto = "Carregando recomendações..." | "Sem recomendação para o tipo selecionado."

    export let disciplinas: Map<string, UIDisciplinaResumo>;
    export let escolhidas: EscolhasSimples;
    export let faltaCursar: Set<string> | null;
    export let podeCursar: Set<string> | null;

    /** flag de usuario logado */
    // let isLogged: boolean;
    // $: isLogged = $userStore !== null;

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
        texto = "Carregando recomendações...";
        let req = await coletarRecomendacoes(escolhidas);
        if (req !== null && req.length > 0) {
            disciplinasRecomendadas = req;
            console.log("disciplinasRecomendadas: ", disciplinasRecomendadas.length);
            console.log("req", req.length);
            console.log("podeCursar", podeCursar?.size);
            console.log("faltaCursar", faltaCursar?.size);
            console.log("escolhidasSet", escolhidasSet.size);
            disciplinasExibidas = filtrarRecomendacoes(
                req, 
                qtdRecomendacao, 
                podeCursar, 
                faltaCursar,
                escolhidasSet, 
                modoRecomendacao
            );
        } else {
            console.log("Erro ao carregar as recomendacoes");
        }
        texto = "Sem recomendação para o tipo selecionado.";
    }

    onMount(() => {
        hasMounted = true;
        atualizarRecomendacoes();
    })

</script>

<div id="lista-recomendacao">
    <span>Recomendações</span>
    <ModoBotaoRecomendacao bind:value={modoRecomendacao}/>

    <!-- {#if isLogged} -->
        <div id="lista-disciplinas">
            {#if disciplinasExibidas.length > 0}   
                {#each disciplinasExibidas as disciplina}
                    {#if disciplinas.has(disciplina.cod)}
                        <div class="disciplina-recomendada">
                            <DisciplinaBox
                                info={disciplinas.get(disciplina.cod)}
                                pesos={filtrarPesos(disciplina.pes)}
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
    <!-- {:else}
        <div id="aviso">
            <span>
                Para ativar as recomendações, 
                entre com sua conta do SAU
            </span>
        </div>
    {/if} -->
</div>

<style>
    #lista-recomendacao {
        /* posicionamento da lista no container */
        box-sizing: border-box;
        grid-column: 1 / span 1;
        grid-row: 1 / span 1;

        height: 100%;
        position: relative;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: center;
        
        background: #eee;
        /* para justificar a barra de scroll da grade abaixo*/
        margin-right: 10px;

        border: 3px solid green
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
        color: #555;
    }
</style>