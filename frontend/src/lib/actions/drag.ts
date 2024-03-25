import type { Action } from 'svelte/action';

export const dragOnBody: Action<HTMLInputElement> = (node: HTMLInputElement) => {
	let counter = 0;

	const onEnter = () => {
		if (counter === 0) {
			node.dispatchEvent(new CustomEvent('dragEnterBody'));
		}

		counter++;
	};

	const onLeave = () => {
		counter--;

		if (counter === 0) {
			node.dispatchEvent(new CustomEvent('dragLeaveBody'));
		}
	};

	const onDrop = () => {
		counter = 0;
	};

	document.body.addEventListener('dragenter', onEnter);
	document.body.addEventListener('dragleave', onLeave);
	document.body.addEventListener('drop', onDrop);

	return {
		destroy() {
			document.body.removeEventListener('dragenter', onEnter);
			document.body.removeEventListener('dragleave', onLeave);
			document.body.removeEventListener('drop', onDrop);
		}
	};
};

export const dropOnBody: Action<HTMLInputElement> = (node: HTMLInputElement) => {
	const onDrop = (event: DragEvent) => {
		const file = getAcceptedFile(node, event);
		if (file) {
			event.preventDefault();

			node.dispatchEvent(
				new CustomEvent<File>('dropBody', {
					detail: file
				})
			);
		}
	};

	const onDragOver = (event: DragEvent) => {
		event.preventDefault();
	};

	document.body.addEventListener('drop', onDrop);
	document.body.addEventListener('dragover', onDragOver);

	return {
		destroy() {
			document.body.removeEventListener('drop', onDrop);
			document.body.removeEventListener('dragover', onDragOver);
		}
	};
};

function getAcceptedFile(node: HTMLInputElement, event: DragEvent): File | undefined {
	const file = event.dataTransfer?.files?.item(0);
	if (!file) {
		return undefined;
	}

	if (!node.accept.includes(file.type)) {
		return undefined;
	}

	return file;
}
