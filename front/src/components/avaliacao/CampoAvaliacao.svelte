<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ItemGenerico, SubmitAvaliacaoEvent } from "../../types/data";
	import Estrelas from "../common/Estrelas.svelte";

    export let info: ItemGenerico | null = null;
    export let statusMessage: string | null = null;

    let notaAtual: number = 0;
    $: notaAtual = info?.nota || 0;
    $: console.log(notaAtual);

    let placeHolder: string;
    $: placeHolder = info?.nota?.toString() ?? "-"

    function atualizar(e: CustomEvent<number>) {
        notaAtual = e.detail;
    }

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

</script>

<form id="campo-avaliacao" on:submit|preventDefault={submit}>
    {#if info !== null}
        <div id="avaliacao-nome">
            {info.nome}
        </div>

        <div id="avaliacao-atual">
            <span class="descricao">Avaliação atual:</span>
            {#if info.media !== null}
                <div class="avaliacao-avaliacao"><Estrelas nota={info.media} enable={false}/></div>
                <span class="avaliacao-qtd">de um total de {info.qtd.toFixed(0)} avaliações</span>
            {:else}
                <div class="avaliacao-avaliacao"><Estrelas nota={0} enable={false} /></div>
                <span class="avaliacao-qtd">Nenhuma avaliação</span>
            {/if}
        </div>

        <div id="avaliacao-atual">
            <span class="descricao">Sua avaliação:</span>
            <span class="avaliacao-avaliacao grande">
                <Estrelas nota={notaAtual} enable={true} on:nota={atualizar}/>
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
        gap: 10%;

        background: var(--color-tria-3);
        color: var(--color-whitef);
        border-radius: var(--border-radius);
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

    #avaliacao-atual {
        width: 100%;

        display: flex;
        flex-flow: column;
        justify-content: flex-start;
        align-items: center;
    }

    .avaliacao-avaliacao {
        height: 1.5rem;
        margin-top: 0.5rem;
        margin-bottom: 0.3rem;
    }

    .grande {
        height: 2.5rem;
    }

    .avaliacao-qtd {
        font-size: 0.9rem;
        font-style: italic;
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
        color: var(--color-whitef);
        background: var(--color-mono-1);
        border: none;
        border-radius: var(--border-radius);
        cursor: pointer;

        transition: background 0.2s;
    }

    .submit-botao:hover {
        background: var(--color-main-2);
    }

    .explicacao {
        text-align: center;
        color: var(--color-whiteff);
    }

    .status {
        color: var(--color-whiteff);
        text-align: center;
        font-size: 0.8rem;
    }

</style>