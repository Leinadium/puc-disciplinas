<script lang="ts">
	import { createEventDispatcher, onMount } from "svelte";
	import type { ItemDisciplina, ItemGenerico, ItemProfessor, SelectAvaliacaoEvent } from "../../types/data";
	import type { UITipoAvaliacao } from "../../types/ui";
	import ItemAvaliacao from "./ItemAvaliacao.svelte";
	import { coletarDisciplinaInfo } from "$lib/api";

    export let disciplinas: ItemDisciplina[];
    export let professores: ItemProfessor[];
    export let cursadas: Set<string> | null = null;

    // binds
    let tipo: UITipoAvaliacao = "disciplina";
    let texto: string = "";
    
    let discGenericos: ItemGenerico[];
    let profGenericos: ItemGenerico[];
    $: discGenericos = disciplinas.map(d => ({
        codigo: d.codigo,
        nome: `[${d.codigo}] ${d.nome}`.toUpperCase(),
        nota: d.nota,
        media: d.media,
        qtd: d.qtd,
    }));
    $: profGenericos = professores.map(p => ({
        codigo: p.nome,
        nome: p.nome.toUpperCase(),
        nota: p.nota,
        media: p.media,
        qtd: p.qtd,
    }));

    // ordena os itens
    // (talvez nao seja a maneira mais eficiente)
    function mySort(a: ItemGenerico, b: ItemGenerico) {
        let x: string = a.nome;
        let y: string = b.nome;
        if (tipo == "disciplina") {
            x = (cursadas?.has(a.codigo) ? "0" : "1") || "1" + x;
            y = (cursadas?.has(b.codigo) ? "0" : "1") || "1" + y;
        }
        return x.localeCompare(y);
    }

    let infosExibidos: ItemGenerico[] = [];
    function pesquisar() {
        const t = texto.toUpperCase();
        const filtro = (i: ItemGenerico) => {
            return t.length == 0 || i.nome.includes(t)
        }
        const items = tipo == "disciplina" ? discGenericos : profGenericos;
        infosExibidos = items.filter(i => filtro(i));
        // ordenando pelo nome
        infosExibidos.sort(mySort);

    }
    // se o texto ou tipo mudarem, atualiza os resultados
    $: texto, tipo, pesquisar();
    
    // evento
    let dispatch = createEventDispatcher<{click: SelectAvaliacaoEvent}>();
    function selecionar(conteudo: ItemGenerico) {
        dispatch('click', {conteudo: conteudo, tipo: tipo});
    }

    // se as disciplinas ou professores mudarem, pesquisa de novo
    $: disciplinas, professores, pesquisar();
    

</script>

<div id="campo-pesquisa">
    <div id="inputs-pesquisa">
        <select name="tipo" id="semestre" bind:value={tipo}>
            <option value="disciplina">Disciplina</option>
            <option value="professor">Professor</option>
        </select>
        <input 
            type="text" placeholder="Pesquisar"
            bind:value={texto}
        />
    </div>
    

    <div id="resultados-pesquisa">
        <div class="header">
            <span class='nome'>Nome</span>
            <span class='notas'>Nota</span>
            <span class='notas'>Média</span>
        </div>

        {#if infosExibidos.length == 0}
            <span class="explicacao">Nenhum resultado encontrado</span>
        {/if}

        {#if !cursadas && tipo == "disciplina"}
            <span class="explicacao"> Cadastre seu histórico para avaliar disciplinas.</span>
        {/if}

        {#each infosExibidos as x (x.codigo)}
            <ItemAvaliacao
                {tipo}
                info={x}
                enable={tipo == "professor" || cursadas?.has(x.codigo)}
                on:click={() => selecionar(x)}
            />
        {/each}
    </div>
</div>

<style>
    #campo-pesquisa {
        box-sizing: border-box;
        grid-column: 2 / span 1;
        grid-row: 1 / span 1;
        width: 100%;
        height: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: stretch;
        gap: 1rem;
    }

    #inputs-pesquisa {
        box-sizing: border-box;
        padding: 0 0.5rem;
        width: 100%;
        margin-top: 2%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    input {
        box-sizing: border-box;
        width: 80%;
        height: 100%;
        padding: 0.5rem;
        border-radius: var(--border-radius);
        border: none;
        color: var(--color-whitef);
        background: var(--color-tria-3);
    }
    select {
        box-sizing: border-box;
        width: 20%;
        height: 100%;
        padding: 0.5rem;
        border-radius: var(--border-radius);
        border: none;
        font-weight: bold;
        color: var(--color-whitef);
        background: var(--color-tria-3);
    }


    #resultados-pesquisa {
        box-sizing: border-box;
        position: relative;
        width: 100%;
        height: 100%;
        padding: 1%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: stretch;
        gap: 1rem;

        overflow-y: scroll;
    }

    .header {
        position: sticky;
        top: 0;
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;

        z-index: 2;

        font-weight: bold;
        box-sizing: border-box;
        width: 100%;
        padding: 0.6rem 0.5rem;
        border-radius: var(--border-radius);
        background: var(--color-mono-4);
    }
    .explicacao {
        box-sizing: border-box;
        width: 100%;
        text-align: center;
        color: var(--color-whiteff);
    }

    .nome {
        width: 60%;
        text-align: center
    }
    .notas {
        width: 20%;
        text-align: center
    }
</style>