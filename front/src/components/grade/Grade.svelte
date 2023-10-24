<script lang="ts">
	import EscolhaBox from "./EscolhaBox.svelte";
    import { formatHora, generateGrade, getDuracao, removeExtraFromEscolha } from "$lib/grade";

    import type { UIDiaDisciplina, UIHoraDisciplina } from "../../types/ui";
    import type { EscolhaInfoExtra } from "../../types/data";

    export let info: EscolhaInfoExtra[] = [
        {
            codigo: "ENG4403",
            nome: "Engenharia de Dados",
            turma: "3WA",
            professor: "Marcos Villas",
            horarios: [
                {dia: "QUA", inicio: 7, fim: 9}
            ],
            shf: 0,
        }, 
    ];

    $: console.log(info);

    const diasColunas: UIDiaDisciplina[] = ["SEG", "TER", "QUA", "QUI", "SEX"];
    const horasLinhas: UIHoraDisciplina[] = [7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20];

    // copia das escolhas que é SHF ou a distância)
    let shf: EscolhaInfoExtra[] = [];
    
    function filtroShfDistancia(e: EscolhaInfoExtra) {
        return e.shf > 0 || e.horarios.filter(h => h.dia === "xxx").length > 0
    };

    $: shf = info.filter(filtroShfDistancia);

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
                            <td rowspan={getDuracao(g[h][d], d)}>
                                <EscolhaBox 
                                    info={removeExtraFromEscolha(g[h][d])} 
                                    on:popup
                                    on:remove
                                />
                            </td>
                        {:else if g[h][d] === null}
                            <td></td>
                        {/if}
                    {/each}
                </tr>
            {/each}

        </tbody>
    </table>
    {#if true}
        <div id="shf">
            {#each shf as e}
                <EscolhaBox info={removeExtraFromEscolha(e)} on:popup on:remove />
            {/each}
        </div>
    {/if}
    
</div>

<style>
    #grade {
        position: relative;
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

    #shf {
        z-index: 3;
        position: sticky;
        bottom: 0;
        left: 0;

        background: rgba(255, 0, 0, 0.8);

        height: 20%;
        width: 100%;
    } 
</style>