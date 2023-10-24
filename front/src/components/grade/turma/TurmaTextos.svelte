<script lang="ts">
	import type { DisciplinaPreRequisitosInfo } from "../../../types/data";

    export let ementa: string;
    export let preReqs: DisciplinaPreRequisitosInfo[];

    const join = (arr: string[]) => arr.join(', ');
</script>

<div id="textos">
    <span class="minititulo">Ementa</span>
    <span id="ementa">{ementa}</span>
    <span class="minititulo">Pré-requisitos</span>
    <span id="prereqs">
        {#if preReqs.length > 0}
            {#each preReqs as preReq}
                <div class="prereq">
                    <span>{join(preReq.preReqs)}</span>
                </div>
            {/each}
        {:else}
            <span>Sem pré-requisitos</span>
        {/if}
    </span>
</div>

<style>
    div {
        box-sizing: border-box;
    }

    #textos {
        width: 40%;
        height: 100%;
        padding: 1%;
        overflow-y: scroll;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: stretch;

        text-align: justify;

        background: #eee;
        border-radius: 15px;
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

        background: #bbb;
        border-radius: 10px;
        padding: 3%;
        margin-bottom: 5%;
    }

    #prereqs {
        height: 30%;
        font-size: 1.0em;
        overflow-y: scroll;

        background: #bbb;
        border-radius: 10px;
        padding: 3%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        gap: 3%;

        font-size: 0.8em;

    }

    .prereq {
        width: fit-content;
        height: fit-content;
        padding: 1%;
        background: #ddd;
        border-radius: 10px;
    }
</style>