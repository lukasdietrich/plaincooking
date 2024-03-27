<script>
	import '@unocss/reset/tailwind.css';
	import 'lucide-static/font/lucide.css';
	import '@fontsource-variable/readex-pro';
	import '@fontsource-variable/crimson-pro';
	import '@fontsource-variable/inconsolata';

	import { page } from '$app/stores';
	import { Header } from '$lib/components/wrapper';
	import { Notifications } from '$lib/components/notifications';

	$: title = $page.data['title'] ?? 'Plaincooking';
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<div class="min-h-screen bg-white font-sans">
	<Header />
	<Notifications />

	<main class="max-w-screen-lg mx-auto p-5">
		<slot />
	</main>
</div>
