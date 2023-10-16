<script lang="ts">
	import { fazerLogin } from "$lib/api";
    import { userStore } from "$lib/stores";

    let usuario: string = "";
    let senha: string = "";

    const login = async () => {
        try {
            let nome: string = await fazerLogin(usuario, senha);
            userStore.set(nome);
        } catch (error) {
            console.log(error);
        }
    }
</script>

<div id="login-teste">
    <input type="text" bind:value={usuario} />
    <input type="password" bind:value={senha} />
    <button on:click={login}>Login</button>
</div>