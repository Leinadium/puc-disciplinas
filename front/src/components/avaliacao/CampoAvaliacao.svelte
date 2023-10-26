<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ItemGenerico, SubmitAvaliacaoEvent } from "../../types/data";

    export let info: ItemGenerico | null = null;
    export let statusMessage: string | null = null;
    let notaAtual: number | null;

    let placeHolder: string;
    $: placeHolder = info?.nota?.toString() ?? "-"

    // eventos
    let dispatch = createEventDispatcher<{submit: SubmitAvaliacaoEvent, remove: null}>();
    const submit = () => {
        if (notaAtual != null) {
            dispatch("submit", {
                avaliacao: notaAtual,
            });
        }
    }
    const remove = () => {
        if (info?.nota != null) {
            dispatch("remove");
        }
    }

    const limpar = () => {
        notaAtual = null;
    }
    $: info, limpar();
</script>

<form id="campo-avaliacao" on:submit|preventDefault={submit}>
    {#if info !== null}
        <div id="avaliacao-nome">
            {#if info.codigo != info.nome}
                <span id="avaliacao-codigo">{info.codigo}</span>
            {/if}
            {info.nome}
        </div>

        <div id="avaliacao-atual">
            <span class="descricao">Avaliação atual:</span>
            {#if info.media !== null}
                <span id="avaliacao-avaliacao">
                    <span class="grande">{info.media?.toFixed(1)}</span> / 5.0
                </span>
                <span id="avaliacao-qtd">
                    de um total de {info.qtd.toFixed(0)} avaliações
                </span>
            {:else}
                <span id="avaliacao-avaliacao">
                    <span class="grande">-</span> / 5.0
                </span>
                <span id="avaliacao-qtd">
                    Nenhuma avaliação
                </span>
            {/if}
        </div>

        <div id="avaliacao-atual">
            <span class="descricao">Sua avaliação:</span>
            <span id="avaliacao-avaliacao">
                <input type="number" placeholder={placeHolder} min=0 max=5 step=1 bind:value={notaAtual}> / 5
            </span>
        </div>

        
        <div class="submit">
            {#if statusMessage != null}
                <span class="status">{statusMessage}</span>
            {/if}
            <div class="submit-botoes">
                {#if notaAtual && notaAtual != info.nota }
                    <input class="submit-botao" type="submit" value="Salvar">
                {/if}
                {#if info.nota !== null}
                    <input class="submit-botao" value="Remover" on:click|preventDefault={remove}>
                {/if}
            </div>
            
        </div>
        
    {:else}
        <div class="explicacao">Selecione uma disciplina ou professor</div>
    {/if}
</form>

<style>
    #campo-avaliacao {
        box-sizing: border-box;
        width: 100%;
        height: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        padding: 15% 5%;
        gap: 10%
    }

    #avaliacao-nome {
        width: 100%;

        display: flex;
        flex-flow: column;
        justify-content: flex-start;
        align-items: center;

        text-align: center;
        font-size: 1.3rem;
    }

    #avaliacao-codigo {
        font-size: 1.5rem;
        font-weight: bold;
    }

    #avaliacao-atual {
        width: 100%;

        display: flex;
        flex-flow: column;
        justify-content: flex-start;
        align-items: center;
    }

    #avaliacao-avaliacao {
        font-size: 1.5rem;
        font-weight: bold;
    }

    .grande {
        font-size: 2rem;
    }

    #avaliacao-qtd {
        font-size: 0.9rem;
        font-style: italic;
    }

    input[type="number"] {
        width: 4rem;
        font-size: 1.5rem;
        font-weight: bold;
        text-align: center;
    }

    .submit {
        width: 70%;
        height: 6rem;
        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-end;
        align-items: center;
        
    }

    .submit-botoes {
        width: 100%;
        height: 3rem;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }

    .submit-botao {
        box-sizing: border-box;
        width: 45%;
        height: 100%;
        font-size: 1.0rem;
        font-weight: bold;
        text-align: center;
        background: #fff;
        border: 1px solid #444;
        border-radius: 10px;
        cursor: pointer;

        transition: background 0.2s;
    }

    .submit-botao:hover {
        background: #ccc;
    }

    .explicacao {
        text-align: center;
        color: #333;
    }

    .status {
        color: #ccc;
        text-align: center;
        font-size: 0.8rem;
    }

</style>