<script lang="ts">
	import Footer from "../components/Footer.svelte";
	import Header from "../components/Header.svelte";
</script>

<Header titulo="Macro" />
    <main>
        <slot></slot>
    </main>
<Footer />

<style>
    main {
        width: 100%;
        height: 100%;
    }
</style>