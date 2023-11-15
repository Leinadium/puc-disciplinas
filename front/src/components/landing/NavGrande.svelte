<script lang="ts">
    export let color: number = 1 | 2;
    export let link: string = "/#";

    let click = () => window.location.href = link;
</script>

<button class="nav-grande color-{color}" on:click|preventDefault={click}>
    <slot></slot>
</button>

<style>
    .nav-grande {
        border: none;
        cursor: pointer;
        text-decoration: none;

        box-sizing: border-box;
        border-radius: var(--border-radius);
        background: var(--color-main-2);
        color: var(--color-whitef);
        box-shadow: 0px 0px 15px 0px var(--color-black3);
        
        display: flex;
        flex-flow: column;
        justify-content: center;
        align-items: center;

        width: 20%;

        aspect-ratio: 1 / 1;
        transition: all 0.2s ease-in-out;
    }

    .nav-grande:hover {
        transform: scale(1.05);
    }

    .color-1 {
        background: var(--color-anal-3);
    }

    .color-2 {
        background: var(--color-anal-4);
    }
</style>