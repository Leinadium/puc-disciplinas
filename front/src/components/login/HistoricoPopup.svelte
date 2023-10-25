<script lang="ts">
	import { armazenarHistorico } from "$lib/api";
	import { createEventDispatcher, onMount } from "svelte";
	import type { LoadingStatus } from "../../types/api";
	import type { PostHistorico } from "../../types/data";
	import Popup from "../common/Popup.svelte";

    const SAU_URL = "https://sau.puc-rio.br/WebLoginPucOnline/Default.aspx?sessao=VmluY3Vsbz1BJlNpc3RMb2dpbj1QVUNPTkxJTkVfQUxVTk8mQXBwTG9naW49TE9HSU4mRnVuY0xvZ2luPUxPR0lOJlNpc3RNZW51PVBVQ09OTElORV9BTFVOTyZBcHBNZW51PU1FTlUmRnVuY01lbnU9TUVOVQ__";

    let status: LoadingStatus | null = "loading";
    let statusMessage: string = "Carregando...";

    async function submit() {
        status = "loading";
        statusMessage = "Carregando...";

        const fileElement = document.getElementById("file-input") as HTMLInputElement; 
        if (fileElement.files == null) return;
        const file = fileElement.files[0];

        try {
            let res: PostHistorico = await armazenarHistorico(file);
            statusMessage = 'Histórico armazenado com sucesso!';
            if (res.curriculo) {
                statusMessage += ` Currículo: ${res.curriculo}`;
            }
            if (res.inseridos) {
                statusMessage += ` Inseridos: ${res.inseridos}`;
            }

            status = "success";
        } catch (e: any) {
            statusMessage = e.message;
            status = "error";
        }
    }

    let dispatch = createEventDispatcher();

    const close = () => { dispatch('close'); }

    onMount(() => {
        status = "success";
        statusMessage = "";
    })
</script>

<Popup on:close>
    <form id="historico-popup" on:submit|preventDefault={submit}>
        <div id="historico-textos">
            <p>
                Envie seu histórico escolar.
            </p>
            <p class="small">
                Acesse o <a href={SAU_URL}>SAU</a> e clique em "Histórico Escolar".<br>
                Clique com o botão direito em qualquer lugar da página, selecione "Salvar como..." e salve o arquivo em seu computador.<br>
                Envie o arquivo .htm ou .html gerado. 
            </p>
        </div>

        <input id="file-input" type="file" name="historico">
        
        <div id="historico-submit">
            <div class="status {status}">
                {statusMessage}
            </div>
            <button type="submit" on:click|preventDefault={submit}>
                Enviar
            </button>
        </div>

        <a id="historico-close" href="/#" on:click|preventDefault={close}>Voltar</a>
    </form>
</Popup>

<style>
    #historico-popup {
        box-sizing: border-box;
        padding: 50px;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;

        background: #eee;
        border: 5px solid #aaa;
        border-radius: 30px;
        gap: 10px;
        color: #111;
    }

    #historico-textos {
        text-align: center;
        margin: 0;
        font-size: 1.1rem;
    }

    #historico-textos p.small {
        font-size: 0.8rem;
    }

    #file-input {
        width: 100%;
        padding: 10px;
        border-radius: 10px;
        border: 1px solid #aaa;
        background: #ccc;
    }

    #historico-submit {
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 10px;
        width: 100%;
    }

    button {
        box-sizing: border-box;
        padding: 1% 5%;
        border-radius: 10px;
        background: #fff;
        font-size: 0.9rem;
        font-weight: 100;
        color: #444;
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
        color: #aaa;
    }

    .success {
        color: #3c3;
    }

    #historico-close {
        font-size: 0.8rem;
        color: #444;
    }
</style>