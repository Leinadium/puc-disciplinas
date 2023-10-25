<script lang="ts">
	import { userStore } from "$lib/stores";
	import LoginHeader from "./login/LoginHeader.svelte";
	import LoginPopup from "./login/LoginPopup.svelte";

    export let titulo: String = "Titulo";
</script>

<header>
    <nav>
        <a class="nav-item" href="/">Home</a>
        <a class="nav-item" href="/grade">Grade</a>
        <a class="nav-item" href="/avaliacao">Avaliação</a>
    </nav>

    <h1>{titulo}</h1>

    <div id="nav-login"><LoginHeader /></div>
</header>

<style>
    header {
        background-color: #333;
        color: #fff;
        box-sizing: border-box;
        padding: 1rem 10%;
        width: 100%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }

    nav, h1, #nav-login {
        flex: 1;
    }

    nav {
        box-sizing: border-box;
        
        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-start;
        align-items: center;
        gap: 10%;
    }

    .nav-item {
        text-decoration: none;
        color: #eee;
        box-sizing: border-box;
        padding: 2%;
        border-radius: 0.5rem;
        background: transparent;

        transition: background 0.2s;
    }
    .nav-item:hover {
        background: #666;
    }

    h1 {
        margin: 0;
        text-align: center;
    }

    #nav-login {
        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-end;
        align-items: center;
    }
</style>