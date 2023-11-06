<script lang="ts">
	import { createEventDispatcher, onMount } from "svelte";
	import type { ItemDisciplina, ItemGenerico, ItemProfessor, SelectAvaliacaoEvent } from "../../types/data";
	import type { UITipoAvaliacao } from "../../types/ui";
	import ItemAvaliacao from "./ItemAvaliacao.svelte";

    export let disciplinas: ItemDisciplina[];
    export let professores: ItemProfessor[];

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


    let infosExibidos: ItemGenerico[] = [];
    function pesquisar() {
        const t = texto.toUpperCase();
        const filtro = (i: ItemGenerico) => {
            return t.length == 0 || i.nome.includes(t)
        }
        const items = tipo == "disciplina" ? discGenericos : profGenericos;
        infosExibidos = items.filter(i => filtro(i));
        // ordenando pelo nome
        infosExibidos.sort((a, b) => a.nome.localeCompare(b.nome));

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
        <input 
            type="text" placeholder="Pesquisar"
            bind:value={texto}
        />
        <select name="tipo" id="semestre" bind:value={tipo}>
            <option value="disciplina">Disciplina</option>
            <option value="professor">Professor</option>
        </select>
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

        {#each infosExibidos as x (x.codigo)}
            <ItemAvaliacao {tipo} info={x} on:click={() => selecionar(x)} />
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

        border: 1px solid black;
    }

    #inputs-pesquisa {
        box-sizing: border-box;
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
        width: 70%;
        height: 100%;
        padding: 0.5rem;
        border-radius: 0.5rem;
        border: 1px solid #aaa;
    }
    select {
        box-sizing: border-box;
        width: 20%;
        height: 100%;
        padding: 0.5rem;
        
    }


    #resultados-pesquisa {
        box-sizing: border-box;
        width: 100%;
        height: 100%;
        padding: 1%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: stretch;
        gap: 1rem;

        overflow-y: scroll;
        border: 1px solid purple;
    }

    .header {
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;

        font-weight: bold;
        box-sizing: border-box;
        width: 100%;
        padding: 0.6rem 0.5rem;
        border-radius: 0.5rem;
        background: #aaa;
    }
    .explicacao {
        box-sizing: border-box;
        width: 100%;
        text-align: center;
        color: #777;
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