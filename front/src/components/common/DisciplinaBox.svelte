<script lang="ts">
	import type { DisciplinaBasicaType } from "../../types/disciplinas";
    import GenericBox from "./GenericBox.svelte";

    export let info: DisciplinaBasicaType = {
        codigo: "INF9999",
        nome: "Técnicas de Programação II",
        creditos: 6,
        qtdTurmas: 3,
        qtdVagas: 120
    };

</script>

<GenericBox color="green">
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