<script lang="ts">
	import Estrelas from "../../common/Estrelas.svelte";

    export let descricao: string;
    export let nota: number;
    export let base: number;
    export let textoPequeno: string;
    export let isValid: boolean = true;
    export let isEstrela: boolean = true;

    $: notaTrue = isValid ? nota.toFixed(1) : "-";
</script>

<div class="nota">
    <span class="descricao">{descricao}</span>
    <div class="notas">
        {#if isEstrela}
            <Estrelas nota={isValid ? nota : 0} enable={false}/>
        {:else}
            <span class="nota-grande">{notaTrue}</span>
            <span class="nota-pequena">/{base}</span>
        {/if}
    </div>
    <div class="texto-pequeno">{textoPequeno}</div>
</div>

<style>
    .nota {
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
    }

    .descricao {
        box-sizing: border-box;
        font-size: 1.2rem;
        font-weight: bold;
        padding-bottom: 2%;
    }

    .notas {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        gap: 1%;

        height: 2rem;
    }

    .nota-grande {
        font-size: 2rem;
        font-weight: bold;
    }

    .nota-pequena {
        font-size: 1.5rem;
    }

    .texto-pequeno {
        font-size: 0.7rem;
        font-style: italic;
    }
</style>