import type { Readable } from 'svelte/store';
import { writable } from 'svelte/store';

export interface NotificationInit {
	type: 'info' | 'warn' | 'error';
	text: string;
	duration?: number;
}

interface Notification extends NotificationInit {
	id: string;
}

const { subscribe, update } = writable<Notification[]>([]);

const push = (notification: Notification) => {
	update((notifications) => [...notifications, notification]);
};

const remove = (notification: Notification) => {
	update((notifications) => notifications.filter((n) => n !== notification));
};

export const notifications: Readable<Notification[]> = { subscribe };

export const notify = (init: NotificationInit) => {
	const notification: Notification = {
		...init,
		id: Math.random().toString(16)
	};

	const duration = init.duration ?? 5000;

	push(notification);
	setTimeout(() => remove(notification), duration);
};
