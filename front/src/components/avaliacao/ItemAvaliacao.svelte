<script lang="ts">
	import type { ItemDisciplina, ItemGenerico, SelectAvaliacaoEvent } from "../../types/data";
    import { createEventDispatcher } from "svelte";

    export let info: ItemGenerico;

    // parsing do nome
    let texto: string;
    let identificador: string;
    if (info.tipo == "disciplina") {
        const x = (info.conteudo as ItemDisciplina)
        texto = `${x.codigo} - ${x.nome}`;
        identificador = x.codigo;
    } else if (info.tipo == "professor") {
        texto = (info.conteudo as ItemDisciplina).nome;
        identificador = texto;
    } else {
        texto = "";
    }

    let dispatch = createEventDispatcher<{click: SelectAvaliacaoEvent}>();

    function click() {
        dispatch("click", {
            tipo: info.tipo,
            id: identificador
        });
    }

</script>

<div class="item-avaliacao">
    <a href="/#" class="texto" on:click|preventDefault={click}>
        {texto}
    </a>
</div>

<style>

</style>