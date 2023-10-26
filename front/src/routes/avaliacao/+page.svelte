<script lang="ts">
	import { armazenarAvaliacao, coletarAvaliacao, removerAvaliacao } from "$lib/api";
	import { onMount } from "svelte";
	import CampoAvaliacao from "../../components/avaliacao/CampoAvaliacao.svelte";
	import CampoPesquisa from "../../components/avaliacao/CampoPesquisa.svelte";
	import type { ItemDisciplina, ItemGenerico, ItemProfessor, SelectAvaliacaoEvent, SubmitAvaliacaoEvent } from "../../types/data";
	import type { UITipoAvaliacao } from "../../types/ui";

    let disciplinas: ItemDisciplina[] = [];
    let professores: ItemProfessor[] = [];

    let escolhido: ItemGenerico | null = null;
    let tipo: UITipoAvaliacao | null = null;
    let statusMessage: string | null = null;

    async function submit(e: CustomEvent<SubmitAvaliacaoEvent>) {
        try {
            if (tipo == null || escolhido == null) {
                throw new Error("Escolha uma disciplina ou professor");
            }
            // chama a API
            await armazenarAvaliacao(tipo, escolhido.codigo, e.detail.avaliacao)
            // exibe a mensagem de sucesso
            statusMessage = "Avaliação enviada com sucesso!";
            // atualiza os dados
            await atualizar();
            escolhido.nota = e.detail.avaliacao;
        
        } catch (e: any) {
            // exibe o erro no campo de avaliacao mesmo
            statusMessage = e.message;
        } finally {
            // remove a mensagem de status depois de 3 segundos
            setTimeout(() => {
                statusMessage = null;
            }, 3000);
        }
    }

    async function remove() {
        if (escolhido && escolhido.nota != null && tipo != null) {
            try {
                // faz a chamada da api
                await removerAvaliacao(tipo, escolhido.codigo);
                // exibe a mensagem de scuesso
                statusMessage = "Avaliação removida com sucesso!";
                // atualiza os dados
                await atualizar();
                escolhido.nota = null;
            
            } catch (e: any) {
                // exibe o erro no campo de avaliacao mesmo
                statusMessage = e.message;
            } finally {
                // remove a mensagem de status depois de 3 segundos
                setTimeout(() => {
                    statusMessage = null;
                }, 3000);
            }
        }
    }

    function select(e: CustomEvent<SelectAvaliacaoEvent>) {
        escolhido = e.detail.conteudo;
        tipo = e.detail.tipo;
    }

    async function atualizar() {
        try {
            let data = await coletarAvaliacao();
            if (data != null) {
                disciplinas = data.disciplinas;
                professores = data.professores;
                console.log(disciplinas);
                console.log(professores);
            }
        } catch (e) {
            console.log(e);
        }
    }

    onMount(() => {
        atualizar();
    })


</script>

<div id="avaliacao-page">
    <div id="avaliacao-container">
        <CampoAvaliacao 
            info={escolhido}
            {statusMessage} 
            on:submit={submit}
            on:remove={remove}
        />
        <CampoPesquisa
            {disciplinas}
            {professores}
            on:click={select}
        />
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