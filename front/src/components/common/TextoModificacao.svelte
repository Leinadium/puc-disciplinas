<script lang="ts">
	import { calcularDiferenca } from "$lib/utils";
	import type { Modificacao } from "../../types/data";

    export let modificacao: Modificacao | null = null;

    let diferenca: string = "";
    $: diferenca = modificacao?.dataGeral ? calcularDiferenca(new Date(modificacao.dataGeral)) + " atrás" : "Não disponível";
</script>

Última modificacação: {diferenca}