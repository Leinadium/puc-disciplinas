<script lang="ts">
	import Container from "./Container.svelte";

	import { checkLogin } from "$lib/api";
	import { onMount } from "svelte";

	onMount(checkLogin);
</script>

<div id="grade-page">
	<Container />
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