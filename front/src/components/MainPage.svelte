<script lang="ts">
    import Header from "./Header.svelte";
    import Footer from "./Footer.svelte";
</script>

<Header titulo="Macro" />
    <main>
        <slot />
    </main>
<Footer />
