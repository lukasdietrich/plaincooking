<script lang="ts" generics="T">
	import { setClose } from './context';

	let dialog: HTMLDialogElement;
	let resolve: ((value: T | undefined) => void) | undefined;

	export const show = async (): Promise<T | undefined> => {
		dialog.showModal();

		if (resolve) {
			resolve(undefined);
		}

		return new Promise((_resolve) => (resolve = _resolve));
	};

	const close = (value?: T) => {
		dialog.close();

		if (resolve) {
			resolve(value);
			resolve = undefined;
		}
	};

	setClose<T>(close);
</script>

<dialog bind:this={dialog} class="bg-white rounded-sm shadow-sm shadow-black/30">
	<div class="relative px-10 py-8">
		<div class="absolute top-0 right-0">
			<button
				on:click={() => close()}
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

<style lang="css">
	dialog::backdrop {
		--at-apply: bg-black/40; /**/
	}
</style>
