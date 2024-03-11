<script lang="ts">
	import Portal from 'svelte-portal';
</script>

<Portal target="#actions">
	<div class="flex text-xl">
		<slot />
	</div>
</Portal>
