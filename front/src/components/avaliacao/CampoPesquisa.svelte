<script lang="ts">
	import type { ItemGenerico } from "../../types/data";
	import type { UITipoAvaliacao } from "../../types/ui";
	import ItemAvaliacao from "./ItemAvaliacao.svelte";

    export let infos: ItemGenerico[] = [
        {
            tipo: "disciplina",
            conteudo: {
                codigo: "ENG1234",
                nome: "Grande disciplina"
            }
        },
        {
            tipo: "professor",
            conteudo: {
                nome: "Grande professor"
            }
        },
        {
            tipo: "disciplina",
            conteudo: {
                codigo: "ENG2345",
                nome: "Pequena materia"
            }
        },
        {
            tipo: "professor",
            conteudo: {
                nome: "Meio professor"
            }
        },
        {
            tipo: "disciplina",
            conteudo: {
                codigo: "ENG4567",
                nome: "Pequena disciplina"
            }
        },
        {
            tipo: "professor",
            conteudo: {
                nome: "Grade nada"
            }
        }
    ];

    // binds
    let tipo: UITipoAvaliacao = "todos";
    let texto: string = "";
    
    const filtroGeral = (i: ItemGenerico) => {
        return tipo === "todos" || i.tipo === tipo
    }

    let infosExibidos: ItemGenerico[] = infos;
    function pesquisar() {

        if (texto.length == 0) {
            infosExibidos = infos;
            return;
        }

        texto = texto.toUpperCase();

        const filtro = (i: ItemGenerico) => {
            if (i.tipo == "disciplina") {
                const d = i.conteudo as {codigo: string, nome: string};
                return (d.codigo + d.nome).toUpperCase().includes(texto);
            } else if (i.tipo == "professor") {
                const p = i.conteudo as {nome: string};
                return p.nome.toUpperCase().includes(texto);
            } else {
                return false;
            }
        }

        infosExibidos = infos.filter(i => filtro(i) && filtroGeral(i));
        console.log(infosExibidos);
    }

    $:tipo, pesquisar()
</script>

<div id="campo-pesquisa">
    <div id="inputs-pesquisa">
        <input 
            type="text" placeholder="ENGXXXX"
            on:input={pesquisar}
            on:change={pesquisar}
            bind:value={texto}
        />
        <select name="tipo" id="semestre" bind:value={tipo}>
            <option value="disciplina">Disciplina</option>
            <option value="professor">Professor</option>
            <option value="todos">Todos</option>
        </select>
    </div>
    

    <div id="resultados-pesquisa">
        {#each infosExibidos as x (x.conteudo.nome)}
            <ItemAvaliacao info={x} on:click />
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
</style>