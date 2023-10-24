<script lang="ts">
	import GradeContainer from "./GradeContainer.svelte";

	import { checkLogin } from "$lib/api";
	import { onMount } from "svelte";
	import { page } from "$app/stores";

	let codigoGrade: string | null = null;

	onMount(() => {
		// verifica o login
		checkLogin()	

		// verifica se a pagina possui um codigo de grade
		const params = $page.url.searchParams;
		const codigo = params.get("g");
		if (codigo !== null) {
			// salva o codigo e apaga da historia
			codigoGrade = codigo;
			window.history.replaceState({}, "", $page.url.pathname);
		}
	});
</script>

<div id="grade-page">
	<GradeContainer {codigoGrade} />
</div>

<style>
	#grade-page {
		width: 100%;
		height: 100%;

		display: flex;
		flex-flow: column nowrap;
		justify-content: center;
		align-items: center;
		gap: 1rem;
	}
</style>