<script lang="ts">
	import type { ItemDisciplina, ItemGenerico, SelectAvaliacaoEvent } from "../../types/data";
    import { createEventDispatcher } from "svelte";
	import type { UITipoAvaliacao } from "../../types/ui";

    export let info: ItemGenerico;
    export let tipo: UITipoAvaliacao;

    let dispatch = createEventDispatcher<{click: SelectAvaliacaoEvent}>();

    function click() {
        dispatch("click", {
            tipo: tipo,
            conteudo: info,
        });
    }

</script>

<a class="item-avaliacao" href="/#" on:click|preventDefault={click}>
    <span class="nome">
        {info.nome}
    </span>
    <div class="outros">
        {#if info.nota != null}
            <span class="nota">{info.nota.toFixed(1)}</span>
        {:else}
            <span class="nota">Não avaliada</span>
        {/if}

        <span class="nota-media">
            {#if info.media != null}    
                {info.media.toFixed(1)}
                ({info.qtd})
            {:else}
                -
            {/if}
        </span>
    </div>
</a>

<style>
    .item-avaliacao {
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
        gap: 5%;

        text-decoration: none;
        background: #eee;
        box-sizing: border-box;
        padding: 0.5rem;
        border-radius: 0.5rem;
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

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
        color: #000;
    }

    .outros > span {
        width: 50%;
        text-align: center;
    }

</style>