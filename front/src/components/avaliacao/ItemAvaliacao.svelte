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
            {#if !enable}
                <span class="nao-cursada">Não cursada</span>
            {:else if info?.nota}
                <Estrelas nota={info.nota} enable={false} />  
            {:else}
                <span> - </span>
            {/if}
        </div>
        <span class="nota">
            {#if info?.media}
                <Estrelas nota={info.media} enable={false} />
            {:else}
                <span> - </span>
            {/if}
        </span>
    </div>
</a>

<style>
    .item-avaliacao {
        position: relative;
        height: 8%;
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
        gap: 5%;

        text-decoration: none;
        background: var(--color-main-2);
        color: var(--color-whitef);
        box-sizing: border-box;
        padding: 0.5rem;
        border-radius: var(--border-radius);
    }

    .item-avaliacao.disable {
        background: var(--color-tria-2) !important;
        cursor: not-allowed !important;
    }

    .nome {
        width: 55%;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .outros {
        width: 40%;
        height: 70%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }

    .nota {
        width: 50%;
        height: 100%;
        text-align: center;

        display: flex;
        flex-flow: column;
        justify-content: center;
    }

    .nao-cursada {
        font-size: 0.8rem;
        text-align: center;
    }

</style>