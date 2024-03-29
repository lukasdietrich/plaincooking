import type { Middleware } from 'openapi-fetch';

const globalMiddleware: Middleware[] = [];

export function register(middleware: Middleware) {
	globalMiddleware.push(middleware);
}

export function middleware(): Middleware[] {
	return globalMiddleware;
}
