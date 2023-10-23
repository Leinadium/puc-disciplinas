<script lang="ts">
	import type { UIDisciplinaResumo } from "../../types/ui";
    import GenericBox from "./GenericBox.svelte";
    import { createEventDispatcher } from "svelte";

    export let info: UIDisciplinaResumo = {
        codigo: "XXX9999",
        nome: "Sem nome",
        creditos: 0,
        qtdTurmas: 0,
        qtdVagas: 0
    };

    let dispatch = createEventDispatcher();

    const popup = () => {
        dispatch("popup", info.codigo);
    }

</script>

<GenericBox color="green" clickCallback={popup}>
    <div class="textos">
        <span id="codigo">{info.codigo.toUpperCase()}</span>
        <span id="nome">{info.nome.toUpperCase()}</span>
    </div>
    
    <div class="numeros">
        <span id="creditos">{info.creditos}<span class="small">crd</span></span>
        <span id="turmas">{info.qtdTurmas}<span class="small">tur</span></span>
        <span id="vagas">{info.qtdVagas}<span class="small">vgs</span></span>
    </div>
</GenericBox>

<style>

    .textos {
        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        height: 70%;
    }

    .small {
        font-size: 0.55rem;
    }

    #nome {
        font-size: 0.65rem;
        overflow: hidden;
        /* white-space: nowrap; */
        text-overflow: ellipsis;
    }

    .numeros {
        font-size: 0.5rem;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: flex-end;
        height: 30%;
    }
</style>