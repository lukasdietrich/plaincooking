<script lang="ts">
	import Portal from 'svelte-portal';
</script>

<Portal target="#actions">
	<div class="inline-flex h-14 bg-emerald-700 text-xl">
		<slot />
	</div>
</Portal>
