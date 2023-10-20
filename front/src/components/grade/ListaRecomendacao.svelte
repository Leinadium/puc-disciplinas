<script lang="ts">
	import { coletarRecomendacoes } from "$lib/api";
    import { userStore } from "$lib/stores";
    import DisciplinaBox from "../common/DisciplinaBox.svelte";
    import { createEventDispatcher } from "svelte";
    import { filtrarRecomendacoes, type ModoRecomendacao } from "$lib/recomendacao";
    import type { UIDisciplinaResumo } from "../../types/ui";
    import type { DisciplinaRecomendacao, Escolha } from "../../types/data";

    export let disciplinas: Map<string, UIDisciplinaResumo>;
    export let escolhidas: Escolha[];
    export let faltaCursar: Set<string>;
    export let podeCursar: Set<string>;

    /** dispatcher de eventos */
    let dispatch = createEventDispatcher();

    /** flag de usuario logado */
    $: isLogged = $userStore !== null;

    /** lista de disciplinas recomendadas */
    let disciplinasRecomendadas: DisciplinaRecomendacao[] = [];
    let modoRecomendacao: ModoRecomendacao = "eletivas";
    const qtdRecomendacao = 10;

    // atualiza as recomendacoes, se alguma das variaveis mudar
    $: disciplinas, escolhidas, faltaCursar, isLogged, atualizarRecomendacoes();
    
    /**
     * Atualiza a lista de recomendacoes.
     * Se o usuario nao estiver logado, nao faz nada.
     */
    async function atualizarRecomendacoes() {
        if (!isLogged) return;
        let req = await coletarRecomendacoes(escolhidas);
        if (req !== null) {
            disciplinasRecomendadas = filtrarRecomendacoes(
                req, qtdRecomendacao, podeCursar, faltaCursar, modoRecomendacao
            );
            console.log(disciplinasRecomendadas);
        } else {
            console.log("Erro ao carregar as recomendacoes");
        }
    }

</script>

<div id="lista-recomendacao">
    {#if isLogged}
        <div id="botoes">
            {#each disciplinasRecomendadas as disciplina}
                {#if disciplinas.has(disciplina.cod)}
                    <DisciplinaBox
                        info={disciplinas.get(disciplina.cod)}
                        on:click={() => dispatch("click", disciplina.cod)}
                    />
                {/if}
            {/each}
        </div>
        <div id="lista">

        </div>    
    {:else}
        <div class="aviso">
            <span>
                Para ativar as recomendações, 
                entre com sua conta do SAU
            </span>
        </div>
    {/if}
</div>

<style>
    
</style>