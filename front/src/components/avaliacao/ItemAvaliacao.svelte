<script lang="ts">
	import type { ItemDisciplina, ItemGenerico, SelectAvaliacaoEvent } from "../../types/data";
    import { createEventDispatcher } from "svelte";
	import type { UITipoAvaliacao } from "../../types/ui";
	import Estrelas from "../common/Estrelas.svelte";

    export let info: ItemGenerico;
    export let tipo: UITipoAvaliacao;
    export let enable: boolean = true;

    let dispatch = createEventDispatcher<{click: SelectAvaliacaoEvent}>();

    function click() {
        if (enable) {
            dispatch("click", {
                tipo: tipo,
                conteudo: info,
            });
        }
    }

</script>

<a class="item-avaliacao" class:disable={!enable} href="/#" on:click|preventDefault={click}>
    <span class="nome">
        {info.nome}
    </span>
    <div class="outros">
        <div class="nota">
            {#if info?.nota}
                <Estrelas nota={info.nota} enable={false} />  {info.nota}  
            {:else}
                <span> - </span>
            {/if}
        </div>
        <span class="nota">
            {#if info?.media}
                <Estrelas nota={info.media} enable={false} /> {info.media}
            {:else}
                <span> - </span>
            {/if}
        </span>
    </div>
</a>

<style>
    .item-avaliacao {
        height: 8%;
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
        gap: 5%;

        text-decoration: none;
        background: #333;
        box-sizing: border-box;
        padding: 0.5rem;
        border-radius: var(--border-radius);
    }

    .item-avaliacao.disable {
        background: #f33 !important;
        cursor: not-allowed !important;
    }

    .nome {
        width: 55%;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        color: #111;
    }

    .outros {
        width: 40%;
        height: 70%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
        color: #000;
    }

    .nota {
        width: 50%;
        height: 100%;
        text-align: center;

        display: flex;
        flex-flow: column;
        justify-content: center;
    }

</style>