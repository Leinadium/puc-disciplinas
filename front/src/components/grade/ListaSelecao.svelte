<script lang="ts">

</script>


<div id="lista-selecao">
    <div id="search-bar">
        <input type="text" placeholder="Pesquisar disciplina" />
    </div>

    <div id="resultados">
        <span>Resultados</span>
        <span>Resultados</span>
        <span>Resultados</span>
        <span>Resultados</span>
        <span>Resultados</span>
        <span>Resultados</span>
        <span>Resultados</span>
    </div>
</div>


<style>
    #lista-selecao {
        width: 10%;
        height: 100%;

        flex-grow: 1;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 1rem;
    }

    #resultados {
        width: 100%;
        height: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 1rem;
    }

    
</style>