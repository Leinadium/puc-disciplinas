<script lang="ts">
	import { pluralize } from "$lib/utils";
	import TurmaNota from "./TurmaNota.svelte";

    export let creditos: number;
    export let qtdTurmas: number;
    export let qtdAlunos: number;
    export let qtdAvaliacoes: number;
    export let avaliacaoMedia: number;
    export let grauMedio: number;

    // left pad
    function formatar(num: number): string {
        if (num < 0) {
            return '?';
        }
        else if (num < 10) {
            return '0' + num;
        }
        return num + '';
    }

    // quem diria, o grau é armazenado como base 100 no banco
    // mas na realidade era pra ser base 10...
    // fazendo essa alteração no front
    let grauCorrigido: number;
    $: grauCorrigido = grauMedio / 10;


</script>

<div id="info">
    <div id="info-top">
        <span class="info-numero">
            {formatar(creditos)}
            <span class="info-menor">{pluralize(creditos, "crédito")}</span>
        </span>
        <span class="info-numero">
            {formatar(qtdTurmas)}
            <span class="info-menor">{pluralize(qtdTurmas, "turma")}</span>
        </span>
    </div>
    <div id="info-bot">
        <div id="avaliacao">
            <TurmaNota
                descricao="Avaliação média"
                nota={avaliacaoMedia}
                base={5}
                textoPequeno="a partir de {qtdAvaliacoes} avaliaç{qtdAvaliacoes != 1 ? 'ões' : 'ão'}"
                isValid={qtdAvaliacoes > 0}
                isEstrela={true}
            />
        </div>
        <div id="grau">
            <TurmaNota
                descricao="Grau médio"
                nota={grauCorrigido}
                base={10}
                textoPequeno="a partir de {qtdAlunos} {pluralize(qtdAlunos, "aluno")}"
                isValid={qtdAlunos > 0}
                isEstrela={false}
            />
        </div>
    </div>
</div>

<style>
    div {
        box-sizing: border-box;
    }

    #info {
        width: 48%;
        height: 100%;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
        align-items: stretch;

        border-radius: var(--border-radius);
        background-color: var(--color-mono-1);
        color: var(--color-whitef);
    }

    #info-top {
        width: 100%;
        height: 30%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }

    .info-numero {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: flex-end;
        gap: 5%;

        font-size: 2.0rem;
        font-weight: bold;
    }

    .info-menor {
        font-size: 1.0rem !important;
        font-weight: 300;
        padding-bottom: 0.2rem;
    }

    #info-bot {
        width: 100%;
        height: 70%;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }
</style>