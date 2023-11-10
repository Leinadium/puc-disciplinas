<script lang="ts">
	import CheckLogin from "../components/CheckLogin.svelte";
    import Footer from "../components/Footer.svelte";
	import Header from "../components/Header.svelte";
</script>

<svelte:head>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
</svelte:head>

<Header titulo="Macro" />
    <CheckLogin />
    <main>
        <slot></slot>
    </main>
<Footer />

<style>
    main {
        width: 100%;
        min-height: 100%;

        background: var(--color-main-1);
    }

    :global(html) {
        --color-main-1: #04364A;
        --color-main-2: #176B87;
        --color-main-3: #64CCC5;
        --color-main-4: #DAFFFB;

        --color-mono-1: #1F4754;
        --color-mono-2: #87D4ED;
        --color-mono-3: #C4F0FF;
        --color-mono-4: #D4F4FF;

        --color-anal-1: #17877F;
        --color-anal-2: #174887;
        --color-anal-3: #178759;
        --color-anal-4: #172487;

        --color-tria-1: #8C8111;
        --color-tria-2: #8C1140;
        --color-tria-3: #193237;
        --color-tria-4: #FFF69E;

        --color-quad-1: #528717;
        --color-quad-2: #875017;
        --color-quad-3: #771787;
        --color-quad-4: #6FC1DC;

        --color-main-4f: #B8DDD9;
        --color-whitef: #ccc;

        --color-blackf: #777;

        --border-radius: 15px;

        /* font-family: Verdana, Geneva, Tahoma, sans-serif; */
        font-family: Lato, Verdana, Geneva, Tahoma, sans-serif;
    }
</style>