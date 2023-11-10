<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import BotaoGenerico from "./BotaoGenerico.svelte";
	import JanelaSalvar from "./JanelaSalvar.svelte";
	import type { GradeAtualExtra, SalvarGradeEvent } from "../../../types/data";
	import { armazenarGrade } from "$lib/api";
	import type { LoadingStatus } from "../../../types/api";
	import { page } from "$app/stores";
	import { userStore } from "$lib/stores";
	import TextoAcima from "./TextoAcima.svelte";

    export let gradeAtual: GradeAtualExtra;
    export let enableSalvar: boolean = false;
    export let enableLimpar: boolean = false;

    let status: LoadingStatus | null = null;
    let linkCodigo: string | null = null;

    const currentUrl = $page.url.pathname;

    let dispatch = createEventDispatcher();
    async function salvar() {
        if ($userStore === null) {
            status = "unauthorized";
            return;
        }

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

    let qtdDisciplinas: number;
    $: qtdDisciplinas = gradeAtual.escolhas.length;
    let qtdCreditos: number;
    $: qtdCreditos = gradeAtual.escolhas.reduce((acc, cur) => acc + cur.creditos > 0 ? cur.creditos : 0, 0);

</script>

{#if status !== null}
    <JanelaSalvar 
        status={status}
        link={linkCodigo}
        on:close={() => {status = null}}
    />
{/if}

<div class="grupo-botoes">
    <TextoAcima disciplinas={qtdDisciplinas} creditos={qtdCreditos} />
    <BotaoGenerico color="blue" enable={enableSalvar} text="💾 Salvar grade" on:click={salvar} />
    <BotaoGenerico color="red" enable={enableLimpar} text="🗑️ Limpar grade" on:click={limpar} />
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
        gap: 5%;

        box-sizing: border-box;
        width: 100%;
        height: 100%;
    }
</style>