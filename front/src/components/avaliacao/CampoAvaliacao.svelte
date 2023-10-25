<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ItemDisciplina, ItemGenericoExtra, SubmitAvaliacaoEvent } from "../../types/data";

    export let item: ItemGenericoExtra;
    let codigo: string = item.tipo == "disciplina" ? (item.conteudo as ItemDisciplina).codigo : "";
    let value: number = 0;
    let dispatch = createEventDispatcher<{submit: SubmitAvaliacaoEvent}>();

    const submit = () => {
        dispatch("submit", {
            avaliacao: value,
        });
    }

</script>

<form id="campo-avaliacao" on:submit>
    <div id="avaliacao-nome">
        {#if codigo}
            <span id="avaliacao-codigo">{codigo}</span>
        {/if}
        {item.conteudo.nome}
    </div>

    <div id="avaliacao-atual">
        <span class="descricao">Avaliação atual:</span>
        <span id="avaliacao-avaliacao">
            <span class="grande">{item.avaliacao.toFixed(1)}</span> / 5.0
        </span>
        <span id="avaliacao-qtd">
            de um total de {item.qtdAvaliacoes.toFixed(0)} avaliações
        </span>
    </div>

    <div id="avaliacao-atual">
        <span class="descricao">Sua avaliação:</span>
        <span id="avaliacao-avaliacao">
            <input type="number" min=0 max=5 step=1 bind:value={value}> / 5
        </span>
    </div>

    <input type="submit" value="Enviar">
</form>

<style>
    #campo-avaliacao {
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

        font-size: 1.5rem;
    }

    #avaliacao-codigo {
        font-size: 1.3rem;
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

    input[type="submit"] {
        width: 70%;
        height: 3rem;
        font-size: 1.5rem;
        font-weight: bold;
        text-align: center;
        background: #fff;
        border: 1px solid #444;
        border-radius: 10px;
        cursor: pointer;

        transition: background 0.2s;
    }

    input[type="submit"]:hover {
        background: #ccc;
    }

</style>