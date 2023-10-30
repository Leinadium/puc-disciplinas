<script lang="ts">
	import TurmaSelecao from "../components/grade/turma/TurmaSelecao.svelte";
	import type { TabelaHorarios } from "../types/data";
	import type { UIDiaDisciplina, UIHoraDisciplina } from "../types/ui";

    let x: TabelaHorarios = new Map<UIDiaDisciplina, Set<UIHoraDisciplina>>();
</script>

<div></div>