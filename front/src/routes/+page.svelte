<script lang="ts">
	import NavGrande from "../components/landing/NavGrande.svelte";
	import TextoAbaixo from "../components/landing/TextoAbaixo.svelte";

</script>

<div id="landing-page">
    <h1>Planeje seu semestre de maneira inteligente</h1>

    <p class="descricao">
        <b>Macro</b> é um serviço que permite planejar seu próximo semestre
        de maneira rápida e inteligente. Com ele, você pode montar sua grade
        de horários, pesquisar e avaliar professores e disciplinas, e muito mais!
    </p>

    <div id="opcoes">
        <NavGrande color={1}>
            <h3>Grade</h3>
            <p class="menor">Monte sua grade se utilizando de diversas informações e recomendações</p>
        </NavGrande>
        <NavGrande color={2}>
            <h3>Avaliação</h3>
            <p class="menor">Pesquise e avalie professores e suas disciplinas já cursadas</p>
        </NavGrande>
    </div>

    <TextoAbaixo />

</div>

<style>
    #landing-page {
        width: 100%;
        height: 100%;

        color: var(--color-whitef);

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
    }

    h1 {
        margin: 0;
        margin-top: 5%;
        padding: 0;
    }

    .descricao {
        width: 60%;
        font-size: 1rem;
        font-style: italic;
        margin-bottom: 3%;
        text-align: center;
    }

    #opcoes {
        width: 100%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        gap: 10%;
    }

    h3 {
        font-size: 2rem;
    }

    .menor {
        width: 80%;
        font-size: 1.0rem;
        font-style: italic;
        padding-bottom: 15%;
    }
</style>