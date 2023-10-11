<script lang="ts">
	import { pesquisarDisciplinas } from "$lib/api";
	import type { DisciplinaBasica } from "../../types/disciplinas";
    import type { ErrorApi } from "../../types/api";
	import DisciplinaBox from "../common/DisciplinaBox.svelte";

    let disciplinas: DisciplinaBasica[] = [];

    let queryStr: string = "";
    let queryLast: string = "";

    /** Faz a pesquisa */
    const pesquisar = (texto: string) => {
        queryStr = texto;
        pesquisarDisciplinas(texto)
            .then((ds: DisciplinaBasica[]) => {
                disciplinas = ds;
                queryLast = queryStr;
            })
            .catch((e: ErrorApi) => {
                console.log(e.message)
            })
    }

    const pesquisarChange = (e: Event) => {
        pesquisar((e.target as HTMLInputElement).value);
    }

    const pesquisarComDelay = (e: Event) => {
        queryStr = (e.target as HTMLInputElement).value;
        const queryNow = queryStr;
        setTimeout(() => {
            if (queryNow == queryStr && queryNow !== queryLast) {
                pesquisar(queryNow);
            }
        }, 500);
    }
</script>


<div id="lista-selecao">
    <div id="pesquisa">
        <input
            type="text" placeholder="ENGXXXX"
            on:input={pesquisarComDelay}
            on:change={pesquisarChange}
        />
    </div>

    <div id="resultados">
        {#each disciplinas as d}
            <DisciplinaBox info={d} />
        {/each}
    </div>
</div>


<style>
    #lista-selecao {
        width: 15%;
        height: 100%;
        overflow-y: scroll;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 1rem;

        border: 1px solid black;
    }

    #pesquisa {
        width: 100%;
    }

    input {
        box-sizing: border-box;
        margin: 0% 5%;
        width: 90%;
    }

    #resultados {
        width: 100%;
        height: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 1rem;
    }

    
</style>