<script lang="ts">
	import EscolhaBox from "./EscolhaBox.svelte";
    import { formatHora, generateGrade, removeExtraFromEscolha } from "$lib/grade";

    import type { UIDiaDisciplina, UIHoraDisciplina, UIEscolha } from "../../types/ui";
    import type { EscolhaInfoExtra } from "../../types/data";

    export let info: EscolhaInfoExtra[] = [
        {
            codigo: "INF1350",
            nome: "Programação Reativa",
            turma: "3WA",
            professor: "Adriano de Souza",
            dia: "SEG",
            inicio: 9,
            horas: 2,
        }, 
        {
            codigo: "INF1350",
            nome: "Programação Reativa",
            turma: "3WA",
            professor: "Adriano de Souza",
            dia: "QUA",
            inicio: 9,
            horas: 2,
        }, 
        {
            codigo: "ENG4403",
            nome: "Engenharia de Dados",
            turma: "3WA",
            professor: "Marcos Villas",
            dia: "QUA",
            inicio: 7,
            horas: 2,
        }, 
    ];


    const diasColunas: UIDiaDisciplina[] = ["SEG", "TER", "QUA", "QUI", "SEX"];
    const horasLinhas: UIHoraDisciplina[] = [7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20];

    // grade horaria a partir das disciplinas
    let g: any = {}
    $: g = generateGrade(diasColunas, horasLinhas, info);

</script>


<div id="grade">
    <table>
        <thead>
            <tr>
                <th></th>
                {#each diasColunas as dia}
                    <th>{dia.toUpperCase()}</th>
                {/each}
            </tr>
        </thead>
        <tbody>
            {#each horasLinhas as h}
                <tr>
                    <td>{formatHora(h)}</td>
                    {#each diasColunas as d}
                        {#if !!g[h][d] }
                            <td rowspan={g[h][d]['horas']}>
                                <EscolhaBox info={removeExtraFromEscolha(g[h][d])} />
                            </td>
                        {:else if g[h][d] === null}
                            <td></td>
                        {/if}
                    {/each}
                </tr>
            {/each}

        </tbody>
    </table>
</div>

<style>
    #grade {
        box-sizing: border-box;
        /* posicionamento da grade no container */
        grid-column: 1 / span 1;
        grid-row: 2 / span 1;

        overflow-y: scroll;
        
        border: 1px solid #ccc;
        
        /* para compensar a barra de scroll*/
        padding-right: 10px;
    }
	table {
        box-sizing: border-box;
        table-layout: fixed;
        border-collapse: collapse;
		height: 100%;
        width: 100%;

        border: 3px solid purple;
	}

    td, th {
        border: 1px solid black;
        padding: 5px;
        margin: 0;
        width: 100%;
    }

    td {
        height: 100%;
    }

    tr {
        height: 60px;
    }

    /* primeiro td (horas) */
    td:first-child, th:first-child {
        /* width: 10% !important; */
        width: 40px;
        padding: 0;
        font-size: 0.8em;
        white-space: nowrap;
        text-align: center;
    }

    thead > tr {
        height: 20px;
        padding: 0;
    }
</style>