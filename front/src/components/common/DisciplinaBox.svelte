<script lang="ts">
	import { codigoCores, pesosCores, pesosDescricao } from "$lib/utils";
    import type { GenericBackground, GenericColor } from "../../types/style";
	import type { UIDisciplinaResumo, UIPeso, UITag } from "../../types/ui";
    import GenericBox from "./GenericBox.svelte";
    import { createEventDispatcher } from "svelte";

    export let info: UIDisciplinaResumo = {
        codigo: "XXX9999",
        nome: "Sem nome",
        creditos: 0,
        qtdTurmas: 0,
        qtdVagas: 0
    };

    export let pesos: UIPeso[] = [];    // lista com os pesos a serem mostrados

    export let cursada: boolean = false;

    export let eletiva: boolean | null = null;
    let color: GenericColor;
    $: color = cursada ? "yellow" : (eletiva ? "green" : "blue");

    let dispatch = createEventDispatcher<{popup: string}>();

    const popup = () => {
        dispatch("popup", info.codigo);
    }

</script>

<GenericBox color={color} clickCallback={popup}>
    <div class="textos">
        <span class="codigo">
            {info.codigo.toUpperCase()}
        </span>
        
        <span class="nome">{info.nome.toUpperCase()}</span>
        
        {#if info.creditos > 0}
            <span class="creditos">{info.creditos} créditos</span>
        {:else}
            <span class="creditos">1 créditos</span>
        {/if}
    </div>

    <div class="quantidades">
        <div class="quantidade">
            <div class="numero">{info.qtdTurmas}</div>
            <span class="quantidade-texto">Turmas</span>
        </div>
        {#if info.qtdVagas >= 0}
            <div class="quantidade">
                <div class="numero">{info.qtdVagas}</div>
                <span class="quantidade-texto">Vagas</span>
            </div>
        {/if}
    </div>
    
    <div class="icones">
        {#each pesos as p}
            <div class="num num-{p}" title={pesosDescricao[p]}>{codigoCores[p]}</div>
        {/each}
    </div>
</GenericBox>

<style>
    .textos {
        box-sizing: border-box;
        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        
        overflow-x: hidden;
        text-overflow: ellipsis;
        color: #111;
        overflow-y: hidden;
    }

    .codigo {
        font-size: 0.85rem;
        font-weight: bold;
        height: 1.2rem;
    }

    .marca {
        position: absolute;
        top: 5%;
        right: 10%;
        font-size: 1.1rem;
        font-weight: bold;
        color: var(--color-blackf);
    }

    .nome {
        font-size: 0.75rem;
    }

    .creditos {
        font-size: 0.75rem;
        font-style: italic;
    }

    .quantidades {
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }

    .quantidade {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: flex-end;
        gap: 0.2rem;
    }

    .quantidade-texto {
        font-size: 0.65rem;
        padding-bottom: 0.15rem;
    }

    .numero {
        font-size: 0.9rem;
        font-weight: bold;
    }

    .icones {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
    }

    .num {
        display: block;
        box-sizing: border-box;
        min-width: 1.3rem;
        width: fit-content;
        height: 1.3rem;
        padding: 0.15rem 0.3rem 0.2rem 0.2rem;
        border-radius: var(--border-radius);

        margin: 0 0.4rem;

        font-size: 0.75rem;
        text-align: center;
        font-weight: bold;
        color: #EEE;
    }

    .num-c {
        background-color: var(--color-tria-2);
    }

    .num-h {
        background-color: var(--color-tria-3);
    }

    .num-o {
        background-color: var(--color-quad-3);
    }

    .num-p {
        background-color: var(--color-tria-1);
    }

    .num-a {
        background-color: var(--color-quad-2);
    }

</style>