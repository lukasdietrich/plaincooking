import { setContext, getContext } from 'svelte';

export type Close<T> = (value?: T) => void;

const closeFunctionKey = {};

export function setClose<T>(close: Close<T>) {
	setContext(closeFunctionKey, close);
}

export function getClose<T>(): Close<T> {
	return getContext(closeFunctionKey);
}
