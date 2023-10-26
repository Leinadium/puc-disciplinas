<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { UIDisciplinaResumo } from "../../../types/ui";
	import DisciplinaBox from "../../common/DisciplinaBox.svelte";
	
    export let disciplinas: Map<string, UIDisciplinaResumo>;
    export let faltaCursar: Set<string>;
    export let podeCursar: Set<string>;
    
    type ModoPesquisa = "codigo" | "nome" | null;

    /** conjunto contendo as iniciais das disciplinas */
    let iniciais: Set<string> = new Set<string>();
    
    /** query string da ultima pesquisa */
    let queryLast: string = "";

    /** lista com o codigo das disciplinas a ser exibidas */
    let disciplinasExibidas: string[] = [];
    let modoPesquisa: ModoPesquisa = null;

    /** filtro de disciplinas nao cursadas que estao no curriculo */
    function filtroFaltaCursar(s: string) {
        return faltaCursar.size == 0 || faltaCursar.has(s)
    };
    /** filtro de disciplina que o usuario pode cursar */
    function filtroPodeCursar(s: string){
        return podeCursar.size == 0 || podeCursar.has(s)
    };
    
    /** Executa a pesquisa,a atualizando o valor do resultado*/
    function pesquisar(texto: string) {
        // se nao tiver nenhum texto util, exibe as disciplinas que faltam cursar
        if (texto.length < 2) {
            disciplinasExibidas = Array.from(disciplinas.keys()).filter((c) => filtroFaltaCursar(c));
            // ordena alfabeticamente pelo codigo
            disciplinasExibidas.sort();
            modoPesquisa = null;
            return;
        }

        texto = texto.toUpperCase();

        // se a query comecar com uma inicial e depois for composta só de numeros, entao a pesquisa deve ser feita pelo codigo
        if (iniciais.has(texto.substring(0, 3)) && texto.substring(3).match(/^[0-9]*$/)) {
            const filtro = (c: string) => c.startsWith(texto) && filtroPodeCursar(c);
            disciplinasExibidas = Array.from(disciplinas.keys()).filter((c) => filtro(c));
            // ordena alfabeticamente pelo codigo
            disciplinasExibidas.sort();
            modoPesquisa = "codigo";
        }

        // se nao, a pesquisa deve ser feita pelo nome
        else {
            const filtro = (d: UIDisciplinaResumo) => d.nome.toUpperCase().includes(texto) // && filtroPodeCursar(d.codigo);
            disciplinasExibidas = Array.from(disciplinas.values()).filter((d) => filtro(d)).map((d) => d.codigo);
            // ordena alfabeticamente pelo nome
            disciplinasExibidas.sort(
                (a: string, b: string) => disciplinas.get(a)!.nome.localeCompare(disciplinas.get(b)!.nome)
            );   
            modoPesquisa = "nome";
        }
        console.log(modoPesquisa, disciplinasExibidas.length);
    }

    /** Executa uma pesquisa (se for um novo valor) com o resultado */
    function pesquisarChange(e: Event) {
        let value = (e.target as HTMLInputElement).value;
        if (value !== queryLast) {
            queryLast = value;
            pesquisar(value);
        }
    }

    function refresh() {
        disciplinas.forEach((d: UIDisciplinaResumo) => {
            iniciais.add(d.codigo.toUpperCase().substring(0, 3)); // salva a inicial
        });
        pesquisar('');
    }
    $: disciplinas, refresh();

    let dispatch = createEventDispatcher();

</script>

<div id="lista-selecao">
    <div id="pesquisa">
        <input
            type="text" placeholder="ENGXXXX"
            on:input={pesquisarChange}
            on:change={pesquisarChange}
        />
    </div>

    <div id="resultados">
        {#each disciplinasExibidas as cod}
            {#if disciplinas.has(cod)}
                <div class="disciplina-selecao">
                    <DisciplinaBox 
                        info={disciplinas.get(cod)}
                        on:popup
                    />
                </div>
            {/if}
        {/each}
    </div>
</div>


<style>
    #lista-selecao {
        /* posicionamento da grade no container */
        grid-column: 2 / span 1;
        grid-row: 2 / span 1;

        box-sizing: border-box;
        width: 100%;
        height: 100%;
        overflow-y: scroll;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
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
        justify-content: flex-start;
        align-items: stretch;
        gap: 1rem;
    }

    .disciplina-selecao {
        width: 100%;
        height: 20%;
    }

    
</style>