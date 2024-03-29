<script lang="ts">
	import { fly } from 'svelte/transition';
	import { t } from '$lib/i18n';
	import { notifications } from './store';
</script>

<div class="fixed bottom-5 left-5 flex flex-col space-y-2">
	{#each $notifications as notification (notification.id)}
		{@const type = notification.type}
		<div
			transition:fly={{ x: -100, duration: 600 }}
			class="flex items-center space-x-3 bg-slate-100 px-5 py-3 rounded border-b-2"
			class:info={type === 'info'}
			class:warn={type === 'warn'}
			class:error={type === 'error'}
		>
			<i
				class="text-xl"
				class:icon-info={type === 'info'}
				class:icon-triangle-alert={type === 'warn'}
				class:icon-octagon-x={type === 'error'}
			></i>
			<span>{$t(notification.text)}</span>
		</div>
	{/each}
</div>

<style>
	.info {
		--at-apply: bg-blue-600 text-blue-50 border-blue-800;
	}

	.warn {
		--at-apply: bg-yellow-600 text-yellow-50 border-yellow-800;
	}

	.error {
		--at-apply: bg-red-600 text-red-50 border-red-800;
	}
</style>
