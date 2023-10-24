<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import BotaoGenerico from "./BotaoGenerico.svelte";
	import JanelaSalvar from "./JanelaSalvar.svelte";

    let dispatch = createEventDispatcher();

    const salvar = () => {dispatch("salvar")}
    const limpar = () => {dispatch("limpar")}
</script>

<JanelaSalvar />
<div class="grupo-botoes">
    <BotaoGenerico text="Salvar grade" on:click={salvar} />
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