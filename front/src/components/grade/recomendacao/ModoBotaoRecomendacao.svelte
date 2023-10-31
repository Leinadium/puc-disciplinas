<script lang="ts">
	import type { ModoRecomendacao } from "$lib/recomendacao";

    export let value: ModoRecomendacao = "todas";
</script>

<div id="modo-botao-recomendacao">
    <select bind:value={value}>
        <option value="eletivas">Eletivas</option>
        <option value="obrigatorias">Obrigatórias</option>
        <option value="todas">Todas</option>
    </select>
</div>

