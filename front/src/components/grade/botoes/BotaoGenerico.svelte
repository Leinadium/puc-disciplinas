<script lang="ts">
	import type { GenericColor } from "../../../types/style";

    export let text: string = "";
    export let color: GenericColor = "white";
</script>

<button type="button" class="{color}" on:click|preventDefault>
    {text}
</button>

<style>
    button {
        box-sizing: border-box;
        width: 80%;
        height: fit-content;
        padding: 3%;
        border-radius: 5px;

        font-size: 0.8em;
        cursor: pointer;
    }

    .green {
        background: #33ff33;
        border: 1px solid #33cc33;
    }
    .green:hover {
        background: #33cc33;
        border: 1px solid #339933;
    }

    .red {
        background: #ff3333;
        border: 1px solid #cc3333;
    }
    .red:hover {
        background: #cc3333;
        border: 1px solid #993333;
    }

    .white {
        background: #ffffff;
        border: 1px solid #cccccc;
        color: black;
    }
    .white:hover {
        background: #cccccc;
        border: 1px solid #999999;
    }
</style>