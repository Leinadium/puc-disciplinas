<script lang="ts">
	import BotaoGenerico from "./BotaoGenerico.svelte";


</script>

<div class="grupo-botoes">
    <BotaoGenerico />
    <BotaoGenerico />
</div>

<style>
    .grupo-botoes {
        /* posicionamento da grade no container */
        grid-column: 2 / span 1;
        grid-row: 1 / span 1;
    }
</style>