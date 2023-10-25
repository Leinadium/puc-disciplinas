<script lang="ts">
	import CampoPesquisa from "../../components/avaliacao/CampoPesquisa.svelte";


</script>

<div id="avaliacao-page">
    <div id="avaliacao-container">
        <CampoPesquisa />
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