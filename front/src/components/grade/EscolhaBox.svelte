<script lang="ts">
	import type { UIEscolha } from "../../types/ui";
    import GenericBox from "../common/GenericBox.svelte";

    export let info: UIEscolha = {
        codigo: "INF9999",
        nome: "Técnicas de Programação II",
        turma: "3WA",
        professor: "Fulano de Tal"
    }

    // export let hourHeight: number = 2;

    // // 10px de padding
    // const height: string = (hourHeight * 50 - 10) + "px";
    // const width: string = "100%";

</script>

<GenericBox color="yellow">
    <div class="wrapper">
        <div class="upper">
            <span id="codigo">{info.codigo.toUpperCase()}</span>
            <span id="turma">{info.turma.toUpperCase()}</span>
        </div>
        
        <div class="lower">
            <span id="nome">{info.nome}</span>
            <span id="professor">{info.professor}</span>
        </div>
    </div>
</GenericBox>

<style>
    .wrapper {
        width: 100%;
        height: 100%;
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 0.2rem;
    }

    span {
        /* width: 90%; */
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }

    .upper {
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;

        font-weight: bold;
    }

    .lower {
        width: 100%;
        display: flex;
        flex-flow: column nowrap;
        justify-content: space-around;
        align-items: center;
        gap: 0.2rem;
    }

    #codigo {
        font-size: 0.75rem;
    }

    #turma {
        font-size: 0.65rem;
        font-style: italic;
    }

    #nome {
        width: 100%;
        font-size: 0.75rem;
    }

    #professor {
        font-size: 0.65rem;
        font-style: italic;
    }
</style>