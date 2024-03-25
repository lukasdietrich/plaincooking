declare module 'svelte/elements' {
	export interface HTMLInputAttributes {
		'on:dragEnterBody'?: () => void;
		'on:dragLeaveBody'?: () => void;
		'on:dropBody'?: (event: CustomEvent<File>) => void;
	}
}

export {};
