<script lang="ts">
    import { calcularDiferenca } from "$lib/utils";
	import type { Modificacao } from "../../../types/data";
	import Legenda from './Legenda.svelte';

    export let modificacao: Modificacao | null = null;
    $: console.log(modificacao);

    let diferenca: string = "";
    $: diferenca = modificacao?.dataGeral ? calcularDiferenca(new Date(modificacao.dataGeral)) + " atrás" : "Não disponível";

    

</script>


<div id="textos-descricao">
    <span class="descricao-grande">
        Monte sua grade horária
    </span>
    
    <p class="descricao">
        Crie sua grade horária para o seu próximo semestre. <b>Macro</b> disponibiliza
        todas as informações necessárias para você montar sua grade
        de forma rápida e fácil. Além disso, você pode salvar sua grade e acessá-la a qualquer momento.
        <br><br>
        Entre com sua conta do SAU e cadastre seu histórico escolar para obter recomendações
        personalizadas e filtros de acordo com o seu currículo.
    </p>

    <Legenda />

    <div class="modificacao">
        <span>Última atualização: {diferenca}</span>
        {#if modificacao?.modoFallback}
            <span class="modo-fallback">
                ⚠️
                O microhorário está em manutenção. 
                A quantidade de vagas e créditos pode estar indisponível.
                ⚠️
            </span>
        {/if}
    </div>

</div>

<style>
    #textos-descricao {
        width: 50%;
        text-align: center;
        padding-top: 1%;

        color: var(--color-main-4f);

        display: flex;
        flex-flow: column;
        justify-content: center;
        gap: 1rem;
    }

    .descricao-grande {
        font-size: 2.1rem;
        font-weight: bold;
    }

    .descricao {
        font-size: 1.1rem;
        font-weight: 300;
    }

    .modificacao {
        font-size: 0.9rem !important;
        
        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-end;
        align-items: center;
        gap: 0.1rem;
    }

    .modo-fallback {
        color: var(--color-tria-1);
        font-weight: bold;
        font-size: 0.9rem;
    }
</style>