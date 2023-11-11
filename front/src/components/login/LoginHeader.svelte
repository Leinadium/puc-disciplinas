<script lang="ts">
	import { checkLogin, fazerLogout } from "$lib/api";
	import { userStore } from "$lib/stores";
	import HistoricoPopup from "./HistoricoPopup.svelte";
	import LoginPopup from "./LoginPopup.svelte";

    // popup de login
    let toggleLogin: boolean = false;
    const clickLogin = () => { toggleLogin = true; }
    const closeLogin = () => { toggleLogin = false; }

    // popup de histórico
    let toggleHistorico: boolean = false;
    const clickHistorico = () => { toggleHistorico = true; }
    const closeHistorico = () => { toggleHistorico = false; }

    // logout
    const logout = async () => {
        if (await fazerLogout()) {
            userStore.set(null);
            checkLogin();
        }
    }

</script>

{#if $userStore != null}

    <span>Olá, {$userStore}</span>

    <a class="log" href="/#" on:click|preventDefault={clickHistorico}>
        Histórico
    </a>    

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
{#if toggleHistorico}
    <HistoricoPopup on:close={closeHistorico}/>
{/if}

<style>
    span {
        font-size: 1.1rem;
        font-weight: bold;
    }

    .log {
        text-decoration: none;
        color: var(--color-main-4);
        box-sizing: border-box;
        padding: 2%;
        border-radius: var(--border-radius);
        background: transparent;

        transition: background 0.2s;
    }
    .log:hover {
        background: var(--color-main-3);
    }
</style>