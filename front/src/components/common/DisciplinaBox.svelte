<script lang="ts">
	import type { DisciplinaBasicaType } from "../../types/disciplinas.type";

    export let info: DisciplinaBasicaType = {
        codigo: "INF9999",
        nome: "Técnicas de Programação II",
        creditos: 6,
        qtdTurmas: 3,
        qtdVagas: 120
    };

</script>

<div class="disciplina-box">
    <div class="textos">
        <span id="codigo">{info.codigo.toUpperCase()}</span>
        <span id="nome">{info.nome.toUpperCase()}</span>
    </div>
    
    <div class="numeros">
        <span id="creditos">{info.creditos} <span class="small">créditos</span></span>
    </div>
    

    <div class="numeros">
        <span id="turmas">{info.qtdTurmas} <span class="small">turmas</span></span>
        <span id="vagas">{info.qtdVagas} <span class="small">vagas</span></span>
    </div>
</div>

<style>
    .disciplina-box {
        box-sizing: border-box;
        width: max(100px, 10vw);
        height: 100px;

        background-color: #aaa;
        border-radius: 10px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        
        padding: 10px;
        margin: 10px;

        text-align: center;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
    }

    .textos {
        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
    }

    .small {
        font-size: 0.6rem;
    }

    #nome {
        font-size: 0.8rem;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }

    .numeros {
        font-size: 0.8rem;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
    }
</style>