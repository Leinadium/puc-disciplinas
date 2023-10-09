<script lang="ts">
    import type { DisciplinaBasicaType } from "../../types/disciplinas";
	import DisciplinaBox from "./DisciplinaBox.svelte";

    let disciplinas: DisciplinaBasicaType[] = [];

    let queryStr: string = "";
    let queryLast: string = "";

    function pesquisar(texto: string) {
        queryStr = texto;
        fetch(`http://localhost:8080/disciplinas/pesquisa?query=${texto}`)
            .then(res => res.json())
            .then(res => {
                disciplinas = res.data;
                queryLast = queryStr;
            })
    }

    function pesquisarChange(e: any) {
        pesquisar(e.target.value);
    }

    function pesquisarComDelay(event: any) {
        queryStr = event.target.value;
        const queryNow = queryStr;
        console.log("settimeout para " + queryNow);
        setTimeout(() => {
            if (queryNow === queryStr && queryNow !== queryLast) {
                console.log("disparando query para" + queryNow);
                pesquisar(queryNow);
            }
        }, 500);
    }

</script>


<input type="text" placeholder="Pesquisar disciplina..." 
    on:input={pesquisarComDelay}
    on:change={pesquisarChange}
>

<div class="teste">
    {#each disciplinas as disciplina}
        <DisciplinaBox info={disciplina} />
    {/each}
</div>

<style>
    .teste {
        display: flex;
        flex-flow: row wrap;
        justify-content: flex-start;
    }
</style>