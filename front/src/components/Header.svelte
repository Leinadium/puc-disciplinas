<script lang="ts">
	import LoginTeste from "./login/LoginTeste.svelte";

    export let titulo: String = "Titulo";
</script>

<header>
    <h1>{titulo}</h1>
    <LoginTeste />
</header>

<style>
    header {
        background-color: #333;
        color: #fff;
        box-sizing: border-box;
        padding: 1rem;
        text-align: center;
    }

    h1 {
        margin: 0;
    }
</style>