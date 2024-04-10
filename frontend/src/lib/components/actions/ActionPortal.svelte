<script lang="ts">
	import Portal from 'svelte-portal';
</script>

<Portal target="#actions">
	<div class="inline-flex h-full">
		<slot />
	</div>
</Portal>
