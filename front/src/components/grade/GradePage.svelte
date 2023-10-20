<script lang="ts">
	import GradeContainer from "./GradeContainer.svelte";

	import { checkLogin } from "$lib/api";
	import { onMount } from "svelte";

	onMount(async () => {
		await checkLogin()	// verifica o login
	});
</script>

<div id="grade-page">
	<GradeContainer />
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