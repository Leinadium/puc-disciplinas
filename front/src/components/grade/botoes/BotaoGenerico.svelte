<script lang="ts">
	import type { GenericColor } from "../../../types/style";

    export let enable: boolean = true;
    export let text: string = "";
    export let color: GenericColor = "white";
</script>

<button
    type="button" 
    class="{color}"
    on:click|preventDefault
    disabled={!enable}
>
    {text}
</button>

<style>
    button {
        box-sizing: border-box;
        width: 80%;
        height: fit-content;
        padding: 3%;
        border-radius: var(--border-radius);

        font-size: 0.75em;
        font-weight: bold;
        cursor: pointer;
        border: none;
        color: var(--color-whitef);
    }

    button:disabled {
        cursor: not-allowed;
        background: transparent !important;
    }


    .blue {
        background: var(--color-anal-4);
        border: 3px solid var(--color-anal-4);
        
    }

    .red {
        background: var(--color-tria-2);
        border: 3px solid var(--color-tria-2);
    }
</style>