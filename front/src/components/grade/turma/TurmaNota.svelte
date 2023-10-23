<script lang="ts">
    export let descricao: string;
    export let nota: number;
    export let base: number;
    export let textoPequeno: string;
    export let isValid: boolean = true;

    $: notaTrue = isValid ? nota.toFixed(1) : "-";
</script>

<div class="nota">
    <span class="descricao">{descricao}</span>
    <div class="notas">
        <span class="nota-grande">{notaTrue}</span>
        <span class="nota-pequena">/{base}</span>
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