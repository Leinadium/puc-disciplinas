<script lang="ts">
	import Popup from "../../common/Popup.svelte";
	import type { LoadingStatus } from "../../../types/api";
    
    export let status: LoadingStatus = "loading";
    export let link: string | null;
    
    let copyStatus: boolean = false;
    function copyLink() {
        if (link != null) {
            navigator.clipboard.writeText(link);
            copyStatus = true;
            setTimeout(() => {copyStatus = false}, 2000);
        }
    }
</script>

<Popup on:close>
    <div id="janela-salvar">
        {#if status == "loading"}
            <div id="loading">Carregando...</div>
        {:else if status == "success"}
            <div id="textos">
                <span id="texto-grande">Grade armazenada com sucesso!</span>
                <span id="texto-pequeno">Link permanente:</span>
            </div>
            
            <div id="div-link">
                <input type="text" id="link" value={link}>

                {#if copyStatus}
                    <span id="texto-minusculo">Copiado!</span>
                {:else}
                    <a href="/#" id="texto-minusculo" on:click|preventDefault={copyLink}>Copiar</a>
                {/if}
            </div>
        {:else}
            <div id="no-info">Erro ao armazenar grade. Tente novamente mais tarde.</div>
        {/if}
    </div>
</Popup>

<style>
    #janela-salvar {
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
    }

    #textos {
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
        gap: 30px;
    }

    #texto-grande {
        font-size: 1.5rem;
        font-weight: bold;
    }

    #texto-pequeno {
        font-size: 0.9rem;
        font-style: italic;
    }

    #texto-minusculo {
        font-size: 0.8rem;
        color: #444;
        text-decoration: none;

        box-sizing: border-box;
        padding: 5px;
        border: 1px solid #444;
        border-radius: 5px;
    }
    
    #div-link {
        display:flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        gap: 10px;
    }
    
    #link {
        box-sizing: border-box;
        border-radius: 10px;
        padding: 1% 5%;
        background: #fff;
        font-size: 0.9rem;
        font-weight: 100;
        color: #444;
        
        white-space: nowrap;
    }
</style>