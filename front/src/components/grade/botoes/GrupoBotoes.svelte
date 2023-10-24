<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import BotaoGenerico from "./BotaoGenerico.svelte";
	import JanelaSalvar from "./JanelaSalvar.svelte";
	import type { GradeAtualExtra, SalvarGradeEvent } from "../../../types/data";
	import { armazenarGrade } from "$lib/api";
	import type { LoadingStatus } from "../../../types/api";
	import { page } from "$app/stores";

    export let gradeAtual: GradeAtualExtra;
    export let enableSalvar: boolean = false;

    let status: LoadingStatus | null = null;
    let linkCodigo: string | null = null;

    const currentUrl = $page.url.pathname;

    let dispatch = createEventDispatcher();
    async function salvar() {
        status = "loading";
        try {
            let codigo = await armazenarGrade(gradeAtual);
            if (codigo !== null) {
                linkCodigo = currentUrl + "?g=" + codigo;
                status = "success";
            } else {
                status = "error";
            }
        } catch (e) {
            console.log(e);
            status = "error"
        }
    }
    const limpar = () => {dispatch("limpar")}
</script>

{#if status !== null}
    <JanelaSalvar 
        status={status}
        link={linkCodigo}
    />
{/if}

<div class="grupo-botoes">
    <BotaoGenerico enable={enableSalvar} text="Salvar grade" on:click={salvar} />
</div>

<style>
    .grupo-botoes {
        /* posicionamento da grade no container */
        grid-column: 2 / span 1;
        grid-row: 1 / span 1;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 10%;

        box-sizing: border-box;
        width: 100%;
        height: 100%;
        background-color: cyan;
    }
</style>