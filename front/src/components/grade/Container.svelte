<script lang="ts">
	import Grade from "./Grade.svelte";
	import ListaSelecao from "./ListaSelecao.svelte";


</script>

<div id="grade-container">
    <Grade />
    <ListaSelecao />
</div>


<style>
    #grade-container {
        width: min(80vw, 1000px);
        height: 80%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: stretch;
        gap: 1rem;
    }
</style>