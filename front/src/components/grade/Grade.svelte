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

    // $: console.log(info);

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
    <table class:com-shf={shf.length > 0}>
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
                    <td>
                        {formatHora(h)}
                    </td>
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
    {#if shf.length > 0}
        <div id="shf">
            <span id="titulo-shf">Sem horário definido / À distância</span>
            <div id="lista-shf">
                {#each shf as e}
                    <div class="disciplina-shf">
                        <EscolhaBox 
                            info={removeExtraFromEscolha(e)}
                            on:popup
                            on:remove
                        />
                    </div>    
                {/each}
            </div>
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

        height: 100%;
        
        /* para compensar a barra de scroll*/
        padding: 5px 10px 5px 5px;

        background: var(--color-main-2);
        border-radius: var(--border-radius);
        --border-row: 1px solid var(--color-main-3);

    }
	table {
        box-sizing: border-box;
        table-layout: fixed;
        border-collapse: collapse;
		height: 100%;
        width: 100%;
	}

    .com-shf {
        height: 87% !important;
    }

    thead {
        height: 4.8%;
        border-bottom: var(--border-row);
    }

    tbody {
        height: 95.2%;
    }

    td, th {
        padding: 5px;
        margin: 0;
        width: 100%;
    }

    tr {
        height: 6.8%;
        border-bottom: var(--border-row);
    }

    /* primeiro td (horas) */
    td:first-child, th:first-child {
        /* width: 10% !important; */
        width: 40px;
        padding: 0;
        font-size: 0.8em;
        white-space: nowrap;
        text-align: center;
        border-right: 1px solid var(--color-main-3);
    }

    #shf {
        z-index: 3;
        position: sticky;
        bottom: 0;
        left: 0;

        background: var(--color-mono-1);

        box-sizing: border-box;
        height: 13%;
        width: 100%;
        display: flex;
        flex-flow: column nowrap;
        justify-content: flex-start;
        align-items: center;
        border-radius: var(--border-radius);

    }

    #lista-shf {
        width: 100%;
        height: 100%;
        overflow-x: scroll;

        box-sizing: border-box;
        padding: 0 1% 1% 1%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-start;
        align-items: flex-end;
        gap: 2%;
    }

    #titulo-shf {
        box-sizing: border-box;
        padding: 0.2rem;
        font-size: 0.8rem;
        font-style: italic;
        color: var(--color-main-4)
    }

    .disciplina-shf {
        flex-shrink: 0;
        width: fit-content;
        height: 100%;
        max-width: 20%;
        
    }
</style>