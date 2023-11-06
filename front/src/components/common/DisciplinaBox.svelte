<script lang="ts">
	import { pesosCores, pesosDescricao } from "$lib/utils";
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
    let bkgd: GenericBackground = cursada ? "outline" : "fill";

    let dispatch = createEventDispatcher<{popup: string}>();

    const popup = () => {
        dispatch("popup", info.codigo);
    }

</script>

<GenericBox color="green" background={bkgd} clickCallback={popup}>
    <div class="textos">
        <span id="codigo">{info.codigo.toUpperCase()}</span>
        <span id="nome">{info.nome.toUpperCase()}</span>
    </div>
    
    <div class="icones">
        <div class="quantidades">
            <div class="num yellow" id="creditos" title="Quantidade de créditos">{info.creditos}</div>
            <div class="num white"  id="turmas" title="Quantidade de turmas">{info.qtdTurmas}</div>
            <div class="num green"  id="vagas" title="Quantidade de vagas">{info.qtdVagas}</div>
            {#each pesos as p}
                <div class="num {pesosCores[p]}" title={pesosDescricao[p]}>{p.toUpperCase()}</div>
            {/each}
        </div>
    </div>
</GenericBox>

<style>
    .textos {
        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
    }

    .num {
        display: block;
        box-sizing: border-box;
        min-width: 1.5rem;
        width: fit-content;
        height: 1.5rem;
        padding: 0.2rem 0.3rem 0.2rem 0.3rem;
        border-radius: 1rem;

        margin: 0 0.4rem;

        font-size: 0.8rem;
        text-align: center;
        font-weight: bold;
    }

    #nome {
        font-size: 0.65rem;
        overflow-x: hidden;
        /* white-space: nowrap; */
        text-overflow: ellipsis;
    }

    .icones {
        width: 100%;
        font-size: 0.5rem;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: flex-end;   
    }

    .quantidades {
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
    }


    .yellow { background: #aa0; color: #000; }
    .green { background: #0a0; color: #000;}
    .red { background: #a00; color: #fff;}
    .blue { background: #00a; color: #fff;}
    .white { background: #ccc; color: #000;}
    .orange { background: #a50; color: #000;}
    .purple { background: #a0a; color: #fff;}
    .pink { background: #f0a; color: #fff;}
</style>