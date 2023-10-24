<script lang="ts">
    import type { GenericColor } from "../../types/style";
    export let color: GenericColor = "default";
    export let clickCallback: () => void = () => {};

    // export let height: string = "100px";
    // export let width: string = "100px";

    // $: customStyle = "--height: " + height + "; --width: " + width + ";"
</script>


<a href="/#" class="generic-box {color}" on:click|preventDefault={clickCallback}>
    <slot></slot>
</a>

<style>
    a {
        text-decoration: none;
        color: #000;
    }

    .generic-box {
        box-sizing: border-box;
        /* width: min(100px, 10vw); */
        /* height: var(--height); */
        width: 100%;
        height: 100%;

        border-radius: 10px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        
        padding: 10px;
        /* margin: 10px; */

        text-align: center;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
        align-items: center;
    }

    .default {
        background-color: #aaa;
    }

    .yellow {
        background-color: #aa3;
    }

    .green {
        background-color: #5f5;
    }
</style>