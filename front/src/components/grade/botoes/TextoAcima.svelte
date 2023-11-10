<script lang="ts">
	import { pluralize } from "$lib/utils";

    export let creditos: number = 0;
    export let disciplinas: number = 0;
</script>

<div id="texto-acima">
    <span>{disciplinas} {pluralize(disciplinas, "disciplina")},</span>
    <span>{creditos} {pluralize(creditos, "credito")}</span>
</div>

<style>
    #texto-acima {
        width: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;

        margin-bottom: 10%;
    }

    span {
        font-size: 0.8em;
        font-weight: bold;
        color: var(--color-whitef);
    }
</style>