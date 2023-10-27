<script lang="ts">
	import CheckLogin from "../components/CheckLogin.svelte";
    import Footer from "../components/Footer.svelte";
	import Header from "../components/Header.svelte";
</script>

<Header titulo="Macro" />
    <CheckLogin />
    <main>
        <slot></slot>
    </main>
<Footer />

<style>
    main {
        width: 100%;
        height: 100%;
    }

    :global(html) {
        --color-main-1: #04364A;
        --color-main-2: #176B87;
        --color-main-3: #64CCC5;
        --color-main-4: #DAFFFB;
        --color-tria-1: #8C8111;
        --color-tria-2: #8C1140;
        --color-tria-3: #11768C;
        --color-tria-4: #193237;
        --color-tria-5: #FFF69E;

        --border-radius: 15px;
    }
</style>