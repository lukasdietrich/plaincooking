<script lang="ts" generics="T">
	let dialog: HTMLDialogElement;
	// eslint-disable-next-line no-undef
	let resolve: ((value: T | undefined) => void) | undefined;

	// eslint-disable-next-line no-undef
	export const show = async (): Promise<T | undefined> => {
		dialog.showModal();

		if (resolve) {
			resolve(undefined);
		}

		return new Promise((_resolve) => (resolve = _resolve));
	};

	// eslint-disable-next-line no-undef
	const close = (value: T | undefined) => {
		dialog.close();

		if (resolve) {
			resolve(value);
			resolve = undefined;
		}
	};
</script>

<dialog bind:this={dialog} class="bg-white rounded-sm shadow-sm shadow-black/30">
	<div class="relative px-10 py-8">
		<div class="absolute top-0 right-0">
			<button
				on:click={() => close(undefined)}
				class="flex items-center text-sm px-2 py-1 transition hover:(bg-red text-white)"
			>
				<i class="icon-x"></i>
			</button>
		</div>
		<div>
			<slot {close} />
		</div>
	</div>
</dialog>

<style>
	dialog::backdrop {
		--at-apply: bg-black/40;
	}
</style>
