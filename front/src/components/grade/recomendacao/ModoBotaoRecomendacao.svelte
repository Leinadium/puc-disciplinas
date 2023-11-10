<script lang="ts">
	import type { ModoRecomendacao } from "$lib/recomendacao";

    export let value: ModoRecomendacao = "todas";
</script>

<select bind:value={value}>
    <option value="eletivas">Eletivas</option>
    <option value="curriculares">Curriculares</option>
    <option value="todas">Todas</option>
</select>


<style>
    select {
        background: var(--color-main-4);
        border-radius: var(--border-radius);

        font-size: 0.9rem;
        font-weight: bold;
    }    
</style>

