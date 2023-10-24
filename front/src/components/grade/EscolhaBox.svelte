<script lang="ts">
	import { fade } from "svelte/transition";
	import type { EscolhaSimples, RemoveDisciplinaEvent } from "../../types/data";
	import type { UIEscolha } from "../../types/ui";
    import GenericBox from "../common/GenericBox.svelte";
    import { createEventDispatcher } from "svelte";

    export let info: UIEscolha = {
        codigo: "INF9999",
        nome: "Técnicas de Programação II",
        turma: "3WA",
        professor: "Fulano de Tal"
    }

    let showBtn: boolean = false;

    let dispatchPopup = createEventDispatcher<{popup: string}>();
    let dispatchRemove = createEventDispatcher<{remove: RemoveDisciplinaEvent}>();

    const popup = () => {
        dispatchPopup("popup", info.codigo);
    }

    const remove = () => {
        dispatchRemove("remove", {disciplina: info.codigo});
    }

    const toggleBtn = () => {
        showBtn = !showBtn;
    }

</script>

<GenericBox color="yellow" clickCallback={toggleBtn}>
    <div class="wrapper">
        <div class="upper">
            <span id="codigo">{info.codigo.toUpperCase()}</span>
            <span id="turma">{info.turma.toUpperCase()}</span>
        </div>
        
        <div class="lower">
            <span id="nome">{info.nome}</span>
            <span id="professor">{info.professor}</span>
        </div>
        
        {#if showBtn }
            <div id="wrapper-btn" transition:fade={{duration: 200}}>
                <a class="btn" id="remover" href="/#" on:click|preventDefault={remove}>x</a>
                <a class="btn" id="informacao" href="/#" on:click|preventDefault={popup}>i</a>
            </div>
        {/if}
    </div>
</GenericBox>

<style>
    .wrapper {
        width: 100%;
        height: 100%;
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 0.2rem;
        
        position: relative;
    }

    span {
        /* width: 90%; */
        overflow: hidden;
        /* white-space: nowrap; */
        text-overflow: ellipsis;
    }

    .upper {
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;

        font-weight: bold;
    }

    .lower {
        width: 100%;
        display: flex;
        flex-flow: column nowrap;
        justify-content: space-around;
        align-items: center;
        gap: 0.2rem;
    }

    #codigo {
        font-size: 0.75rem;
    }

    #turma {
        font-size: 0.65rem;
        font-style: italic;
    }

    #nome {
        width: 100%;
        font-size: 0.75rem;
    }

    #professor {
        font-size: 0.65rem;
        font-style: italic;
    }

    #wrapper-btn {
        position: absolute;
        top: 0;
        right: 0;
        width: 100%;
        height: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;

        z-index: 5;
    }

    .btn {
        width: 30%;
        aspect-ratio: 1/1;
        font-size: 1.5rem;
        text-align: center;
        padding-top: 0.3rem;
        border-radius: 50%;
        text-decoration: none;
    }

    #remover {
        background-color: #a00;
        color: #fff;
    }

    #informacao {
        background-color: #00a;
        color: #fff;
    }
</style>