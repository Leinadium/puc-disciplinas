<script lang="ts">
    import type { GenericBackground, GenericColor } from "../../types/style";
    export let color: GenericColor = "default";
    export let background: GenericBackground = "fill";
    export let clickCallback: () => void = () => {};
</script>


<a 
    href="/#" 
    class="generic-box {color} {background}"
    on:click|preventDefault={clickCallback}
>
    <slot></slot>
</a>

<style>
    a {
        text-decoration: none;
        color: #000;
    }

    .generic-box {
        position: relative;
        box-sizing: border-box;
        /* width: min(100px, 10vw); */
        /* height: var(--height); */
        width: 100%;
        height: 100%;

        border-radius: var(--border-radius);
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        
        padding: 10px;
        /* margin: 10px; */

        text-align: center;

        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: center;
        gap: 0.2rem;
    }

    .outline {
        background-color: transparent !important;
        border: 2px solid #000 !important;
    }

    .default {
        background-color: var(--color-quad-4);
    }

    .yellow {
        background-color: var(--color-tria-1);
    }

    .white {
        background-color: var(--color-tria-4);
    }

    .green {
        background-color: var(--color-anal-3);
    }

    .blue {
        background-color: var(--color-quad-4);
    }
</style>