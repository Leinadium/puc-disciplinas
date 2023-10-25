<script lang="ts">
	import { fazerLogout } from "$lib/api";
	import { userStore } from "$lib/stores";
	import LoginPopup from "./LoginPopup.svelte";

    // popup de login
    let toggleLogin: boolean = false;
    const clickLogin = () => {
        toggleLogin = true;
    }
    const closeLogin = () => {
        toggleLogin = false;
    }

    // logout
    const logout = async () => {
        if (await fazerLogout()) {
            userStore.set(null);
        }
    }

</script>

{#if $userStore != null}
    <span>Olá, {$userStore}</span>
    <a class="log" href="/#" on:click|preventDefault={logout}>
        Sair
    </a>
{:else}
    <a class="log" href="/#" on:click|preventDefault={clickLogin}>
        Entrar com SAU
    </a>
{/if}

{#if toggleLogin}
    <LoginPopup on:close={closeLogin}/>
{/if}

<style>
    span {
        font-size: 1.1rem;
        font-weight: bold;
    }

    .log {
        text-decoration: none;
        color: #eee;
        box-sizing: border-box;
        padding: 2%;
        border-radius: 0.5rem;
        background: transparent;

        transition: background 0.2s;
    }
    .log:hover {
        background: #666;
    }
</style>