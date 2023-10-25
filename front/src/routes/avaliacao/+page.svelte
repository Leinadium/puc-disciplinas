<script lang="ts">
	import CampoAvaliacao from "../../components/avaliacao/CampoAvaliacao.svelte";
	import CampoPesquisa from "../../components/avaliacao/CampoPesquisa.svelte";
	import type { ItemGenerico, ItemGenericoExtra } from "../../types/data";

    let items: ItemGenerico[] = [
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
    let item: ItemGenericoExtra = {
        tipo: "disciplina",
        conteudo: {
            codigo: "ENG1234",
            nome: "Grande disciplina"
        },
        avaliacao: 4.5,
        qtdAvaliacoes: 100
    }

    // TODO: handle clicks
    // interação com o usuario

</script>

<div id="avaliacao-page">
    <div id="avaliacao-container">
        <CampoAvaliacao {item} on:submit/>
        <CampoPesquisa {items} on:click />
    </div>
</div>

<style>
    #avaliacao-page {
        width: 100%;
        height: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    #avaliacao-container {
        box-sizing: border-box;
        max-width: min(80vw, 1000px);
        height: 700px;
        width: 100%;

        display: inline-grid;
        grid-template-columns: 30% 69%;
        grid-template-rows: 100%;
        place-content: stretch;
        gap: 1%;

        border: 3px solid blue;
    }
</style>