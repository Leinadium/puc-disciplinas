<script lang="ts">
	import LoginHeader from "./login/LoginHeader.svelte";

    export let titulo: String = "Titulo";
</script>

<header>
    <nav>
        <a class="nav-item" href="/grade">Grade</a>
        <a class="nav-item" href="/avaliacao">Avaliação</a>
    </nav>

    <h1>
        <a class="nav-item" href="/">{titulo}</a>
    </h1>

    <div id="nav-login"><LoginHeader /></div>
</header>

<style>
    header {
        background-color: var(--color-main-2);
        color: var(--color-main-4);
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
        color: var(--color-main-4);
        box-sizing: border-box;
        padding: 2%;
        border-radius: var(--border-radius);
        background: transparent;
        font-size: 1rem;

        transition: background 0.2s;
    }
    .nav-item:hover {
        background: var(--color-main-3);
    }

    h1 {
        margin: 0;
        text-align: center;
    }
    h1 > a {
        font-size: 2rem !important;
    }

    #nav-login {
        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-end;
        align-items: center;
        gap: 10%;
    }
</style>