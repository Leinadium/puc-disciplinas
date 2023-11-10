<script lang="ts">
    import est0 from "$lib/images/estrela-0.png";
    import est1 from "$lib/images/estrela-1.png";
    import est2 from "$lib/images/estrela-2.png";
    import est3 from "$lib/images/estrela-3.png";
    import est4 from "$lib/images/estrela-4.png";
    import est5 from "$lib/images/estrela-5.png";
    import est6 from "$lib/images/estrela-6.png";
    import est7 from "$lib/images/estrela-7.png";
    import est8 from "$lib/images/estrela-8.png";
    import est9 from "$lib/images/estrela-9.png";
    import est10 from "$lib/images/estrela-10.png";

    import { createEventDispatcher } from "svelte";

    const estrelasImg = [
        est1, est2, est3, est4, est5,
        est6, est7, est8, est9, est10
    ];

    export let nota: number = 0;
    export let enable: boolean = false;

    let estrelas: string[] = [];
    $: {
        estrelas = [est0, est0, est0, est0, est0];
        if (nota > 0) {
            estrelas = [];
            for (let i = 1; i <= 5; i++) {
                if (nota >= i ) {
                    estrelas.push(estrelasImg[9]);
                } else if (i > nota && i < nota + 1) {
                    estrelas.push(estrelasImg[Math.floor((i - nota) * 10) - 1]);    
                } else if (enable) {
                    estrelas.push(est0);
                }
                
            }
        }
    }

    let dispatch = createEventDispatcher<{nota: number}>();
    function click(x: number) {
        if (enable) {
            dispatch("nota", x);
        }
    }

</script>

<div class="estrelas">
    {#each estrelas as estrela, i}
        {#if enable}
            <a href="/#" on:click|preventDefault={() => click(i + 1)}>
                <img src={estrela} alt="Estrela {i+1}"/>
            </a>
        {:else}
            <img src={estrela} alt="Estrela {i+1}"/>
        {/if}
    {/each}
</div>


<style>
    .estrelas {
        height: 100%;
        aspect-ratio: 5/1;

        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
    }

    a {
        height: 100%;
    }

    img {
        height: 100%;
    }

</style>