<script lang="ts">
	import { checkLogin, fazerLogin } from "$lib/api";
	import { userStore } from "$lib/stores";
	import { createEventDispatcher, onMount } from "svelte";
    import Popup from "../common/Popup.svelte";
	import type { LoadingStatus } from "../../types/api";

    let dispatch = createEventDispatcher();

    let usuario: string = "";
    let senha: string = "";

    let status: LoadingStatus | null = "loading";
    let statusMessage: string = "Carregando...";

    async function login() {
        try {
            status = "loading";
            statusMessage = "Carregando...";
            
            let nome: string = await fazerLogin(usuario, senha);
            userStore.set(nome);
            await checkLogin();

            status = "success";
            dispatch("close");
        }
        catch (e: any) {
            status = "error";
            statusMessage = e.message;
        }
    }

    onMount(() => {
        status = "success"
        statusMessage = "";
    })
</script>

<Popup on:close>
    <form id="login-popup" on:submit|preventDefault={login}>
        <div id="login-textos">
            <p>
                Faça seu login utilizando seu usuário e senha do SAU da PUC-Rio.   
            </p>
            <p class="small">
                Suas credenciais NÃO são armazenadas no sistema,<br>
                mas são utilizadas para validar sua matrícula.
            </p>
        </div>
        
        <div id="login-inputs">
            <div class="login-field">
                <label for="login-usuario">Usuário</label>
                <input type="text" bind:value={usuario}>
            </div>
            <div class="login-field">
                <label for="login-senha">Senha</label>
                <input type="password" bind:value={senha}/>
            </div>
            
        </div>
        <div id="login-submit">
            <div class="status {status}">
                {statusMessage}
            </div>
            
            <button type="submit" on:click|preventDefault={login}>
                Entrar
            </button>
        </div>
        
    </form>
</Popup>

<style>
    #login-popup {
        box-sizing: border-box;
        padding: 50px;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 30px;

        background: var(--color-main-1);
        border-radius: var(--border-radius);

        color: var(--color-whitef);
    }

    #login-textos {
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 10px;
    }
    p {
        text-align: center;
        margin: 0;
        font-size: 1.1rem;
        
    }
    p.small {
        font-size: 0.8rem;
    }

    #login-inputs {
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 10px;

    }

    .login-field {
        width: 100%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        gap: 10px;
    }

    label {
        width: 30%;
    }
    input {
        border: none;
        width: 70%;
        background: var(--color-mono-3);
    }

    #login-submit {
        box-sizing: border-box;
        width: 100%;
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
    }

    button {
        box-sizing: border-box;
        padding: 1% 5%;
        border-radius: 10px;
        border: none;
        background: var(--color-anal-3);
        font-size: 0.9rem;
        font-weight: 100;
        color: var(--color-whitef);
    }

    .status {
        font-size: 1rem;
        font-weight: bold;
        font-style: italic;
        height: 1.5rem;
    }

    .error {
        color: #f33;
    }

    .loading {
        color: var(--color-whiteff);
    }

    .success {
        color: transparent;
    }

</style>