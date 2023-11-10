<script lang="ts">
    import legenda from "$lib/images/legenda.png";
    import { scale, slide } from "svelte/transition";

    let show: boolean = false;
    function click() {
        show = !show;
    }

    let icon: string;
    $: icon = !show ? "ℹ️" : "❌" ;
    
    let color: string;
    $: color = !show ? 'blue' : 'red';

</script>

<a class={color} href="/#" on:click|preventDefault={click}>
    {icon}
    {!show ? "Ver" : "Ocultar"}
    legenda
</a>

{#if show}
    <img src={legenda} alt="Legenda" transition:slide={{axis: 'y'}}/>
{/if}


<style>
    a {
        text-decoration: none;
        padding-bottom: 1rem;
    }

    img {
        width: 100%;
        margin: 1rem;
    }

    .blue {
        color: var(--color-mono-2);
    }

    .red {
        color: var(--color-tria-2);
    }
</style>