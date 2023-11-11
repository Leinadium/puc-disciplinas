<script lang="ts">
	import { geraLinkEmenta } from "$lib/utils";
	import type { DisciplinaPreRequisitosInfo } from "../../../types/data";

    export let ementa: string;
    export let preReqs: DisciplinaPreRequisitosInfo[] | null = null;
    export let codigo: string;
    console.log(preReqs);

    const join = (arr: string[]) => arr.join(', ');
</script>

<div id="textos">
    <span class="minititulo">Ementa</span>
    <span id="ementa">{ementa}</span>
    <span class="minititulo">Pré-requisitos</span>
    {#if preReqs && preReqs.length > 0}
        <div id="prereqs">    
            <a 
                id="prereqs-explicacao"    
                target="_blank"
                href={geraLinkEmenta(codigo)}
            >
                Nem todas os códigos podem estar sendo exibidos corretamente. 
                Confira aqui.
            </a>
            
            {#each preReqs as preReq}
                <div class="prereq">
                    <span>{join(preReq.preReqs)}</span>
                </div>
            {/each}
        </div>
            {:else}
        <span>Sem pré-requisitos</span>
    {/if}
    
</div>

<style>
    div {
        box-sizing: border-box;
    }

    #textos {
        width: 50%;
        height: 100%;
        padding: 1rem;
        overflow-y: scroll;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: stretch;

        text-align: justify;

        border-radius: var(--border-radius);
        background-color: var(--color-mono-1);
        color: var(--color-whitef);
    }

    .minititulo {
        height: 10%;
        font-size: 1.2em;
        font-weight: bold;
        text-align: center;
    }

    #ementa {
        height: 45%;
        font-size: 1.0em;
        overflow-y: scroll;

        border-radius: var(--border-radius);
        border: 3px solid var(--color-main-2);
        padding: 3%;
        margin-bottom: 5%;
    }

    #prereqs {
        box-sizing: border-box;

        height: 30%;
        font-size: 1.0em;
        overflow-y: scroll;

        border-radius: var(--border-radius);
        border: 3px solid var(--color-main-2);
        padding: 3%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        gap: 0.5rem;

        font-size: 0.8em;
    }

    .prereq {
        width: fit-content;
        height: fit-content;
        padding: 0.4rem;
        background: var(--color-main-1);
        border-radius: var(--border-radius);
    }

    #prereqs-explicacao {
        width: 100%;
        text-align: center;
        font-size: 0.8em;
        font-style: italic;
        text-decoration: none;
        color: var(--color-whiteff) !important;
        margin-bottom: 0.5rem;
    }

</style>